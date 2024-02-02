package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer() // 启动服务器
		},
	}
	autoMigrateCmd = &cobra.Command{
		Use: "migrate-db",
		Run: func(cmd *cobra.Command, args []string) {
			autoMigrateDB() // 使用gorm同步数据库表结构
		},
	}
	createTableCmd = &cobra.Command{
		Use: "create-tables",
		Run: func(cmd *cobra.Command, args []string) {
			createDbTables() // 通过sql创建数据库表
		},
	}
	initDefaultCmd = &cobra.Command{
		Use: "init-default",
		Run: func(cmd *cobra.Command, args []string) {
			initDefault() // 初始化数据，添加默认app和tenants等信息
		},
	}
)

func init() {
	rootCmd.AddCommand(autoMigrateCmd)
	rootCmd.AddCommand(createTableCmd)
	rootCmd.AddCommand(initDefaultCmd)

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./deploy/dev.yaml", "config file (default is ./deploy/dev.yaml)")
}

func initConfig() {
	if err := InitConfig(cfgFile); err != nil {
		fmt.Println("init config err: ", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
