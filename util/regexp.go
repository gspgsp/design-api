package util

import "regexp"

/**
jwt 匹配过期验证信息
 */
func JwtExpireValidReg(string string) bool {
	re := regexp.MustCompile("token is expired by*")
	if re.MatchString(string) {
		return true
	}

	return false
}
