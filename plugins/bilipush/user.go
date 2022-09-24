package bilipush

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"strings"

	json "github.com/bytedance/sonic"
	"golang.org/x/image/webp"
	"google.golang.org/grpc/metadata"

	"io"
	"net/http"
	"strconv"

	dy "github.com/Mrs4s/go-cqhttp/plugins/bilipush/grpc_api/bilibili/app/dynamic/v2"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Mid       int64  // B站 UID
	Uname     string // B站用户名
	LastOid   int64  // B站最新动态ID
	LastDy    *dy.DynamicItem
	LastDyUrl string // B站最新动态链接
	TopDyOid  string // B站置顶动态ID

	LivingInfo *LivingInfo // 直播信息
	OtherInfo  *OtherInfo  // 其他信息
	//ImgChan      chan []byte  // 图片
	MsgChan      chan PushMsg // 消息推送管道
	ErrorMessage string

	IsListed bool // 是否已经被列入监听列表

	GroupList []int64 // 群号
}

type LivingInfo struct {
	Code    int64               `json:"code"`
	Msg     string              `json:"msg"`
	Message string              `json:"message"`
	Data    map[string]RoomInfo `json:"data"`
}

type RoomInfo struct {
	Title            string `json:"title"`
	RoomId           int64  `json:"room_id"`
	Uid              int64  `json:"uid"`
	Online           int64  `json:"online"`
	LiveTime         int64  `json:"live_time"`
	LiveStatus       int64  `json:"live_status"`
	ShortId          int64  `json:"short_id"`
	Area             int64  `json:"area"`
	AreaName         string `json:"area_name"`
	AreaV2Id         int64  `json:"area_v2_id"`
	AreaV2Name       string `json:"area_v2_name"`
	AreaV2ParentName string `json:"area_v2_parent_name"`
	AreaV2ParentId   int64  `json:"area_v2_parent_id"`
	Uname            string `json:"uname"`
	Face             string `json:"face"`
	TagName          string `json:"tag_name"`
	Tags             string `json:"tags"`
	CoverFromUser    string `json:"cover_from_user"`
	Keyframe         string `json:"keyframe"`
	LockTill         string `json:"lock_till"`
	HiddenTill       string `json:"hidden_till"`
	BroadcastType    int64  `json:"broadcast_type"`
}
type OtherInfo struct {
	Code    int64  `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    []struct {
		Mid     int64  `json:"mid"`
		Name    string `json:"name"`
		Sex     string `json:"sex"`
		Face    string `json:"face"`
		Sign    string `json:"sign"`
		Rank    int64  `json:"rank"`
		Level   int64  `json:"level"`
		Silence int64  `json:"silence"`
		Vip     struct {
			Type       int64 `json:"type"`
			Status     int64 `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int64 `json:"vip_pay_type"`
			ThemeType  int64 `json:"theme_type"`
			Label      struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgStyle               int64  `json:"bg_style"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				UseImgLabel           bool   `json:"use_img_label"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			} `json:"label"`
			AvatarSubscript    int64  `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int64  `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			TvVipStatus        int64  `json:"tv_vip_status"`
			TvVipPayType       int64  `json:"tv_vip_pay_type"`
			TvDueDate          int64  `json:"tv_due_date"`
		} `json:"vip"`
		Pendant struct {
			Pid               int64  `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int64  `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Nameplate struct {
			Nid        int64  `json:"nid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			ImageSmall string `json:"image_small"`
			Level      string `json:"level"`
			Condition  string `json:"condition"`
		} `json:"nameplate"`
		Official struct {
			Role  int64  `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int64  `json:"type"`
		} `json:"official"`
		Birthday       int64 `json:"birthday"`
		IsFakeAccount  int64 `json:"is_fake_account"`
		IsDeleted      int64 `json:"is_deleted"`
		InRegAudit     int64 `json:"in_reg_audit"`
		FaceNft        int64 `json:"face_nft"`
		FaceNftNew     int64 `json:"face_nft_new"`
		IsSeniorMember int64 `json:"is_senior_member"`
	} `json:"data"`
}

type PushMsg struct {
	Type string
	Msg  string
	Img  []byte
}

var UserList = make(map[int64]*User)

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

