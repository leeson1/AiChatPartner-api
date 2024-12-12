/*
 * @Author: LEESON
 * @Date: 2024-12-12 17:02:49
 */
package mysql

func GetUidByUserName(username string) int64 {
	var uid int64 = -1
	if username == "admin" {
		uid = 1001
	}
	return uid
}
