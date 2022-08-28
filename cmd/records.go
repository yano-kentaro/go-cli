//
//  records.go
//  go-cli
//
//  Created by Kentaro Yano on 2022/08/26.
//  Copyright © 2022 InterPark. All rights reserved.
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// recordsCmd represents the records command
var recordsCmd = &cobra.Command{
	Use:   "records",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("records called")
	},
}

func init() {
	fmt.Println("records.go init")
	rootCmd.AddCommand(recordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
