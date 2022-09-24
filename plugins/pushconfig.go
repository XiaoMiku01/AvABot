package plugins

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Pusher interface {
	AddGroup(groupIds ...int64) error // 添加推送群组
	RemoveGroup(groupId int64) error  // 移除推送群组
}

type GroupList []int64

func (g GroupList) Contains(groupId int64) bool {
	for _, v := range g {
		if v == groupId {
			return true
		}
	}
	return false
}

type BiliPusher struct {
	UID       int64     `yaml:"uid"`
	GroupList GroupList `yaml:"group_ids"`
}

func (b *BiliPusher) AddGroup(groupIds ...int64) error {
	for _, groupId := range groupIds {
		if !b.GroupList.Contains(groupId) {
			b.GroupList = append(b.GroupList, groupId)
		}
	}
	return nil
}

func (b *BiliPusher) RemoveGroup(groupId int64) error {
	if b.GroupList.Contains(groupId) {
		for i, v := range b.GroupList {
			if v == groupId {
				b.GroupList = append(b.GroupList[:i], b.GroupList[i+1:]...)
				break
			}
		}
		return nil
	} else {
		return errors.New("groupid not found")
	}
}

type TikTokPusher struct {
	Url       string    `yaml:"url"`
	GroupList GroupList `yaml:"group_ids"`
}

func (t *TikTokPusher) AddGroup(groupIds ...int64) error {
	for _, groupId := range groupIds {
		if !t.GroupList.Contains(groupId) {
			t.GroupList = append(t.GroupList, groupId)
		} else {
			return errors.New("groupid already exists")
		}
	}
	return nil
}

func (t *TikTokPusher) RemoveGroup(groupId int64) error {
	if t.GroupList.Contains(groupId) {
		for i, v := range t.GroupList {
			if v == groupId {
				t.GroupList = append(t.GroupList[:i], t.GroupList[i+1:]...)
				break
			}
		}
		return nil
	} else {
		return errors.New("groupid not found")
	}
}

type WeiBoPusher struct {
	UID       int64     `yaml:"uid"`
	GroupList GroupList `yaml:"group_ids"`
}

func (w *WeiBoPusher) AddGroup(groupIds ...int64) error {
	for _, groupId := range groupIds {
		if !w.GroupList.Contains(groupId) {
			w.GroupList = append(w.GroupList, groupId)
		} else {
			return errors.New("groupid already exists")
		}
	}
	return nil
}

func (w *WeiBoPusher) RemoveGroup(groupId int64) error {
	if w.GroupList.Contains(groupId) {
		for i, v := range w.GroupList {
			if v == groupId {
				w.GroupList = append(w.GroupList[:i], w.GroupList[i+1:]...)
				break
			}
		}
		return nil
	} else {
		return errors.New("groupid not found")
	}
}

type PushConfig struct {
	MasterID int64          `yaml:"master_id"`
	Bili     []BiliPusher   `yaml:"bili"`
	TikTok   []TikTokPusher `yaml:"tiktok"`
	WeiBo    []WeiBoPusher  `yaml:"weibo"`
}

func ReadConfig() (*PushConfig, error) {
	var config PushConfig
	configFile, err := os.Open("pusher_config.yaml")
	if err != nil {
		config.SaveConfig()
		return nil, errors.New("没有找到 pusher_config.yaml 文件，已自动生成，请修改 master_id 后重启")
	}
	defer configFile.Close()
	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (p *PushConfig) SaveConfig() error {
	configFile, err := os.Create("pusher_config.yaml")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer configFile.Close()
	err = yaml.NewEncoder(configFile).Encode(p)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}
func (p *PushConfig) AddUser(pusher Pusher) error {
	switch pusher := pusher.(type) {
	case *BiliPusher:
		for i, v := range p.Bili {
			if v.UID == pusher.UID {
				v.AddGroup(pusher.GroupList...)
				p.Bili[i] = v
				return p.SaveConfig()
			}
		}
		p.Bili = append(p.Bili, *pusher)
	case *TikTokPusher:
		for i, v := range p.TikTok {
			if v.Url == pusher.Url {
				v.AddGroup(pusher.GroupList...)
				p.TikTok[i] = v
				return p.SaveConfig()
			}
		}
		p.TikTok = append(p.TikTok, *pusher)
	case *WeiBoPusher:
		for i, v := range p.WeiBo {
			if v.UID == pusher.UID {
				v.AddGroup(pusher.GroupList...)
				p.WeiBo[i] = v
				return p.SaveConfig()
			}
		}
		p.WeiBo = append(p.WeiBo, *pusher)
	}
	return p.SaveConfig()
}
func (p *PushConfig) RemoveUser(pusher Pusher) error {
	switch pusher := pusher.(type) {
	case *BiliPusher:
		for i, v := range p.Bili {
			if v.UID == pusher.UID {
				temp := pusher
				for _, g := range temp.GroupList {
					v.RemoveGroup(g)
				}
				p.Bili[i] = v
				if len(p.Bili[i].GroupList) == 0 {
					p.Bili = append(p.Bili[:i], p.Bili[i+1:]...)
				}
				break
			}
		}
	case *TikTokPusher:
		for i, v := range p.TikTok {
			if v.Url == pusher.Url {
				temp := pusher
				for _, g := range temp.GroupList {
					v.RemoveGroup(g)
				}
				p.TikTok[i] = v
				if len(p.TikTok[i].GroupList) == 0 {
					p.TikTok = append(p.TikTok[:i], p.TikTok[i+1:]...)
				}
				break
			}
		}
	case *WeiBoPusher:
		for i, v := range p.WeiBo {
			if v.UID == pusher.UID {
				temp := pusher
				for _, g := range temp.GroupList {
					v.RemoveGroup(g)
				}
				p.WeiBo[i] = v
				if len(p.WeiBo[i].GroupList) == 0 {
					p.WeiBo = append(p.WeiBo[:i], p.WeiBo[i+1:]...)
				}
				break
			}
		}
	}
	return p.SaveConfig()
}

var Config *PushConfig

func init() {
	var err error
	Config, err = ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