func NewUser(mid int64, init bool, groupIds ...int64) (*User, error) {
	if len(groupIds) == 0 {
		return nil, fmt.Errorf("没有需要推送的群")
	}
	api := "https://api.vc.bilibili.com/account/v1/user/cards?uids=" + fmt.Sprintf("%d", mid)
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var oinfo OtherInfo
	body, _ := io.ReadAll(resp.Body)
	//uname := gjson.Get(string(body), "card.name").String()
	err = json.Unmarshal(body, &oinfo)
	// uname := oinfo.Card.Name
	if len(oinfo.Data) == 0 || err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	uname := oinfo.Data[0].Name
	//imgChan := make(chan []byte, 1)
	msgChan := make(chan PushMsg, 1)
	face, _ := DownloadImg(oinfo.Data[0].Face)
	if !init {
		msgChan <- PushMsg{
			Type: "new",
			Msg:  fmt.Sprintf("已添加 %s UID: %d 到B站关注列表", uname, mid),
			Img:  face,
		}
	}
	livingInfo := new(LivingInfo)
	return &User{
		Mid:        mid,
		Uname:      uname,
		MsgChan:    msgChan,
		GroupList:  groupIds,
		LivingInfo: livingInfo,
	}, nil
}
func (u *User) HasGroupById(groupId int64) bool {
	for _, v := range u.GroupList {
		if v == groupId {
			return true
		}
	}
	return false
}
func (u *User) AddGroup(groupId int64) {
	u.GroupList = append(u.GroupList, groupId)
	log.Warnf(fmt.Sprintf("[B站推送]  用户 %s(%d) 已被添加到群 %d 的监听列表", u.Uname, u.Mid, groupId))
	face, _ := DownloadImg(u.OtherInfo.Data[0].Face)
	u.MsgChan <- PushMsg{
		Type: "new",
		Msg:  fmt.Sprintf("已添加 %s UID: %d 到关注列表", u.Uname, u.Mid),
		Img:  face,
	}
}
func (u *User) RemoveGroup(groupId int64) {
	for k, v := range u.GroupList {
		if v == groupId {
			u.GroupList = append(u.GroupList[:k], u.GroupList[k+1:]...)
			log.Warnf(fmt.Sprintf("[B站推送]  用户 %s(%d) 已被从群 %d 的监听列表移除", u.Uname, u.Mid, groupId))
		}
	}
}

func (u *User) GetLastOid() error {
	dyClient := dy.NewDynamicClient(GrpcClient)
	dyReq := &dy.DynSpaceReq{
		HostUid:       u.Mid,
		HistoryOffset: "",
	}
	grpcCtx := metadata.NewOutgoingContext(context.Background(), GetMetadata())
	retryTimes := 0
	var dyResp *dy.DynSpaceRsp
	var err error
	for retryTimes < 3 {
		dyResp, err = dyClient.DynSpace(grpcCtx, dyReq)
		if err != nil {
			retryTimes++
			//log.Error("BiliBili GRPC 服务器 出错, 重新连接中... ", err, retryTimes)
			continue
		}
		break
	}
	if retryTimes == 3 {
		return fmt.Errorf("[B站推送]  BiliBili GRPC 服务器 出错")
	}
	for _, d := range dyResp.List {
		isTop := false
		for _, m := range d.Modules {
			if m.ModuleType == dy.DynModuleType_module_author {
				if m.GetModuleAuthor().IsTop {
					isTop = true
					u.TopDyOid = d.Extend.DynIdStr
				}
			}
		}
		oid, _ := strconv.ParseInt(d.Extend.DynIdStr, 10, 64)
		url := "https://t.bilibili.com/h5/dynamic/detail/" + d.Extend.DynIdStr
		if !isTop && oid > u.LastOid {
			u.LastOid = oid
			u.LastDy = d
			u.LastDyUrl = url
			break
		}
		if oid == 0 {
			log.Errorln(" [B站推送] 获取动态失败, 请检查网络连接")
		}
	}
	return nil
}

func (u *User) GetLivingInfo() {
	api := "https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids"
	body := fmt.Sprintf("{\"uids\": [%d]}", u.Mid)
	resp, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	var info LivingInfo
	_ = json.Unmarshal(b, &info)
	if info.Msg != "success" {
		return
	}
	if err != nil {
		log.Errorln(string(b), err)
		u.LivingInfo = &info
		return
	}
	if info.Data[strconv.FormatInt(u.Mid, 10)].Title == "" {
		log.Warnln(u.Uname, u.Mid, "没有开通直播间")
		return
	}
	u.LivingInfo = &info
}
func (u *User) GetOtherInfo() {
	api := "https://api.vc.bilibili.com/account/v1/user/cards?uids=" + fmt.Sprintf("%d", u.Mid)
	resp, err := http.Get(api)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var oinfo OtherInfo
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &oinfo)
	if err != nil || oinfo.Code != 0 {
		log.Errorln(err, string(body))
		return
	}
	u.OtherInfo = &oinfo
}

