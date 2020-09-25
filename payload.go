package jpushclient

import (
	"encoding/json"
)

// PayLoad 整个推送体
type PayLoad struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Thirdpart    interface{} `json:"notification_3rd,omitempty"`
	Options      *Option     `json:"options,omitempty"`
}

// NewPushPayLoad 创建一个新推送体
func NewPushPayLoad() *PayLoad {
	pl := &PayLoad{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}

// SetPlatform 设定推送平台
func (py *PayLoad) SetPlatform(pf *Platform) {
	py.Platform = pf.Os
}

// SetAudience 设定推送受众
func (py *PayLoad) SetAudience(ad *Audience) {
	py.Audience = ad.Object
}

// SetOptions 设定推送配置项
func (py *PayLoad) SetOptions(o *Option) {
	py.Options = o
}

// SetMessage 设定消息
func (py *PayLoad) SetMessage(m *Message) {
	py.Message = m
}

// SetThirdpart 设定厂商推送
func (py *PayLoad) SetThirdpart(t *Thirdpart) {
	py.Thirdpart = t
}

// SetNotice 设定通知内容
func (py *PayLoad) SetNotice(notice *Notice) {
	py.Notification = notice
}

// ToBytes 将发送内容转换成 []bytes
func (py *PayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(py)
	if err != nil {
		return nil, err
	}
	return content, nil
}
