package jpushclient

// Notice 通知结构
type Notice struct {
	Alert    string          `json:"alert,omitempty"`    // 通知内容
	Android  *AndroidNotice  `json:"android,omitempty"`  // 安卓通知内容
	IOS      *IOSNotice      `json:"ios,omitempty"`      // IOS通知内容
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"` // winphone 通知内容
}

// AndroidNotice 安卓通知结构
type AndroidNotice struct {
	Alert       string                 `json:"alert"`
	Title       string                 `json:"title,omitempty"`
	BuilderID   int                    `json:"builder_id,omitempty"`
	Intent      map[string]string      `json:"intent,omitempty"`
	URIActivity string                 `json:"uri_activity,omitempty"`
	URIAction   string                 `json:"uri_action,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// IOSNotice IOS 通知结构
type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            int                    `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

// WinPhoneNotice winphone 通知结构
type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

// SetAlert 设定通知文本
func (nt *Notice) SetAlert(alert string) {
	nt.Alert = alert
}

// SetAndroidNotice 设定安卓通知
func (nt *Notice) SetAndroidNotice(n *AndroidNotice) {
	nt.Android = n
}

// SetIOSNotice 设定iOS通知
func (nt *Notice) SetIOSNotice(n *IOSNotice) {
	nt.IOS = n
}

// SetWinPhoneNotice 设定 winphone 通知
func (nt *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	nt.WINPhone = n
}
