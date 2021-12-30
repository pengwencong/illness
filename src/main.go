package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go/router"
	"go/server"
)

func main() {

	err := server.InitRedis()
	if err != nil {
		fmt.Println(err)
		return
	}

	engine := gin.Default()
	engine.Use(TlsHandler())
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)
	//禁用控制台颜色
	router.Init(engine)

	err = engine.RunTLS(":80", "./runtime/tls/server.pem", "./runtime/tls/server.key")

	fmt.Println("listen err:", err)
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":80",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
