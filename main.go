package main

import (
	"context"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/basic"
	"project/middleware"
	_ "project/plugin/logrus"
	_ "project/plugin/mysql"
	_ "project/plugin/sms"
	"syscall"
	"time"
)

var logger = logrus.WithField("component", "http")

func main() {
	app := cli.NewApp()

	// base application info
	app.Name = "gst server"
	app.Author = "rzy"
	app.Version = "0.0.1"
	app.Copyright = "gst team reserved"
	app.Usage = "gst server"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./configs/config.toml",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Action = serve
	app.Run(os.Args)

}

func serve(c *cli.Context) error {
	basic.Init(c.String("config"))

	var (
		addr      = viper.GetString("webserver.addr")
		cert      = viper.GetString("webserver.certificates.cert")
		key       = viper.GetString("webserver.certificates.key")
		enableSSL = viper.GetBool("webserver.enable_ssl")
	)

	logrus.Infof("Web service addr: %s(enable ssl: %v)", addr, enableSSL)

	mux := startupService()
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         addr,
		Handler:      mux,
	}

	go func() {
		if enableSSL {
			log.Fatal(srv.ListenAndServeTLS(cert, key))
		} else {
			log.Fatal(srv.ListenAndServe())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

	return nil
}

func startupService() http.Handler {
	if !viper.GetBool("core.debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.MaxMultipartMemory = 1 << 20 //上传文件最大1M
	e.Use(
		gin.Recovery(),
		middleware.LogReq(logger),
		middleware.Cors(),
	)

	ginpprof.Wrap(e)
	route(e)

	return e
}
