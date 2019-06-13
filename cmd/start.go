package cmd

import (
	"fmt"
	"net/http"
	"syscall"
	"time"

	"github.com/yitume/shop-api/pkg/prom"

	"github.com/fvbock/endless"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yitume/caller"
	"github.com/yitume/caller/gorm"
	"github.com/yitume/caller/redigo"
	callerzap "github.com/yitume/caller/zap"
	"go.uber.org/zap"

	"github.com/yitume/shop-api/model"
	"github.com/yitume/shop-api/pkg/bootstrap"
	"github.com/yitume/shop-api/router"
	"github.com/yitume/shop-api/service"
)

// startCmd represents the hello command
var startCmd = &cobra.Command{
	Use:  "start",
	Long: `Starts shop-api server`,
	Run:  startFn,
}

func init() {
	startCmd.PersistentFlags().StringVar(&bootstrap.Arg.CfgFile, "config", "conf/conf.toml", "config file (default is $HOME/.cobra-example.yaml)")
	RootCmd.AddCommand(startCmd)
	cobra.OnInitialize(initConfig)
}

func startFn(cmd *cobra.Command, args []string) {
	err := caller.Init(
		bootstrap.Arg.CfgFile,
		gorm.New,
		redigo.New,
		callerzap.New,
	)
	if err != nil {
		panic(err)
	}

	// 配置初始化
	if err := bootstrap.InitConfig(bootstrap.Arg.CfgFile); err != nil {
		model.Logger.Panic(err.Error())
	}
	// 信息初始化到prometheus
	prom.AppBuildInfo.Set(bootstrap.Conf.Build.Version)

	model.Init()
	service.Init()
	service.InitGen()
	// 主服务器
	endless.DefaultReadTimeOut = bootstrap.Conf.ServerApi.ReadTimeout.Duration
	endless.DefaultWriteTimeOut = bootstrap.Conf.ServerApi.WriteTimeout.Duration
	endless.DefaultMaxHeaderBytes = bootstrap.Conf.ServerApi.MaxHeaderBytes
	server := endless.NewServer(bootstrap.Conf.ServerApi.Addr, router.InitApi())
	server.BeforeBegin = func(add string) {
		model.Logger.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}

	serverStats := &http.Server{
		Addr:         bootstrap.Conf.ServerStat.Addr,
		Handler:      router.InitStat(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 异步跑监控
	go func() {
		defer func() {
			serverStats.Close()
		}()
		if err := serverStats.ListenAndServe(); err != nil {
			model.Logger.Error("ServerApi err", zap.String("err", err.Error()))
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		model.Logger.Error("ServerApi err", zap.String("err", err.Error()))
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(bootstrap.Arg.CfgFile)
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
