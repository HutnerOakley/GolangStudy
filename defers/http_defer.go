package defers

/*

HTTP 服务、RPC 服务：**单个请求内部发生 panic（程序 bug），
不能让整个服务进程挂掉**，所以在请求入口加一层 recover，
捕获 panic，记录日志，终止这个请求，其他请求继续跑
*/

import (
	"log"
	"net/http"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if p := recover(); p != nil {
			log.Printf("panic: %v", p)
			http.Error(w, "server internal error", 500)
		}
	}()

	// 业务逻辑处理，万一panic：只崩这一个请求
}
