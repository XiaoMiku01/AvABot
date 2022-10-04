package plugins

import (
	"bytes"
	"fmt"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/modules/servers"
	"github.com/Mrs4s/go-cqhttp/plugins/bilipush"
	"github.com/Mrs4s/go-cqhttp/plugins/tiktok"
	"github.com/Mrs4s/go-cqhttp/plugins/weibo"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

type Plugin struct {
	// 插件名
	Name string
	Func func(e *coolq.Event)
}

type Plugins struct {
	bot        *coolq.CQBot
	pluginList []Plugin
}

var PList Plugins

func init() {
	var newPlugin = &Plugin{
		Name: "指令",
		Func: func(e *coolq.Event) {
			switch e.Raw.PostType {
			case "message":
				groupId, ok := e.Raw.Others["group_id"].(int64)
				if !ok {
					return
				}
				if groupId == 0 || e.Raw.Others["user_id"] != Config.MasterID {
					return
				}
				cmd, ok := e.Raw.Others["raw_message"].(string)
				if !ok {
					return
				}
				msg := message.NewSendingMessage()
				if cmd == "/help" {
					msg.Append(message.NewText("发送\n关注列表 :获取当前群所有推送\n"))
					msg.Append(message.NewText("发送 /关注 + UID :在本群关注该B站用户\n"))
					msg.Append(message.NewText("发送 /取关 + UID :在本群取关该B站用户\n"))
					msg.Append(message.NewText("发送 /微博关注 + 微博ID :在本群关注该微博用户\n"))
					msg.Append(message.NewText("发送 /微博取关 + 微博ID :在本群取关该微博用户\n"))
					msg.Append(message.NewText("发送 /抖音关注 + 抖音主页分享链接 :在本群关注该抖音用户\n"))
					msg.Append(message.NewText("发送 /抖音取关 + 抖音主页分享链接 :在本群取关该抖音用户\n"))
					PList.bot.Client.SendGroupMessage(groupId, msg)
				} else if cmd == "关注列表" {
					msg.Append(message.NewText("已关注的用户:\n"))
					msg.Append(message.NewText("B站:\n"))
					for _, u := range bilipush.UserList {
						if u.HasGroupById(groupId) {
							msg.Append(message.NewText(fmt.Sprintf("%s UID: %d\n", u.Uname, u.Mid)))
						}
					}
					msg.Append(message.NewText("微博:\n"))
					for _, u := range weibo.WeiBoUserList {
						if u.HasGroupById(groupId) {
							msg.Append(message.NewText(fmt.Sprintf("%s ID: %d\n", u.Name, u.Id)))
						}
					}
					msg.Append(message.NewText("抖音:\n"))
					for _, u := range tiktok.TikTokUserList {
						if u.HasGroupById(groupId) {
							msg.Append(message.NewText(fmt.Sprintf("%s 主页: %s\n", u.Name, u.Url)))
						}
					}
					PList.bot.Client.SendGroupMessage(groupId, msg)
				} else if strings.HasPrefix(cmd, "/关注") {
					uidStr := regexp.MustCompile(`\d+`).FindAllString(cmd, -1)
					if len(uidStr) == 0 {
						msg.Append(message.NewText("请输入正确的B站UID"))
						PList.bot.Client.SendGroupMessage(groupId, msg)
						return
					}
					uid, _ := strconv.ParseInt(uidStr[0], 10, 64)
					if user, ok := bilipush.UserList[uid]; ok {
						if user.HasGroupById(groupId) {
							msg.Append(message.NewText(fmt.Sprintf("该用户 %s UID: %d 已经在B站关注列表中", user.Uname, user.Mid)))
							PList.bot.Client.SendGroupMessage(groupId, msg)
						} else {
							user.AddGroup(groupId)
							//msg.Append(message.NewText(fmt.Sprintf("已添加 %s UID: %d 到关注列表", user.Uname, user.Mid)))
							Config.AddUser(&BiliPusher{UID: user.Mid, GroupList: GroupList{groupId}})
						}
					} else {
						user, err := bilipush.NewUser(uid, false, groupId)
						if err != nil {
							msg.Append(message.NewText("获取用户信息失败"))
							PList.bot.Client.SendGroupMessage(groupId, msg)
							return
						}
						go BiliStart(user, PList.bot)
						log.Printf("[B站推送] 用户 %s UID: %d 初始化成功！", user.Uname, user.Mid)
						//msg.Append(message.NewText(fmt.Sprintf("已添加 %s UID: %d 到关注列表", user.Uname, user.Mid)))
						Config.AddUser(&BiliPusher{UID: user.Mid, GroupList: GroupList{groupId}})
						//PList.bot.Client.SendGroupMessage(groupId, msg)
					}
				} else if strings.HasPrefix(cmd, "/取关") {
					uidStr := regexp.MustCompile(`\d+`).FindAllString(cmd, -1)
					if len(uidStr) == 0 {
						msg.Append(message.NewText("请输入正确的B站UID"))
						PList.bot.Client.SendGroupMessage(groupId, msg)
						return
					}
					uid, _ := strconv.ParseInt(uidStr[0], 10, 64)
					if user, ok := bilipush.UserList[uid]; ok {
						if user.HasGroupById(groupId) {
							user.RemoveGroup(groupId)
							msg.Append(message.NewText(fmt.Sprintf("已从B站关注列表中移除 %s UID: %d", user.Uname, user.Mid)))
							if len(user.GroupList) == 0 {
								user.IsListed = false
								delete(bilipush.UserList, uid)
							}
							Config.RemoveUser(&BiliPusher{UID: user.Mid, GroupList: GroupList{groupId}})
						} else {
							msg.Append(message.NewText("该用户不在B站关注列表中"))
						}
					} else {
						msg.Append(message.NewText("该用户不在B站关注列表中"))
					}
					PList.bot.Client.SendGroupMessage(groupId, msg)
				} else if strings.HasPrefix(cmd, "/微博关注") {
					uidStr := regexp.MustCompile(`\d+`).FindAllString(cmd, -1)
					if len(uidStr) == 0 {
						msg.Append(message.NewText("请输入正确的微博ID"))
						PList.bot.Client.SendGroupMessage(groupId, msg)
						return
					}
					uid, _ := strconv.ParseInt(uidStr[0], 10, 64)
					if user, ok := weibo.WeiBoUserList[uid]; ok {
						if user.HasGroupById(groupId) {
							msg.Append(message.NewText(fmt.Sprintf("该用户 %s ID: %d 已经微博在关注列表中", user.Name, user.Id)))
							PList.bot.Client.SendGroupMessage(groupId, msg)
						} else {
							user.AddGroup(groupId)
							//msg.Append(message.NewText(fmt.Sprintf("已添加 %s ID: %d 到微博关注列表", user.Name, user.Id)))
							Config.AddUser(&WeiBoPusher{UID: user.Id, GroupList: GroupList{groupId}})
						}
					} else {
						user, err := weibo.NewWeiBoUser(uid, false, groupId)
						if err != nil {
							msg.Append(message.NewText("获取用户信息失败"))
							PList.bot.Client.SendGroupMessage(groupId, msg)
							return
						}
						go WeiBoStart(user, PList.bot)
						log.Printf("[微博推送] 用户 %s ID: %d 初始化成功！", user.Name, user.Id)
						//msg.Append(message.NewText(fmt.Sprintf("已添加 %s UID: %d 到关注列表", user.Name, user.Id)))
						Config.AddUser(&WeiBoPusher{UID: user.Id, GroupList: GroupList{groupId}})
						//PList.bot.Client.SendGroupMessage(groupId, msg)
					}
				} else if strings.HasPrefix(cmd, "/微博取关") {
					uidStr := regexp.MustCompile(`\d+`).FindAllString(cmd, -1)
					if len(uidStr) == 0 {
						msg.Append(message.NewText("请输入正确的微博ID"))
						PList.bot.Client.SendGroupMessage(groupId, msg)
						return
					}
					uid, _ := strconv.ParseInt(uidStr[0], 10, 64)
					if user, ok := weibo.WeiBoUserList[uid]; ok {
						if user.HasGroupById(groupId) {
							user.RemoveGroup(groupId)
							msg.Append(message.NewText(fmt.Sprintf("已从微博关注列表中移除 %s ID: %d", user.Name, user.Id)))
							if len(user.GroupList) == 0 {
								user.IsListed = false
								delete(weibo.WeiBoUserList, uid)
							}
							Config.RemoveUser(&WeiBoPusher{UID: user.Id, GroupList: GroupList{groupId}})
						} else {
							msg.Append(message.NewText("该用户不在微博关注列表中"))
						}
					} else {
						msg.Append(message.NewText("该用户不在微博关注列表中"))
					}
					PList.bot.Client.SendGroupMessage(groupId, msg)
				} else if strings.HasPrefix(cmd, "/抖音关注") {
					dyUrl := regexp.MustCompile(`https://v.douyin.com/\w+/`).FindAllString(cmd, -1)
					if len(dyUrl) < 1 {
						msg.Append(message.NewText("请输入正确的抖音分享链接"))
						PList.bot.Client.SendGroupMessage(groupId, msg)
						return
					}
					if user, ok := tiktok.TikTokUserList[dyUrl[0]]; ok {
						if user.HasGroupById(groupId) {
							msg.Append(message.NewText(fmt.Sprintf("该用户 %s\n主页:%s 已经抖音在关注列表中", user.Name, user.Url)))
							PList.bot.Client.SendGroupMessage(groupId, msg)
						} else {
							user.AddGroup(groupId)
							msg.Append(message.NewText(fmt.Sprintf("已添加 %s\n主页:%s 到抖音关注列表", user.Name, user.Url)))
							target := message.Source{
								SourceType: message.SourceGroup,
								PrimaryID:  1,
							}
							imgBinary, _ := tiktok.Download(user.Face)
							img := bytes.NewReader(imgBinary)
							imgMsg, _ := PList.bot.Client.UploadImage(target, img)
							msg.Append(imgMsg)
							Config.AddUser(&TikTokPusher{Url: user.Url, GroupList: GroupList{groupId}})
						}
						PList.bot.Client.SendGroupMessage(groupId, msg)
					} else {
						tkuser, err := tiktok.NewTikTokUser(dyUrl[0], false, groupId)
						if err != nil {
							msg.Append(message.NewText("获取用户信息失败"))
							PList.bot.Client.SendGroupMessage(groupId, msg)
							return
						}
						go TikTokStart(tkuser, PList.bot)
						Config.AddUser(&TikTokPusher{Url: dyUrl[0], GroupList: GroupList{groupId}})
						msg.Append(message.NewText(fmt.Sprintf("已添加 %s\n主页:%s 到抖音关注列表", tkuser.Name, tkuser.Url)))
						target := message.Source{
							SourceType: message.SourceGroup,
							PrimaryID:  1,
						}
						imgBinary, _ := tiktok.Download(tkuser.Face)
						img := bytes.NewReader(imgBinary)
						imgMsg, _ := PList.bot.Client.UploadImage(target, img)
						msg.Append(imgMsg)
						PList.bot.Client.SendGroupMessage(groupId, msg)
					}
				} else if strings.HasPrefix(cmd, "/抖音取关") {
					dyUrl := regexp.MustCompile(`https://v.douyin.com/\w+/`).FindAllString(cmd, -1)
					if len(dyUrl) < 1 {
						msg.Append(message.NewText("请输入正确的抖音分享链接"))
						PList.bot.Client.SendGroupMessage(groupId, msg)
						return
					}
					if user, ok := tiktok.TikTokUserList[dyUrl[0]]; ok {
						if user.HasGroupById(groupId) {
							user.RemoveGroup(groupId)
							msg.Append(message.NewText(fmt.Sprintf("已从抖音关注列表中移除 %s\n主页:%s", user.Name, user.Url)))
							if len(user.GroupList) == 0 {
								user.IsListed = false
								delete(tiktok.TikTokUserList, dyUrl[0])
							}
							Config.RemoveUser(&TikTokPusher{Url: dyUrl[0], GroupList: GroupList{groupId}})
						} else {
							msg.Append(message.NewText("该用户不在抖音关注列表中"))
						}
					} else {
						msg.Append(message.NewText("该用户不在抖音关注列表中"))
					}
					PList.bot.Client.SendGroupMessage(groupId, msg)
				}
			}
		},
	}
	PList.pluginList = append(PList.pluginList, *newPlugin)
	servers.RegisterCustom("plugins", func(bot *coolq.CQBot) {
		PList.bot = bot
		for _, plugin := range PList.pluginList {
			log.Println("插件:", plugin.Name, "已加载")
			bot.OnEventPush(plugin.Func)
		}
	})
	servers.RegisterCustom("bilipush", func(bot *coolq.CQBot) {
		err := bilipush.GrpcInit()
		if err != nil {
			log.Println("[B站推送] grpc 服务启动失败！")
			return
		}
		log.Println("[B站推送] grpc 服务启动成功！")
		for _, a := range Config.Bili {
			user, err := bilipush.NewUser(a.UID, true, a.GroupList...) // , 573546889, 686877263
			if err != nil {
				log.Warnf("[B站推送] 用户 %d 初始化失败！: %s", a.UID, err.Error())
				continue
			}
			log.Printf("[B站推送] 用户 %s UID: %d 初始化成功！", user.Uname, user.Mid)
			go BiliStart(user, bot)
		}
	})
	servers.RegisterCustom("tiktok", func(bot *coolq.CQBot) {
		for _, a := range Config.TikTok {
			tkuser, err := tiktok.NewTikTokUser(a.Url, true, a.GroupList...)
			if err != nil {
				log.Warnf("[抖音推送] 用户 %s 初始化失败！: %s", a.Url, err.Error())
				continue
			}
			go TikTokStart(tkuser, bot)
		}
	})
	servers.RegisterCustom("weibo", func(bot *coolq.CQBot) {
		for _, u := range Config.WeiBo {
			wbuser, err := weibo.NewWeiBoUser(u.UID, true, u.GroupList...)
			if err != nil {
				log.Warnln("[微博推送]", err)
				continue
			}
			go WeiBoStart(wbuser, bot)
			weibo.WeiBoUserList[u.UID] = wbuser
		}
	})
}
func BiliStart(u *bilipush.User, bot *coolq.CQBot) {
	go u.ListenUser()
	bilipush.UserList[u.Mid] = u
	for pushMsg := range u.MsgChan {
		msg := message.NewSendingMessage()
		target := message.Source{
			SourceType: message.SourceGroup,
			PrimaryID:  1,
		}
		switch pushMsg.Type {
		case "dynamic":
			msg.Append(message.NewText(pushMsg.Msg))
			if pushMsg.Img != nil {
				img := bytes.NewReader(pushMsg.Img)
				imgMsg, err := bot.Client.UploadImage(target, img)
				if err == nil {
					msg.Append(imgMsg)
				}
			} else {
				msg.Append(message.NewText(fmt.Sprintf("\n截图发生错误: %s\n", u.ErrorMessage)))
			}
			msg.Append(message.NewText(u.LastDyUrl))
		case "live":
			msg.Append(message.NewText(pushMsg.Msg))
			img := bytes.NewReader(pushMsg.Img)
			imgMsg, err := bot.Client.UploadImage(target, img)
			if err == nil {
				msg.Append(imgMsg)
			}
			if strings.Contains(pushMsg.Msg, "开播") {
				msg.Append(message.NewText(fmt.Sprintf("https://live.bilibili.com/%d", u.LivingInfo.Data[strconv.FormatInt(u.Mid, 10)].RoomId)))
			}

		case "other":
			msg.Append(message.NewText(pushMsg.Msg))
			img := bytes.NewReader(pushMsg.Img)
			imgMsg, err := bot.Client.UploadImage(target, img)
			if err == nil {
				msg.Append(imgMsg)
			}
		case "new":
			msg.Append(message.NewText(pushMsg.Msg))
			img := bytes.NewReader(pushMsg.Img)
			imgMsg, err := bot.Client.UploadImage(target, img)
			if err == nil {
				msg.Append(imgMsg)
			}
			bot.Client.SendGroupMessage(u.GroupList[len(u.GroupList)-1], msg)
			continue
		}
		for _, g := range u.GroupList {
			go bot.Client.SendGroupMessage(g, msg)
		}
	}
	log.Warnf("[B站推送] 用户 %s UID: %d 监听结束！", u.Uname, u.Mid)
}

func WeiBoStart(wbuser *weibo.WeiBoUser, bot *coolq.CQBot) {
	go wbuser.Listen()
	weibo.WeiBoUserList[wbuser.Id] = wbuser
	for pushMsg := range wbuser.MsgChan {
		msg := message.NewSendingMessage()
		target := message.Source{
			SourceType: message.SourceGroup,
			PrimaryID:  1,
		}
		msg.Append(message.NewText(pushMsg.Msg))
		if len(pushMsg.Img) > 0 {
			img := bytes.NewReader(pushMsg.Img)
			imgMsg, err := bot.Client.UploadImage(target, img)
			if err == nil {
				msg.Append(imgMsg)
			}
		}
		switch pushMsg.Type {
		case "weibo":
			msg.Append(message.NewText(strings.Split(wbuser.LastCardUrl, "?")[0]))
			for _, groupId := range wbuser.GroupList {
				bot.Client.SendGroupMessage(groupId, msg)
			}
		case "new":
			bot.Client.SendGroupMessage(wbuser.GroupList[len(wbuser.GroupList)-1], msg)
			continue
		}
	}
	log.Warnf("[微博推送] 用户 %s ID: %d 监听结束！", wbuser.Name, wbuser.Id)
}

func TikTokStart(tkuser *tiktok.TikTokUser, bot *coolq.CQBot) {
	go tkuser.Listen()
	tiktok.TikTokUserList[tkuser.Url] = tkuser
	for aweme := range tkuser.AwemeChan {
		msg0 := message.NewSendingMessage()
		msg0.Append(message.NewText(fmt.Sprintf("%s 抖音新作品：%s\nhttps://www.douyin.com/video/%s", tkuser.Name, aweme.Desc, aweme.AwemeId)))
		msg := message.NewSendingMessage()
		target := message.Source{
			SourceType: message.SourceGroup,
			PrimaryID:  1,
		}
		var videoBinary []byte
		var imgBinary []byte
		var err error
		var video *bytes.Reader
		var img *bytes.Reader
		for _, videoUrl := range aweme.Video.PlayAddrLowbr.UrlList {
			videoBinary, err = tiktok.Download(videoUrl)
			if err != nil {
				log.Println(err)
				continue
			}
			if len(videoBinary) > 0 {
				video = bytes.NewReader(videoBinary)
				break
			}
		}
		for _, imgUrl := range aweme.Video.Cover.UrlList {
			imgBinary, err = tiktok.Download(imgUrl)
			if err != nil {
				log.Println(err)
				continue
			}
			if len(imgBinary) > 0 {
				switch tiktok.CheckImgType(imgBinary) {
				case "jpg":
					img = bytes.NewReader(imgBinary)
				case "gif":
					imgBinary, err = tiktok.GIF2JPG(imgBinary)
					if err != nil {
						continue
					}
					img = bytes.NewReader(imgBinary)
				case "png":
					imgBinary, err = tiktok.PNG2JPG(imgBinary)
					if err != nil {
						continue
					}
					img = bytes.NewReader(imgBinary)
				}
			}
		}
		if video != nil && len(videoBinary) > 0 {
			videoMsg, err := bot.Client.UploadShortVideo(target, video, img, 1)
			if err != nil {
				log.Warnln(err)
				msg.Append(message.NewText(fmt.Sprintln("视频下载失败")))
				for _, groupId := range tkuser.GroupList {
					bot.Client.SendGroupMessage(groupId, msg)
				}
				continue
			}
			msg.Append(videoMsg)
		}
		for _, groupId := range tkuser.GroupList {
			bot.Client.SendGroupMessage(groupId, msg0)
			bot.Client.SendGroupMessage(groupId, msg)
		}
	}
	log.Warnf("[抖音推送] 用户 %s ID: %s 监听结束！", tkuser.Name, tkuser.Url)
}
