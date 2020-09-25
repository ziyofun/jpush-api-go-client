package jpushclient

const (
	tag    = "tag"             // 指定标签
	tagAnd = "tag_and"         // 指定标签‘与’条件
	alias  = "alias"           // 指定别名
	id     = "registration_id" // 指定RID
	fileID = "file_id"         // 指定文件ID
)

// Audience 可以指定的受众信息
type Audience struct {
	Object   interface{}
	audience map[string][]string
}

// All 设定推送目标为所有人
func (ad *Audience) All() {
	ad.Object = "all"
}

// SetID 设定推送ID
func (ad *Audience) SetID(ids []string) {
	ad.set(id, ids)
}

// SetTag 设定推送标签
func (ad *Audience) SetTag(tags []string) {
	ad.set(tag, tags)
}

// SetTagAnd 设定标签‘与’条件
func (ad *Audience) SetTagAnd(tags []string) {
	ad.set(tagAnd, tags)
}

// SetAlias 设定推送别名
func (ad *Audience) SetAlias(aliases []string) {
	ad.set(alias, aliases)
}

// SetFile 设定推送文件名
func (ad *Audience) SetFile(fileID string) {
	ad.Object = map[string]interface{}{
		"file": map[string]string{
			"file_id": fileID,
		},
	}
}

// set 为推送受众设定字段值
func (ad *Audience) set(key string, v []string) {
	if ad.Object == nil {
		ad.audience = map[string][]string{key: v}
		ad.Object = ad.audience
	}

	ad.audience[key] = v
}
