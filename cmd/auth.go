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
	"os"

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
	Short: "Set up credentials.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("help called")
			os.Exit(0)
		}

		switch args[0] {
		case "list":
			fmt.Println("list called")
		case "add":
			fmt.Println("add called")
		case "update":
			fmt.Println("update called")
		case "remove":
			fmt.Println("remove called")
		default:
			fmt.Println("help called")
		}

		return nil
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
