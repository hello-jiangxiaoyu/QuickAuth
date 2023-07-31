package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "QuickAuth is a oauth2 provider server.",
		Long:  `QuickAuth is a oauth2 provider server.`,
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
	autoMigrateCmd = &cobra.Command{
		Use:   "migrate db",
		Short: "Auto migrate database by gorm.",
		Long:  `Auto migrate database by gorm.`,
		Run: func(cmd *cobra.Command, args []string) {
			autoMigrateDB()
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of QuickAuth",
		Long:  `All software has versions. This is QuickAuth's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("QuickAuth Static Site Generator v1.0 -- HEAD")
		},
	}
)

func init() {
	rootCmd.AddCommand(autoMigrateCmd)
	rootCmd.AddCommand(versionCmd)

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./deploy/dev.yaml", "config file (default is ./deploy/dev.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
}

func initConfig() {
	fmt.Println("init config: ", cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
