package jpushclient

// Thirdpart 第三方平台
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

// SetContent 设定内容
func (tp *Thirdpart) SetContent(c string) {
	tp.Content = c

}

// SetTitle 设定标题
func (tp *Thirdpart) SetTitle(title string) {
	tp.Title = title
}

// SetURIActivity 设定 uri_activity
func (tp *Thirdpart) SetURIActivity(t string) {
	tp.URIActivity = t
}

// AddExtras 增加额外字段
func (tp *Thirdpart) AddExtras(key string, value interface{}) {
	if tp.Extras == nil {
		tp.Extras = make(map[string]interface{})
	}
	tp.Extras[key] = value
}
