/*
 * @Author: Leeson
 * @Date: 2024-12-12 19:25:06
 */
package filter

import (
	"net/http"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FilterToken(w http.ResponseWriter, r *http.Request) error {

	// 这里写一个过滤器，如果请求没有携带 token，就返回 401
	if !needCheckToken(r) {
		return nil
	}

	// 如果过滤请求中是否有token
	// token := r.Header.Get("Authorization")
	// if token == "" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return nil
	// }

	// 读取请求体
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Failed to read request body", http.StatusBadRequest)
	// 	return err
	// }
	// defer r.Body.Close()

	// // 从body中解析出用户名
	// var data map[string]interface{}
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 	return err
	// }

	// // 提取字段（假设我们想要提取 "username" 和 "password" 字段）
	// username, usernameExists := data["username"].(string)
	// if !usernameExists {
	// 	http.Error(w, "Missing 'username' field", http.StatusBadRequest)
	// 	logx.Error("Missing 'username' field")
	// 	return err
	// }

	// //去mysql拿uid
	// uid := mysql.GetUidByUserName(username)

	// // 同时还需要拿这个token去查一下redis，看看是否存在，如果存在就续期
	// redisKey := strconv.Itoa(int(uid))
	// if _, err := redis.GetRedisClient().Hget(redisKey, "token"); err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	logx.Error("token not exist")
	// 	return err
	// } else {
	// 	// 续期
	// 	redis.GetRedisClient().Expire(redisKey, 3600)
	// }

	// 通过用户名去查找uin
	// 先去redis查,没有再去mysql查
	// userInfo, err := redis.GetRedisClient().Get(req.Username)
	// if userInfo == "" || err != nil {
	// 	logx.Info("[FilterToken] redis not exist, go to mysql err: ", err)
	// 	uid := mysql.GetUidByUserName(req.Username)
	// 	logx.Info("[FilterToken] uin: ", uid)
	// }

	// // 这里想写一个过滤器，如果请求没有携带 token，就返回 401
	// token := r.Header.Get("Authorization")
	// if token == "" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return nil
	// }

	// 同时还需要拿这个token去查一下redis，看看是否存在，如果存在就续期
	// if !redis.Exist(token) {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	return nil
}

func needCheckToken(r *http.Request) bool {
	if r.URL.Path == "/login" {
		return false
	}
	return true
}
