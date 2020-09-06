package jpushclient

// "notification_3rd": {
// 	"content": "Hi,JPush",
// 	"title": "msg",
// 	"channel_id": "channel001",
// 	"uri_activity": "cn.jpush.android.ui.OpenClickActivity",
// 	"uri_action": "cn.jpush.android.intent.CONNECTION"
// 	"badge_add_num": 1
// 	"badge_class": "com.test.badge.MainActivity"
// 	"sound": "sound"
// 	"extras":{
// 			"news_id" : 134,
// 			"my_key" : "a value"
// 	}
// }

type Thirdpart struct {
	Content     string                 `json:"content"`
	Title       string                 `json:"title,omitempty"`
	ChannelID   string                 `json:"channel_id,omitempty"`
	URIActivity string                 `json:"uri_activity,omitempty"`
	URIAction   string                 `json:"uri_action,omitempty"`
	BadgeAddNum int                    `json:"badge_add_num,omitempty"`
	BadgeClass  string                 `json:"badge_class,omitempty"`
	Sound       string                 `json:"sound,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (this *Thirdpart) SetContent(c string) {
	this.Content = c

}

func (this *Thirdpart) SetTitle(title string) {
	this.Title = title
}

func (this *Thirdpart) SetURIActivity(t string) {
	this.URIActivity = t
}

func (this *Thirdpart) AddExtras(key string, value interface{}) {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
}
