package weibo

type WeiBoInfo struct {
	Ok   int64 `json:"ok"`
	Data struct {
		CardlistInfo struct {
			Containerid string `json:"containerid"`
			VP          int64  `json:"v_p"`
			ShowStyle   int64  `json:"show_style"`
			Total       int64  `json:"total"`
			SinceID     int64  `json:"since_id"`
		} `json:"cardlistInfo"`
		Cards []struct {
			CardType      int64  `json:"card_type"`
			ProfileTypeID string `json:"profile_type_id"`
			Itemid        string `json:"itemid"`
			Scheme        string `json:"scheme"`
			Mblog         struct {
				Visible struct {
					Type   int64 `json:"type"`
					ListID int64 `json:"list_id"`
				} `json:"visible"`
				CreatedAt                string        `json:"created_at"`
				ID                       string        `json:"id"`
				Mid                      string        `json:"mid"`
				CanEdit                  bool          `json:"can_edit"`
				ShowAdditionalIndication int64         `json:"show_additional_indication"`
				Text                     string        `json:"text"`
				Source                   string        `json:"source"`
				Favorited                bool          `json:"favorited"`
				PicIds                   []interface{} `json:"pic_ids"`
				IsPaid                   bool          `json:"is_paid"`
				MblogVipType             int64         `json:"mblog_vip_type"`
				User                     struct {
					ID                int64  `json:"id"`
					ScreenName        string `json:"screen_name"`
					ProfileImageURL   string `json:"profile_image_url"`
					ProfileURL        string `json:"profile_url"`
					StatusesCount     int64  `json:"statuses_count"`
					Verified          bool   `json:"verified"`
					VerifiedType      int64  `json:"verified_type"`
					VerifiedTypeExt   int64  `json:"verified_type_ext"`
					VerifiedReason    string `json:"verified_reason"`
					CloseBlueV        bool   `json:"close_blue_v"`
					Description       string `json:"description"`
					Gender            string `json:"gender"`
					Mbtype            int64  `json:"mbtype"`
					Urank             int64  `json:"urank"`
					Mbrank            int64  `json:"mbrank"`
					FollowMe          bool   `json:"follow_me"`
					Following         bool   `json:"following"`
					FollowCount       int64  `json:"follow_count"`
					FollowersCount    string `json:"followers_count"`
					FollowersCountStr string `json:"followers_count_str"`
					CoverImagePhone   string `json:"cover_image_phone"`
					AvatarHd          string `json:"avatar_hd"`
					Like              bool   `json:"like"`
					LikeMe            bool   `json:"like_me"`
					Badge             struct {
						UserNameCertificate int64 `json:"user_name_certificate"`
						PcNew               int64 `json:"pc_new"`
						PartyCardidState    int64 `json:"party_cardid_state"`
					} `json:"badge"`
				} `json:"user"`
				RetweetedStatus struct {
					Visible struct {
						Type   int64 `json:"type"`
						ListID int64 `json:"list_id"`
					} `json:"visible"`
					CreatedAt                string   `json:"created_at"`
					ID                       string   `json:"id"`
					Mid                      string   `json:"mid"`
					CanEdit                  bool     `json:"can_edit"`
					ShowAdditionalIndication int64    `json:"show_additional_indication"`
					Text                     string   `json:"text"`
					TextLength               int64    `json:"textLength"`
					Source                   string   `json:"source"`
					Favorited                bool     `json:"favorited"`
					PicIds                   []string `json:"pic_ids"`
					ThumbnailPic             string   `json:"thumbnail_pic"`
					BmiddlePic               string   `json:"bmiddle_pic"`
					OriginalPic              string   `json:"original_pic"`
					IsPaid                   bool     `json:"is_paid"`
					MblogVipType             int64    `json:"mblog_vip_type"`
					User                     struct {
						ID                int64  `json:"id"`
						ScreenName        string `json:"screen_name"`
						ProfileImageURL   string `json:"profile_image_url"`
						ProfileURL        string `json:"profile_url"`
						StatusesCount     int64  `json:"statuses_count"`
						Verified          bool   `json:"verified"`
						VerifiedType      int64  `json:"verified_type"`
						VerifiedTypeExt   int64  `json:"verified_type_ext"`
						VerifiedReason    string `json:"verified_reason"`
						CloseBlueV        bool   `json:"close_blue_v"`
						Description       string `json:"description"`
						Gender            string `json:"gender"`
						Mbtype            int64  `json:"mbtype"`
						Urank             int64  `json:"urank"`
						Mbrank            int64  `json:"mbrank"`
						FollowMe          bool   `json:"follow_me"`
						Following         bool   `json:"following"`
						FollowCount       int64  `json:"follow_count"`
						FollowersCount    string `json:"followers_count"`
						FollowersCountStr string `json:"followers_count_str"`
						CoverImagePhone   string `json:"cover_image_phone"`
						AvatarHd          string `json:"avatar_hd"`
						Like              bool   `json:"like"`
						LikeMe            bool   `json:"like_me"`
						Badge             struct {
							Enterprise           int64 `json:"enterprise"`
							GongyiLevel          int64 `json:"gongyi_level"`
							Dailv                int64 `json:"dailv"`
							Zongyiji             int64 `json:"zongyiji"`
							VipActivity2         int64 `json:"vip_activity2"`
							SelfMedia            int64 `json:"self_media"`
							Dzwbqlx2016          int64 `json:"dzwbqlx_2016"`
							FollowWhitelistVideo int64 `json:"follow_whitelist_video"`
							Travel2017           int64 `json:"travel_2017"`
							UserNameCertificate  int64 `json:"user_name_certificate"`
							Wenchuan10Th         int64 `json:"wenchuan_10th"`
							Qixi2018             int64 `json:"qixi_2018"`
							Meilizhongguo2018    int64 `json:"meilizhongguo_2018"`
							NationalDay2018      int64 `json:"national_day_2018"`
							Memorial2018         int64 `json:"memorial_2018"`
							Hongbaofei2019       int64 `json:"hongbaofei_2019"`
							Denglong2019         int64 `json:"denglong_2019"`
							Suishoupai2019       int64 `json:"suishoupai_2019"`
							Hongrenjie2019       int64 `json:"hongrenjie_2019"`
							China2019            int64 `json:"china_2019"`
							Hongkong2019         int64 `json:"hongkong_2019"`
							Rrgyj2019            int64 `json:"rrgyj_2019"`
							Shouhuan2019         int64 `json:"shouhuan_2019"`
							Weishi2019           int64 `json:"weishi_2019"`
							Starlight2019        int64 `json:"starlight_2019"`
							Gongjiri2019         int64 `json:"gongjiri_2019"`
							Macao2019            int64 `json:"macao_2019"`
							China20192           int64 `json:"china_2019_2"`
							Hongbao2020          int64 `json:"hongbao_2020"`
							Feiyan2020           int64 `json:"feiyan_2020"`
							PcNew                int64 `json:"pc_new"`
							School2020           int64 `json:"school_2020"`
							Hongrenjie2020       int64 `json:"hongrenjie_2020"`
							China2020            int64 `json:"china_2020"`
							Zhongcaoguan2021     int64 `json:"zhongcaoguan_2021"`
							Yuanlongping2021     int64 `json:"yuanlongping_2021"`
							Ylpshuidao2021       int64 `json:"ylpshuidao_2021"`
							Gaokao2021           int64 `json:"gaokao_2021"`
							PartyCardidState     int64 `json:"party_cardid_state"`
							Aoyun2021            int64 `json:"aoyun_2021"`
							Kaixue212021         int64 `json:"kaixue21_2021"`
							Nihaochuntian2022    int64 `json:"nihaochuntian_2022"`
							Shuidao2022          int64 `json:"shuidao_2022"`
							Hangmu2022           int64 `json:"hangmu_2022"`
							Guoqi2022            int64 `json:"guoqi_2022"`
							Gangqi2022           int64 `json:"gangqi_2022"`
							Zhongqiujie2022      int64 `json:"zhongqiujie_2022"`
							Kaixueji2022         int64 `json:"kaixueji_2022"`
						} `json:"badge"`
					} `json:"user"`
					RepostsCount         int64 `json:"reposts_count"`
					CommentsCount        int64 `json:"comments_count"`
					ReprintCmtCount      int64 `json:"reprint_cmt_count"`
					AttitudesCount       int64 `json:"attitudes_count"`
					PendingApprovalCount int64 `json:"pending_approval_count"`
					IsLongText           bool  `json:"isLongText"`
					Mlevel               int64 `json:"mlevel"`
					ExpireTime           int64 `json:"expire_time"`
					AdState              int64 `json:"ad_state"`
					DarwinTags           []struct {
						ObjectType    string      `json:"object_type"`
						ObjectID      string      `json:"object_id"`
						DisplayName   string      `json:"display_name"`
						EnterpriseUID interface{} `json:"enterprise_uid"`
						BdObjectType  string      `json:"bd_object_type"`
					} `json:"darwin_tags"`
					HotPage struct {
						Fid            string `json:"fid"`
						FeedDetailType int64  `json:"feed_detail_type"`
					} `json:"hot_page"`
					Mblogtype         int64  `json:"mblogtype"`
					Mark              string `json:"mark"`
					Rid               string `json:"rid"`
					ContentAuth       int64  `json:"content_auth"`
					CommentManageInfo struct {
						CommentPermissionType int64 `json:"comment_permission_type"`
						ApprovalCommentType   int64 `json:"approval_comment_type"`
						CommentSortType       int64 `json:"comment_sort_type"`
					} `json:"comment_manage_info"`
					PicNum          int64 `json:"pic_num"`
					NewCommentStyle int64 `json:"new_comment_style"`
					EditConfig      struct {
						Edited bool `json:"edited"`
					} `json:"edit_config"`
					PageInfo struct {
						Type       string `json:"type"`
						ObjectType int64  `json:"object_type"`
						PagePic    struct {
							URL string `json:"url"`
						} `json:"page_pic"`
						PageURL   string `json:"page_url"`
						PageTitle string `json:"page_title"`
						Content1  string `json:"content1"`
					} `json:"page_info"`
					Pics interface{} `json:"pics"`
					Bid  string      `json:"bid"`
				} `json:"retweeted_status"`
				RepostsCount         int64         `json:"reposts_count"`
				CommentsCount        int64         `json:"comments_count"`
				ReprintCmtCount      int64         `json:"reprint_cmt_count"`
				AttitudesCount       int64         `json:"attitudes_count"`
				PendingApprovalCount int64         `json:"pending_approval_count"`
				IsLongText           bool          `json:"isLongText"`
				Mlevel               int64         `json:"mlevel"`
				DarwinTags           []interface{} `json:"darwin_tags"`
				HotPage              struct {
					Fid            string `json:"fid"`
					FeedDetailType int64  `json:"feed_detail_type"`
				} `json:"hot_page"`
				Mblogtype             int64  `json:"mblogtype"`
				Mark                  string `json:"mark"`
				Rid                   string `json:"rid"`
				ExternSafe            int64  `json:"extern_safe"`
				NumberDisplayStrategy struct {
					ApplyScenarioFlag    int64  `json:"apply_scenario_flag"`
					DisplayTextMinNumber int64  `json:"display_text_min_number"`
					DisplayText          string `json:"display_text"`
				} `json:"number_display_strategy"`
				ContentAuth       int64 `json:"content_auth"`
				CommentManageInfo struct {
					CommentPermissionType int64 `json:"comment_permission_type"`
					ApprovalCommentType   int64 `json:"approval_comment_type"`
					CommentSortType       int64 `json:"comment_sort_type"`
				} `json:"comment_manage_info"`
				RepostType        int64  `json:"repost_type"`
				PicNum            int64  `json:"pic_num"`
				NewCommentStyle   int64  `json:"new_comment_style"`
				AbSwitcher        int64  `json:"ab_switcher"`
				RegionName        string `json:"region_name"`
				RegionOpt         int64  `json:"region_opt"`
				MblogMenuNewStyle int64  `json:"mblog_menu_new_style"`
				EditConfig        struct {
					Edited bool `json:"edited"`
				} `json:"edit_config"`
				RawText string `json:"raw_text"`
				Bid     string `json:"bid"`
			} `json:"mblog"`
		} `json:"cards"`
		Scheme      string `json:"scheme"`
		ShowAppTips int64  `json:"showAppTips"`
	} `json:"data"`
}
