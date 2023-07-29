package utils

import (
	"regexp"
)

var (
	reSafeURL     *regexp.Regexp
	safeStringDic map[int32]struct{}
)

func init() {
	reSafeURL, _ = regexp.Compile(`^(http://|https://)?[\w.:/\-_#%]+\??[\w.\-_#%=&]*$`)
	safeStringDic = make(map[int32]struct{}, 0)

	const safeString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	for _, v := range safeString {
		safeStringDic[v] = struct{}{}
	}
}

func IsSafeURLValid(url string) bool {
	return reSafeURL.MatchString(url)
}

func IsSafeStringValid(str string, maxLen int) bool {
	if len(str) > maxLen {
		return false
	}
	for _, v := range str {
		if _, ok := safeStringDic[v]; !ok {
			return false // 包含非白名单字符
		}
	}
	return true
}
