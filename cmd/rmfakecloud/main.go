package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddvk/rmfakecloud/internal/app"
	"github.com/ddvk/rmfakecloud/internal/config"
	"github.com/ddvk/rmfakecloud/internal/storage/fs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var version string

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Println("Version: ", version)
		fmt.Printf(`
Environment Variables:
General:
%s	Log verbosity level (debug, info, warn) (default: info)
%s		Port (default: %s)
%s		Local storage folder (default: %s)
%s	Url the tablet can resolve (default: http://hostname:port)

email sending, smtp:
%s
%s
%s
%s	don't check the server certificate (not recommended)
%s	custom HELO (if your email server needs it)
%s	override the email's From:

myScript hwr (needs a developer account):
%s
%s
`,
			config.EnvLogLevel,
			config.EnvPort,
			config.DefaultPort,
			config.EnvDataDir,
			config.DefaultDataDir,
			config.EnvStorageURL,

			config.EnvSmtpServer,
			config.EnvSmtpUsername,
			config.EnvSmtpPassword,
			config.EnvSmtpInsecureTLS,
			config.EnvSmtpHelo,
			config.EnvSmtpFrom,

			config.EnvHwrApplicationKey,
			config.EnvHwrHmac,
		)
	}
	flag.Parse()
	fmt.Println("run with -h for all available env variables")

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.TextFormatter{})

	if lvl, err := log.ParseLevel(os.Getenv(config.EnvLogLevel)); err == nil {
		fmt.Println("Log level:", lvl)
		logger.SetLevel(lvl)
	}
	cfg := config.FromEnv()

	log.Println("Version: ", version)
	// configs
	log.Println("Documents will be saved in:", cfg.DataDir)
	log.Println("Url the device should use:", cfg.StorageURL)
	log.Println("Listening on port:", cfg.Port)

	fsStorage := &fs.Storage{
		Cfg: *cfg,
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	gin.DefaultWriter = logger.Writer()

	a := app.NewApp(cfg, fsStorage, fsStorage)
	go a.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Stopping the service...")
	a.Stop()
	log.Println("Stopped")
}
