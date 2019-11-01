package util

import "strings"

func GetBetween(text, start, end string) (data []string) {
	for k, v := range strings.Split(text, start) {
		if k > 0 {
			data = append(data, strings.Split(v, end)[0])
		}
	}
	return
}
func InArray(arr []string, str string) bool {
	ret := false
	for _, val := range arr {
		if str == val {
			ret = true
		}
	}
	return ret
}
func Unique(m []string) []string {
	s := make([]string, 0)
	sMap := make(map[string]string)
	for _, value := range m {
		if value != "" {
			//计算map长度
			length := len(sMap)
			sMap[value] = "1"
			//比较map长度, 如果map长度不相等， 说明key不存在
			if len(sMap) != length {
				s = append(s, value)
			}
		}
	}
	return s
}
