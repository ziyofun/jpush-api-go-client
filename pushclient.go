package jpushclient

import (
	"encoding/base64"
	"errors"
	"strings"
)

const (
	// ValidateURI 验证推送地址
	ValidateURI = "https://api.jpush.cn/v3/push/validate"
	// PushURI 推送基础地址
	PushURI = "https://api.jpush.cn/v3/push"
	// PushFileURI 使用文件推送地址
	PushFileURI = "https://api.jpush.cn/v3/push/file"
	// ScheduleURI 定时推送地址
	ScheduleURI = "https://api.jpush.cn/v3/schedules"
	// FileURI 文件API地址
	FileURI = "https://api.jpush.cn/v3/files"
	// DeviceURI 设备API地址
	DeviceURI = "https://device.jpush.cn/v3"
	// ReportURI 获取报告地址
	ReportURI = "https://report.jpush.cn/v3"
	// Base64Table base64 字符表
	Base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(Base64Table)

// PushClient 推送 client 结构
type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseURI      string
	PushFileURI  string
	ScheduleURI  string
}

// NewPushClient 新建推送 client
func NewPushClient(secret, appKey string) *PushClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{
		MasterSecret: secret,
		AppKey:       appKey,
		AuthCode:     auth,
		BaseURI:      PushURI,
		PushFileURI:  PushFileURI,
		ScheduleURI:  ScheduleURI,
	}
	return pusher
}

// Send 发送数据
func (client *PushClient) Send(data []byte) (string, error) {
	return client.SendPushBytes(data)
}

// SendPushString 推送 Post 数据
func (client *PushClient) SendPushString(content string) (string, error) {
	ret, err := SendPostString(client.BaseURI, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

// SendPushBytes 推送 bytes 数据
func (client *PushClient) SendPushBytes(content []byte) (string, error) {
	//ret, err := SendPostBytes(client.BaseURI, content, client.AuthCode)
	ret, err := SendPostBytes2(client.BaseURI, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

// SetSchedule 设定定时推送
func (client *PushClient) SetSchedule(content []byte) (string, error) {

	ret, err := SendPostBytes2(client.ScheduleURI, content, client.AuthCode)
	if err != nil {
		return ret, err
	}

	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

// SendPushFile 设定文件推送
func (client *PushClient) SendPushFile(content []byte) (string, error) {
	ret, err := SendPostBytes2(client.PushFileURI, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}
