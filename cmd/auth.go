//
//  auth.go
//  go-cli
//
//  Created by Kentaro Yano on 2022/08/26.
//  Copyright © 2022 InterPark. All rights reserved.
//

package cmd

import (
	"fmt"
	lib_viper "go-cli/lib"

	"github.com/spf13/cobra"
)

type Authentication lib_viper.Authentication

// 構造体に"./auth.yaml"のデータをマッピングする関数
func authYamlMapping() *Authentication {
	// 認証情報ファイルの指定
	filePath := os.Getenv("HOME")
	lib_viper.LoadConfFile(".auth", "yaml", filePath)

	// 構造体にマッピング
	return lib_viper.Unmarshal[Authentication](".auth", "yaml")
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auth called")
	},
}

func init() {
	authData := authYamlMapping()
	fmt.Println(authData)
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