func (u *User) ListenUser() {
	mid := strconv.FormatInt(u.Mid, 10)
	u.IsListed = true
	for err := u.GetLastOid(); err != nil; {
	}
	u.GetLivingInfo()
	u.GetOtherInfo()
	// go u.ListenUserInfo()
	//log.Println(u.LivingInfo)
	for u.IsListed {
		if _, ok := u.LivingInfo.Data[mid]; ok {
			title := u.LivingInfo.Data[mid].Title                 // 保存上一次的直播间标题
			liveStatus := u.LivingInfo.Data[mid].LiveStatus       // 保存上一次的直播间状态
			coverFromUser := u.LivingInfo.Data[mid].CoverFromUser // 保存上一次的直播间封面
			u.GetLivingInfo()                                     // 获取最新的直播间信息
			if u.LivingInfo.Data[mid].LiveStatus != liveStatus {
				if u.LivingInfo.Data[mid].LiveStatus == 1 {
					log.Warnf("[B站推送] 用户", u.Uname, "开播了")
					img, _ := DownloadImg(u.LivingInfo.Data[mid].CoverFromUser)
					u.MsgChan <- PushMsg{
						Type: "live",
						Msg:  fmt.Sprintf("%s 开播了：\n标题：%s", u.Uname, u.LivingInfo.Data[mid].Title),
						Img:  img,
					}
				} else if u.LivingInfo.Data[mid].LiveStatus == 0 || u.LivingInfo.Data[mid].LiveStatus == 2 {
					log.Warnf("[B站推送] 用户", u.Uname, "下播了")
					img, _ := DownloadImg(coverFromUser)
					u.MsgChan <- PushMsg{
						Type: "live",
						Msg:  fmt.Sprintf("%s 下播了：", u.Uname),
						Img:  img,
					}
				}
			}
			if u.LivingInfo.Data[mid].Title != title {
				log.Warnf("[B站推送] 用户", u.Uname, "直播间标题变更为", u.LivingInfo.Data[mid].Title)
				img, _ := DownloadImg(coverFromUser)
				u.MsgChan <- PushMsg{
					Type: "live",
					Msg:  fmt.Sprintf("%s 直播间标题变更：\n原: %s\n新: %s", u.Uname, title, u.LivingInfo.Data[mid].Title),
					Img:  img,
				}
			}
			if u.LivingInfo.Data[mid].CoverFromUser != coverFromUser {
				log.Warnf("[B站推送] 用户", u.Uname, "直播间封面变更")
				img, _ := DownloadImg(coverFromUser)
				u.MsgChan <- PushMsg{
					Type: "live",
					Msg:  fmt.Sprintf("%s 直播间封面变更", u.Uname),
					Img:  img,
				}
			}
		}
		oid := u.LastOid // 保存上一次的动态ID
		for err := u.GetLastOid(); err != nil; {
		} // 获取最新动态ID
		if u.LastOid > oid {
			log.Warnf("[B站推送] 用户", u.Uname, "发布了新动态", u.LastDyUrl)
			img, err := GetScreenshot(u.LastDyUrl)
			if err != nil {
				log.Errorln(err)
				if err.Error() == "动态进入审核" {
					var descText string
					var imgs [][]byte
					var newImg []byte
					for _, m := range u.LastDy.Modules {
						if m.ModuleType == dy.DynModuleType_module_dynamic {
							dynamic := m.GetModuleDynamic()
							if dynamic == nil {
								continue
							}
							pics := dynamic.GetDynDraw()
							if pics == nil {
								continue
							}
							for _, p := range pics.Items {
								img, err := DownloadImg(p.Src)
								if strings.HasSuffix(p.Src, ".webp") {
									webpImg, _ := Webp2Jpeg(img)
									imgs = append(imgs, webpImg)
									continue
								}
								if err != nil {
									continue
								}
								imgs = append(imgs, img)
							}
						}
						if m.ModuleType == dy.DynModuleType_module_desc {
							desc := m.GetModuleDesc()
							descText = desc.Text
						}
					}
					if len(imgs) > 0 {
						newImg = MergeImages(imgs...)
					}
					msg := fmt.Sprintf("%s 发布新动态,但是进入审核,无法截图,原动态如下:\n %s", u.Uname, descText)
					u.MsgChan <- PushMsg{
						Type: "dynamic",
						Msg:  msg,
						Img:  newImg,
					}
				} else {
					u.ErrorMessage = err.Error()
					u.MsgChan <- PushMsg{
						Type: "dynamic",
						Msg:  fmt.Sprintf("%s 发布新动态：", u.Uname),
						Img:  nil,
					}
				}
			} else {
				u.MsgChan <- PushMsg{
					Type: "dynamic",
					Msg:  fmt.Sprintf("%s 发布新动态：", u.Uname),
					Img:  img,
				}
			}
		}

		name := u.OtherInfo.Data[0].Name // 保存上一次的用户名
		face := u.OtherInfo.Data[0].Face // 保存上一次的头像
		sign := u.OtherInfo.Data[0].Sign // 保存上一次的签名
		// attentions := u.OtherInfo.Card.Attentions // 保存上一次的关注数
		pendants := u.OtherInfo.Data[0].Pendant // 保存上一次的装扮
		u.GetOtherInfo()                        // 获取最新信息
		if u.OtherInfo.Data[0].Name != name {
			log.Warnf("[B站推送] %s 昵称变更 -> %s", name, u.Uname)
			img, _ := DownloadImg(u.OtherInfo.Data[0].Face)
			u.MsgChan <- PushMsg{
				Type: "other",
				Msg:  fmt.Sprintf("%s 昵称变更 -> %s", name, u.Uname),
				Img:  img,
			}
		}
		if u.OtherInfo.Data[0].Face != face {
			log.Warnf("[B站推送] %s 头像变更", u.Uname)
			img, _ := DownloadImg(u.OtherInfo.Data[0].Face)
			u.MsgChan <- PushMsg{
				Type: "other",
				Msg:  fmt.Sprintf("%s 头像变更", u.Uname),
				Img:  img,
			}
		}
		if u.OtherInfo.Data[0].Sign != sign {
			log.Warnf("[B站推送] %s 签名变更: %s", u.Uname, u.OtherInfo.Data[0].Sign)
			u.MsgChan <- PushMsg{
				Type: "other",
				Msg:  fmt.Sprintf("%s 签名变更:\n%s", u.Uname, u.OtherInfo.Data[0].Sign),
				Img:  nil,
			}
		}
		if u.OtherInfo.Data[0].Pendant.Pid != pendants.Pid {
			log.Warnf("[B站推送] %s 装扮变更\n%s -> %s", u.Uname, pendants.Name, u.OtherInfo.Data[0].Pendant.Name)
			//img, _ := DownloadImg(u.OtherInfo.Data[0].Face)
			u.MsgChan <- PushMsg{
				Type: "other",
				Msg:  fmt.Sprintf("%s 装扮变更\n%s -> %s", u.Uname, pendants.Name, u.OtherInfo.Data[0].Pendant.Name),
				Img:  nil,
			}
		}
	}
	close(u.MsgChan)
}

