/*
 * @Author: Leeson
 * @Date: 2024-12-12 22:19:28
 */
package middle

import (
	"net/http"

	"github.com/rs/cors"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 在请求处理之前执行的逻辑

		corsConfig := cors.Default()
		corsMiddleware := func(next http.Handler) http.Handler {
			return corsConfig.Handler(next)
		}

		corsMiddleware(next).ServeHTTP(w, r)
		// 调用下一个处理器
		next(w, r)

		// 在请求处理之后执行的逻辑
	}
}
