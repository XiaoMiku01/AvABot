package weibo

import (
	"fmt"
	json "github.com/bytedance/sonic"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type WeiBoUser struct {
	Id          int64
	Name        string
	Face        string
	Containerid string
	LastCardMid int64
	LastCardUrl string
	MsgChan     chan PushMsg // 消息推送管道

	IsListed  bool
	GroupList []int64
}

type PushMsg struct {
	Type string
	Msg  string
	Img  []byte
}

var WeiBoUserList = make(map[int64]*WeiBoUser)

func (u *WeiBoUser) HasGroupById(groupId int64) bool {
	for _, v := range u.GroupList {
		if v == groupId {
			return true
		}
	}
	return false
}
func (u *WeiBoUser) AddGroup(groupId int64) {
	u.GroupList = append(u.GroupList, groupId)
	log.Warnf(fmt.Sprintf("[微博推送]  用户 %s(%d) 已被添加到群 %d 的监听列表", u.Name, u.Id, groupId))
	face, _ := DownloadImg(u.Face)
	u.MsgChan <- PushMsg{
		Type: "new",
		Msg:  fmt.Sprintf("已添加 %s UID: %d 到微博关注列表", u.Name, u.Id),
		Img:  face,
	}
}
func (u *WeiBoUser) RemoveGroup(groupId int64) {
	for k, v := range u.GroupList {
		if v == groupId {
			u.GroupList = append(u.GroupList[:k], u.GroupList[k+1:]...)
			log.Warnf(fmt.Sprintf("[微博推送]  用户 %s(%d) 已被从群 %d 的监听列表移除", u.Name, u.Id, groupId))
		}
	}
}

func NewWeiBoUser(id int64, init bool, groupIds ...int64) (*WeiBoUser, error) {
	var msgChan chan PushMsg = make(chan PushMsg, 1)
	var wb = WeiBoUser{Id: id, MsgChan: msgChan, GroupList: groupIds}
	wb.GetContainerId()
	wb.GetLastCard()
	if wb.Name != "" {
		log.Println("[微博推送] 微博用户:", wb.Name)
		if !init {
			face, _ := DownloadImg(wb.Face)
			wb.MsgChan <- PushMsg{
				Type: "new",
				Msg:  fmt.Sprintf("已添加 %s ID: %d 到微博关注列表", wb.Name, wb.Id),
				Img:  face,
			}
		}
		return &wb, nil
	}
	return nil, fmt.Errorf("[微博推送] 未发现用户")
}

func (wb *WeiBoUser) GetContainerId() {
	var url0 = "https://m.weibo.cn/u/" + strconv.FormatInt(wb.Id, 10)
	resp, err := http.Get(url0)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	cookies := resp.Cookies()
	for _, c := range cookies {
		if c.Name == "M_WEIBOCN_PARAMS" {
			urlencode, _ := url.QueryUnescape(c.Value)
			fid := strings.Split(urlencode, "&")[0]
			id := strings.Split(fid, "=")[1]
			idInt64, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				panic(err)
			}
			wb.Containerid = strconv.FormatInt(idInt64+7098e+10, 10)
			return
		}
	}
}

func (wb *WeiBoUser) GetLastCard() int64 {
	var api = "https://m.weibo.cn/api/container/getIndex?containerid=" + wb.Containerid
	client := &http.Client{}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return 0
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/105.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	if resp.StatusCode != 200 {
		log.Warnln("[微博推送] 风控:", resp.StatusCode, "暂停5分钟")
		time.Sleep(time.Minute * 5)
	}
	var weiboinfo WeiBoInfo
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &weiboinfo)
	if err != nil || weiboinfo.Ok != 1 {
		return 0
	}
	var maxMid int64 = 0
	for _, card := range weiboinfo.Data.Cards {
		mid, _ := strconv.ParseInt(card.Mblog.Mid, 10, 64)
		if mid > maxMid {
			maxMid = mid
			wb.LastCardUrl = card.Scheme
			wb.LastCardMid = mid
			wb.Name = card.Mblog.User.ScreenName
			wb.Face = card.Mblog.User.ProfileImageURL
		}
	}
	return maxMid
}
func DownloadImg(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (wb *WeiBoUser) Listen() {
	wb.IsListed = true
	for wb.IsListed {
		nowMid := wb.LastCardMid
		lastMid := wb.GetLastCard()
		if lastMid > nowMid {
			log.Warnf(wb.Name, "发表新微博", wb.LastCardUrl)
			img, _ := GetScreenshot(wb.LastCardUrl)
			wb.MsgChan <- PushMsg{
				Type: "weibo",
				Msg:  fmt.Sprintf("%s 发表新微博：", wb.Name),
				Img:  img,
			}
		}
		time.Sleep(time.Second * 5)
	}
	close(wb.MsgChan)
}