func Webp2Jpeg(i []byte) ([]byte, error) {
	img, err := webp.Decode(bytes.NewReader(i))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MergeImages(img ...[]byte) []byte {
	var images []image.Image
	for _, i := range img {
		img, _, _ := image.Decode(bytes.NewReader(i))
		images = append(images, img)
	}
	maxWidth := 0
	totalHeight := 0
	for _, img := range images {
		b := img.Bounds()
		w := b.Max.X
		h := b.Max.Y
		if w > maxWidth {
			maxWidth = w
		}
		totalHeight += h
	}
	newImg := image.NewRGBA(image.Rect(0, 0, maxWidth, totalHeight))
	draw.Draw(newImg, newImg.Bounds(), image.White, image.Point{}, draw.Src)
	offsetY := 0
	for _, img := range images {
		b := img.Bounds()
		w := b.Max.X
		h := b.Max.Y
		draw.Draw(newImg, image.Rect(0, offsetY, w, offsetY+h), img, image.Point{}, draw.Src)
		offsetY += h
	}
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, newImg, nil)
	return buf.Bytes()
}

func GetChangedAttention(old, new []int64) (OtherInfo, error) {
	var uidChanged int64
	if len(old) > len(new) {
		// 取关了
		for _, o := range old {
			if !Contains(new, o) {
				uidChanged = o
				break
			}
		}
	} else {
		// 关注了
		for _, n := range new {
			if !Contains(old, n) {
				uidChanged = n
				break
			}
		}
	}
	var info OtherInfo
	api := "https://account.bilibili.com/api/member/getCardByMid?mid=" + fmt.Sprintf("%d", uidChanged)
	resp, err := http.Get(api)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, err
	}
	return info, nil
}

func Contains(int64s []int64, o int64) bool {
	for _, i := range int64s {
		if i == o {
			return true
		}
	}
	return false
}
