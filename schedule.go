// Package jpushclient 定时模块相关功能支持
// http://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule/
package jpushclient

import (
	"encoding/json"
)

// Single 表示定时任务
type Single struct {
	Time string `json:"time"` // time 为必选项且格式为 "yyyy-MM-dd HH:mm:ss"，定时任务的最晚时间不能超过一年
}

// Periodical 表示定期任务
type Periodical struct {
	Start     string   `json:"start"`           // 表示定期任务有效起始时间，格式与必须严格为: "yyyy-MM-dd HH:mm:ss"，且时间必须为 24 小时制
	End       string   `json:"end"`             // 表示定期任务的过期时间，格式同上。定期任务最大跨度不能超过一年
	Time      string   `json:"time"`            // 表示触发定期任务的定期执行时间，格式严格为: "HH:mm:ss"，且为 24 小时制
	TimeUnit  string   `json:"time_unit"`       // 表示定期任务的执行的最小时间单位有："day"、"week" 及 "month" 三种。大小写不敏感
	Frequency int      `json:"frequency"`       // 与 time_unit 的乘积共同表示的定期任务的执行周期，如 time_unit 为 day，frequency 为 2，则表示每两天触发一次推送，目前支持的最大值为 100
	Point     []string `json:"point,omitempty"` // 一个列表，此项与 time_unit 相对应
}

// Trigger 触发条件，支持定期和定时
type Trigger struct {
	Single     *Single     `json:"single,omitempty"`
	Periodical *Periodical `json:"periodical,omitempty"`
}

// Schedule 定时配置
type Schedule struct {
	Cid     string   `json:"cid,omitempty"` // 防止重复发送的ID
	Name    string   `json:"name"`          // 表示定时任务的名字
	Enabled bool     `json:"enabled"`       // 表示任务当前状态，布尔值，必须为 true 或 false，创建任务时必须为 true
	Trigger *Trigger `json:"trigger"`       // 表示 schedule 任务的触发条件，当前只支持定时任务（single）与定期任务（periodical）
	Push    *PayLoad `json:"push"`          // push消息结构
}

// SetCid 设定 cid
func (sche *Schedule) SetCid(cid string) {
	sche.Cid = cid
}

// SetName 设定 name
func (sche *Schedule) SetName(name string) {
	sche.Name = name
}

// SetEnabled 设定当前任务状态
func (sche *Schedule) SetEnabled(enabled bool) {
	sche.Enabled = enabled
}

// SetPayload 设定定时推送内容
func (sche *Schedule) SetPayload(payload *PayLoad) {
	sche.Push = payload
}

// SetSingleSchedule 为定时推送设定时间
func (sche *Schedule) SetSingleSchedule(time string) {
	single := &Single{Time: time}

	if sche.Trigger == nil {
		trigger := &Trigger{Single: single}
		sche.Trigger = trigger
	} else {
		sche.Trigger.Single = single
	}

}

// SetPeriodicalSchedule 为定期推送设定定期规则
func (sche *Schedule) SetPeriodicalSchedule(periodical *Periodical) {

	if sche.Trigger == nil {
		trigger := &Trigger{Periodical: periodical}
		sche.Trigger = trigger
	} else {
		sche.Trigger.Periodical = periodical
	}
}

// ToBytes 将定时任务转换为 bytes 数据
func (sche *Schedule) ToBytes() ([]byte, error) {
	content, err := json.Marshal(sche)
	if err != nil {
		return nil, err
	}
	return content, nil
}
