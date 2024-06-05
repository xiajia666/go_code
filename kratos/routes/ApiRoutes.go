package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/admin"
)

func ApiRoutes(r *gin.Engine) {
	r.Group("/")
	{
		r.GET("/apiget", admin.Upload)
		r.GET("/setsession", func(r *gin.Context) {
			// 设置cookies，六秒钟
			r.SetCookie("username", "我叫夏嘉，我是cookies", 6, "/", "localhost", false, false)
			session := sessions.Default(r) //设置session
			session.Options(sessions.Options{
				Path:     "",
				Domain:   "",
				MaxAge:   6, //设置过期时间
				Secure:   false,
				HttpOnly: false,
				SameSite: 0,
			})
			session.Set("username", "法外狂徒")
			_ = session.Save()
		})
	}
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置 session 中间件，参数 mysession，指的是 session 的名字，也是 cookie 的名字 // store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("myCookieSession", store))
	// 初始化基于 redis 的存储引擎
	// 参数说明:
	// 第 1 个参数 - redis 最大的空闲连接数
	// 第2个参数 - 数通信协议tcp或者udp
	// 第 3 个参数 - redis 地址, 格式，host:port
	// 第4个参数 - redis密码
	// 第 5 个参数 - session 加密密钥
	store_redis, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("myRedisSession", store_redis))
}
