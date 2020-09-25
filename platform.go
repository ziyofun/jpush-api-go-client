package jpushclient

import (
	"errors"
)

const (
	ios      = "ios"
	android  = "android"
	winphone = "winphone"
)

// Platform 推送平台
type Platform struct {
	Os     interface{}
	osArry []string
}

// All 指定为全平台推送
func (pf *Platform) All() {
	pf.Os = "all"
}

// Add 增加推送系统
func (pf *Platform) Add(os string) error {
	if pf.Os == nil {
		pf.osArry = make([]string, 0, 3)
	} else {
		switch pf.Os.(type) {
		case string:
			panic("platform is all")
		default:
		}
	}

	//判断是否重复
	for _, value := range pf.osArry {
		if os == value {
			return nil
		}
	}

	switch os {
	case ios:
		fallthrough
	case android:
		fallthrough
	case winphone:
		pf.osArry = append(pf.osArry, os)
		pf.Os = pf.osArry
	default:
		return errors.New("unknow platform")
	}

	return nil
}

// AddIOS 增加 IOS 平台推送
func (pf *Platform) AddIOS() {
	pf.Add(ios)
}

// AddAndrid 增加 Android 平台推送
func (pf *Platform) AddAndrid() {
	pf.Add(android)
}

// AddWinphone 增加 winphone 平台推送
func (pf *Platform) AddWinphone() {
	pf.Add(winphone)
}
