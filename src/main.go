package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	pprof.Register(engine)
	//禁用控制台颜色
	router.Init(engine)

	hs := &http.Server{
		Addr:":8080",
		Handler:engine,
	}

	go hs.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println("pid:", os.Getpid())
	
	restart(hs)
}

func restart(hs *http.Server){
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <- sigs:
		hs.Shutdown(context.TODO())
	}
}