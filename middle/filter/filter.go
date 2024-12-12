/*
 * @Author: Leeson
 * @Date: 2024-12-12 19:25:06
 */
package filter

import (
	"net/http"
)

func FilterToken(w http.ResponseWriter, r *http.Request) error {
	// username := r.Header.Get("username")
	// uid := mysql.GetUidByUserName(username)

	// 这里想写一个过滤器，如果请求没有携带 token，就返回 401
	// token := r.Header.Get("token")
	// if token == "" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// 同时还需要拿这个token去查一下redis，看看是否存在，如果存在就续期
	// if !redis.Exist(token) {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	return nil
}
