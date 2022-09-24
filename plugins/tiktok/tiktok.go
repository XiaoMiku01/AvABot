package tiktok

import (
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type TikTokInfo struct {
	StatusCode int64   `json:"status_code"`
	AwemeList  []Aweme `json:"aweme_list"`
	MaxCursor  int64   `json:"max_cursor"`
	MinCursor  int64   `json:"min_cursor"`
	HasMore    bool    `json:"has_more"`
	Extra      struct {
		Now   int64  `json:"now"`
		Logid string `json:"logid"`
	} `json:"extra"`
}

type Aweme struct {
	AwemeId string `json:"aweme_id"`
	Desc    string `json:"desc"`
	Author  struct {
		Uid          string `json:"uid"`
		ShortId      string `json:"short_id"`
		Nickname     string `json:"nickname"`
		Signature    string `json:"signature"`
		AvatarLarger struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"avatar_larger"`
		AvatarThumb struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"avatar_thumb"`
		AvatarMedium struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"avatar_medium"`
		FollowStatus           int64       `json:"follow_status"`
		AwemeCount             int64       `json:"aweme_count"`
		FollowingCount         int64       `json:"following_count"`
		FollowerCount          int64       `json:"follower_count"`
		FavoritingCount        int64       `json:"favoriting_count"`
		TotalFavorited         string      `json:"total_favorited"`
		CustomVerify           string      `json:"custom_verify"`
		UniqueId               string      `json:"unique_id"`
		StoryOpen              bool        `json:"story_open"`
		WithCommerceEntry      bool        `json:"with_commerce_entry"`
		VerificationType       int64       `json:"verification_type"`
		EnterpriseVerifyReason string      `json:"enterprise_verify_reason"`
		IsAdFake               bool        `json:"is_ad_fake"`
		FollowersDetail        interface{} `json:"followers_detail"`
		Region                 string      `json:"region"`
		PlatformSyncInfo       interface{} `json:"platform_sync_info"`
		WithShopEntry          bool        `json:"with_shop_entry"`
		Secret                 int64       `json:"secret"`
		HasOrders              bool        `json:"has_orders"`
		Geofencing             interface{} `json:"geofencing"`
		VideoIcon              struct {
			Uri     string        `json:"uri"`
			UrlList []interface{} `json:"url_list"`
		} `json:"video_icon"`
		WithFusionShopEntry bool        `json:"with_fusion_shop_entry"`
		UserCanceled        bool        `json:"user_canceled"`
		IsGovMediaVip       bool        `json:"is_gov_media_vip"`
		PolicyVersion       interface{} `json:"policy_version"`
		SecUid              string      `json:"sec_uid"`
		Rate                int64       `json:"rate"`
		TypeLabel           []int64     `json:"type_label"`
		CardEntries         interface{} `json:"card_entries"`
		MixInfo             interface{} `json:"mix_info"`
		UserBadRate         int64       `json:"user_bad_rate"`
	} `json:"author"`
	ChaList interface{} `json:"cha_list"`
	Video   struct {
		PlayAddr struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"play_addr"`
		Cover struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"cover"`
		Height       int64 `json:"height"`
		Width        int64 `json:"width"`
		DynamicCover struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"dynamic_cover"`
		OriginCover struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"origin_cover"`
		Ratio        string `json:"ratio"`
		DownloadAddr struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"download_addr"`
		HasWatermark  bool `json:"has_watermark"`
		PlayAddrLowbr struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"play_addr_lowbr"`
		BitRate     interface{} `json:"bit_rate"`
		Duration    int64       `json:"duration"`
		Vid         string      `json:"vid"`
		IsLongVideo int64       `json:"is_long_video,omitempty"`
	} `json:"video"`
	Statistics struct {
		AwemeId      string `json:"aweme_id"`
		CommentCount int64  `json:"comment_count"`
		DiggCount    int64  `json:"digg_count"`
		PlayCount    int64  `json:"play_count"`
		ShareCount   int64  `json:"share_count"`
		ForwardCount int64  `json:"forward_count"`
	} `json:"statistics"`
	TextExtra []struct {
		Start       int64  `json:"start"`
		End         int64  `json:"end"`
		Type        int64  `json:"type"`
		HashtagName string `json:"hashtag_name"`
		HashtagId   int64  `json:"hashtag_id"`
		UserId      string `json:"user_id,omitempty"`
	} `json:"text_extra"`
	VideoLabels  interface{} `json:"video_labels"`
	AwemeType    int64       `json:"aweme_type"`
	ImageInfos   interface{} `json:"image_infos"`
	CommentList  interface{} `json:"comment_list"`
	Geofencing   interface{} `json:"geofencing"`
	VideoText    interface{} `json:"video_text"`
	LabelTopText interface{} `json:"label_top_text"`
	Promotions   interface{} `json:"promotions"`
	LongVideo    interface{} `json:"long_video"`
	Images       interface{} `json:"images"`
}

type AwemeList []Aweme

func (a Aweme) AwemeId2Uinit() uint64 {
	awemeId, _ := strconv.ParseUint(a.AwemeId, 10, 64)
	return awemeId
}

func (al AwemeList) MaxAwemeId() Aweme {
	var max Aweme
	for _, aweme := range al {
		if aweme.AwemeId2Uinit() > max.AwemeId2Uinit() {
			max = aweme
		}
	}
	return max
}

func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//println(len(body))
	return body, nil
}

func PNG2JPG(bin []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(bin))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func GIF2JPG(bin []byte) ([]byte, error) {
	img, err := gif.Decode(bytes.NewReader(bin))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var magicTable = map[string]string{
	"\xff\xd8\xff":      "jpg",
	"\x89PNG\r\n\x1a\n": "png",
	"GIF87a":            "gif",
	"GIF89a":            "gif",
}

func CheckImgType(incipit []byte) string {
	incipitStr := []byte(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(string(incipitStr), magic) {
			return mime
		}
	}
	return "unknown"
}
