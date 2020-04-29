package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var CFG *ini.File


// Setup 初始化配置
func Setup(cfgPath string) {
	cfg, err := ini.Load(cfgPath)
	if err != nil {
		panic(fmt.Sprintf("no such config: %s", cfgPath))
	}
	CFG = cfg

}

// Get 获取字符串
func Get(sectionKey string) string {
	section, key := getSectionAndKey(sectionKey)
	return CFG.Section(section).Key(key).String()
}

// GetInt 获取整型
func GetInt(sectionKey string) int {
	section, key := getSectionAndKey(sectionKey)
	return CFG.Section(section).Key(key).MustInt(-999)
}


// getSectionAndKey 分割带区和key的值
// sectionKey 格式（section.key)
// 不带分区的话就直接传key即可
func getSectionAndKey(sectionKey string) (section, key string) {
	s := strings.Split(sectionKey, ".")
	switch len(s) {
	case 2:
		section = s[0]
		key = s[1]
	case 1:
		key = s[0]
	default:
		// todo 这里直接panic是否合适？
		panic(fmt.Sprintf("wrong sectionKey: %s", sectionKey))
	}
	return
}
