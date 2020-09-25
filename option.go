package jpushclient

// Option 推送配置项
type Option struct {
	SendNo            int                    `json:"sendno,omitempty"`
	TimeLive          int                    `json:"time_to_live,omitempty"`
	ApnsProduction    bool                   `json:"apns_production"`
	OverrideMsgID     int64                  `json:"override_msg_id,omitempty"`
	BigPushDuration   int                    `json:"big_push_duration,omitempty"`
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel,omitempty"`
}

// SetSendno 设定推送批次号
func (op *Option) SetSendno(no int) {
	op.SendNo = no
}

// SetTimelive 推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送。默认 86400 （1 天）
func (op *Option) SetTimelive(timelive int) {
	op.TimeLive = timelive
}

// SetOverrideMsgID 如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 msg_id 就会产生覆盖效果
func (op *Option) SetOverrideMsgID(id int64) {
	op.OverrideMsgID = id
}

// SetApns True 表示推送生产环境，False 表示要推送开发环境；如果不指定则为推送生产环境
func (op *Option) SetApns(apns bool) {
	op.ApnsProduction = apns
}

// SetBigPushDuration 又名缓慢推送，把原本尽可能快的推送速度，降低下来
func (op *Option) SetBigPushDuration(bigPushDuration int) {
	op.BigPushDuration = bigPushDuration
}
