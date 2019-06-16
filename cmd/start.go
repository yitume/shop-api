package cmd

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yitume/muses"
	"github.com/yitume/muses/pkg/cache/redis"
	"github.com/yitume/muses/pkg/database/mysql"
	"github.com/yitume/muses/pkg/server/gin"
	"github.com/yitume/muses/pkg/server/stat"
	"go.uber.org/zap"
	"syscall"

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
	if err := muses.Container(
		bootstrap.Arg.CfgFile,
		mysql.Register,
		redis.Register,
		gin.Register,
		stat.Register,
	); err != nil {
		panic(err)
	}
	// 配置初始化
	if err := bootstrap.InitConfig(bootstrap.Arg.CfgFile); err != nil {
		model.Logger.Panic(err.Error())
	}

	model.Init()
	service.Init()
	service.InitGen()
	// 主服务器
	endless.DefaultReadTimeOut = gin.Config().Muses.Server.Gin.ReadTimeout.Duration
	endless.DefaultWriteTimeOut = gin.Config().Muses.Server.Gin.WriteTimeout.Duration
	endless.DefaultMaxHeaderBytes = 100000000000000
	server := endless.NewServer(gin.Config().Muses.Server.Gin.Addr, router.InitApi())
	server.BeforeBegin = func(add string) {
		model.Logger.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}

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
