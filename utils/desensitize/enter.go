package desensitize

import "strings"

// DesensitizeTel 电话号码脱敏
func DesensitizeTel(telNum string) string {
	if len(telNum) != 11 {
		return ""
	}
	return telNum[:3] + "****" + telNum[7:]
}

func DesensitizeEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	domain := parts[1]
	return parts[0][0:2] + "****@" + domain

}
