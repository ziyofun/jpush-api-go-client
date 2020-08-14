package jpushclient

import (
	"encoding/base64"
	"errors"
	"strings"
)

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_NAME_SSL = "https://api.jpush.cn/v3/push"
	PUSH_FILE_URL = "https://api.jpush.cn/v3/push/file"
	SCHEDULE_HOST_NAME = "https://api.jpush.cn/v3/schedules"
	BASE64_TABLE  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
	PushFileUrl  string
	ScheduleUrl  string
}

func NewPushClient(secret, appKey string) *PushClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{
		MasterSecret: secret,
		AppKey: appKey,
		AuthCode: auth,
		BaseUrl: HOST_NAME_SSL,
		PushFileUrl: PUSH_FILE_URL,
		ScheduleUrl: SCHEDULE_HOST_NAME,
	}
	return pusher
}

func (this *PushClient) Send(data []byte) (string, error) {
	return this.SendPushBytes(data)
}

func (this *PushClient) SendPushString(content string) (string, error) {
	ret, err := SendPostString(this.BaseUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (this *PushClient) SendPushBytes(content []byte) (string, error) {
	//ret, err := SendPostBytes(this.BaseUrl, content, this.AuthCode)
	ret, err := SendPostBytes2(this.BaseUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}


func (this *PushClient) SetSchedule(content []byte) (string, error) {
	
	ret , err := SendPostBytes2(this.ScheduleUrl, content, this.AuthCode )
	if err != nil {
		return ret, err
	}

	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (this *PushClient) SendPushFile(content []byte) (string, error) {
	ret, err := SendPostBytes2(this.PushFileUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}
