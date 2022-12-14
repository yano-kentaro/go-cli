//
//  root.go
//  go-cli
//
//  Created by Kentaro Yano on 2022/08/26.
//  Copyright © 2022 InterPark. All rights reserved.
//

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Url string `yaml:"url"`
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	// cfg, err := loadConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cfg.Url)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root.go init")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.go-cli.yaml)",
	)

	rootCmd.Flags().BoolP(
		"toggle",
		"t",
		false,
		"Help message for toggle",
	)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		filePath := os.Getenv("HOME") + "/data"
		viper.AddConfigPath(filePath)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-cli")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// 設定ファイル情報を取得
// 不要かもしれない
// func loadConfig() (*Config, error) {
// 	// viperの設定を呼び出し
// 	initConfig()

// 	var cfg Config
// 	if err := viper.Unmarshal(&cfg); err != nil {
// 		return nil, err
// 	}

// 	return &cfg, nil
// }
