package cmd

import (
	"QuickAuth/biz"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	ExitCmd        = 1
	ExitInitGlobal = 2
	ExitServer     = 3
	ExitMigrate    = 4
	ExitExecSql    = 5
	ExitReadFile   = 6
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "QuickAuth is a oauth2 provider server.",
		Run: func(cmd *cobra.Command, args []string) {
			startServer() // 启动服务器
		},
	}
	autoMigrateCmd = &cobra.Command{
		Use:   "migrate-db",
		Short: "Auto migrate database by gorm.",
		Run: func(cmd *cobra.Command, args []string) {
			autoMigrateDB() // 使用gorm同步数据库表结构
		},
	}
	createTableCmd = &cobra.Command{
		Use:   "create-tables",
		Short: "Create db by sql.",
		Run: func(cmd *cobra.Command, args []string) {
			createDbTables() // 通过sql创建数据库表
		},
	}
	initDefaultCmd = &cobra.Command{
		Use:   "init-default",
		Short: "init default data.",
		Run: func(cmd *cobra.Command, args []string) {
			initDefault() // 初始化数据，添加默认app和tenants等信息
		},
	}
	versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("QuickAuth v1.0 -- HEAD") // 版本信息
		},
	}
)

func init() {
	rootCmd.AddCommand(autoMigrateCmd)
	rootCmd.AddCommand(createTableCmd)
	rootCmd.AddCommand(initDefaultCmd)
	rootCmd.AddCommand(versionCmd)

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./deploy/dev.yaml", "config file (default is ./deploy/dev.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
}

func initConfig() {
	if err := internal.InitConfig(cfgFile); err != nil {
		fmt.Println("init config err: ", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(ExitCmd)
	}
}
