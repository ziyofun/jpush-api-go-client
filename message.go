package jpushclient

// Message 推送信息
type Message struct {
	Content     string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// SetContent 设定推送内容
func (mg *Message) SetContent(c string) {
	mg.Content = c

}

// SetTitle 设定推送标题
func (mg *Message) SetTitle(title string) {
	mg.Title = title
}

// SetContentType 设定内容类型
func (mg *Message) SetContentType(t string) {
	mg.ContentType = t
}

// AddExtras 增加额外信息
func (mg *Message) AddExtras(key string, value interface{}) {
	if mg.Extras == nil {
		mg.Extras = make(map[string]interface{})
	}
	mg.Extras[key] = value
}
