package tiktok

import (
	"errors"
	"fmt"
	json "github.com/bytedance/sonic"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type TikTokUser struct {
	Url   string
	SecId string
	Name  string
	Face  string

	LastAweme Aweme
	AwemeChan chan Aweme

	IsListed  bool
	GroupList []int64
}

func (tk *TikTokUser) HasGroupById(groupId int64) bool {
	for _, v := range tk.GroupList {
		if v == groupId {
			return true
		}
	}
	return false
}
func (u *TikTokUser) AddGroup(groupId int64) {
	u.GroupList = append(u.GroupList, groupId)
	log.Warnf(fmt.Sprintf("[抖音推送]  用户 %s(%s) 已被添加到群 %d 的监听列表", u.Name, u.Url, groupId))
	//face, _ := DownloadImg(u.Face)
	//u.MsgChan <- PushMsg{
	//	Type: "new",
	//	Msg:  fmt.Sprintf("已添加 %s UID: %d 到微博关注列表", u.Name, u.Id),
	//	Img:  face,
	//}
}
func (u *TikTokUser) RemoveGroup(groupId int64) {
	for k, v := range u.GroupList {
		if v == groupId {
			u.GroupList = append(u.GroupList[:k], u.GroupList[k+1:]...)
			log.Warnf(fmt.Sprintf("[抖音推送]  用户 %s(%s) 已被从群 %d 的监听列表移除", u.Name, u.Url, groupId))
		}
	}
}

var TikTokUserList = make(map[string]*TikTokUser)

func NewTikTokUser(url string, init bool, groupIds ...int64) (*TikTokUser, error) {
	var user = new(TikTokUser)
	user.Url = url
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/105.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	path := strings.Split(resp.Request.URL.Path, "/")
	if len(path) < 3 {
		return user, errors.New("error url")
	}
	SecUid := path[3]
	nameRegexp := regexp.MustCompile("<div class=\"name\">(.+?)</div>")
	faceRegexp := regexp.MustCompile("<img class=\"avatar\" src=\"(.+?)\"")
	names := nameRegexp.FindStringSubmatch(string(body))
	if len(names) < 2 {
		return user, errors.New("获取用户名失败")
	}
	faces := faceRegexp.FindStringSubmatch(string(body))
	if len(faces) < 2 {
		return user, errors.New("获取用户头像失败")
	}
	log.Printf("[抖音推送] 用户 %s 初始化成功！", names[1])
	user = &TikTokUser{
		Url:       url,
		SecId:     SecUid,
		Name:      names[1],
		Face:      faces[1],
		AwemeChan: make(chan Aweme),
		GroupList: groupIds,
	}
	return user, nil
}

func (tk *TikTokUser) GetLastAweme() error {
	api := fmt.Sprintf("https://m.douyin.com/web/api/v2/aweme/post/?reflow_source=reflow_page&sec_uid=%s&count=21&max_cursor=0", tk.SecId)
	client := &http.Client{}
	req, err := http.NewRequest("GET", api, nil)
	req.Header.Set("User-Agent", "Mozill/a/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/108.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Warnln(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Warnln("[抖音推送] 读取响应失败")
		return err
	}
	var ttinfo TikTokInfo
	err = json.Unmarshal(body, &ttinfo)
	if err != nil {
		log.Warnln(string(body), err)
		return err
	}
	al := AwemeList(ttinfo.AwemeList)
	if al.MaxAwemeId().AwemeId2Uinit() <= tk.LastAweme.AwemeId2Uinit() {
		return nil
	}
	tk.LastAweme = al.MaxAwemeId()
	return nil
}

func (tk *TikTokUser) Listen() {
	for err := tk.GetLastAweme(); err != nil; {
		//log.Warnln(err)
		log.Warnln(err)
		time.Sleep(time.Second * 5)
	}
	_ = tk.GetLastAweme()
	tk.IsListed = true
	for tk.IsListed {
		temp := tk.LastAweme.AwemeId2Uinit()
		_ = tk.GetLastAweme()
		if tk.LastAweme.AwemeId2Uinit() > temp {
			log.Warnf("[抖音推送] 用户 %s 抖音新作品 %s %s", tk.Name, tk.LastAweme.Desc, tk.LastAweme.AwemeId)
			tk.AwemeChan <- tk.LastAweme
		}
		time.Sleep(time.Second * 2)
	}
	close(tk.AwemeChan)
}
