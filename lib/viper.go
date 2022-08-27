//
//  viper.go
//  go-cli
//
//  Created by Kentaro Yano on 2022/08/27.
//  Copyright © 2022 InterPark. All rights reserved.
//

package lib_viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// "./auth.yaml"のデータをマッピングする構造体
type Authentication struct {
	User []struct {
		Name     string
		ApiKey   string
		ApiToken string
	}
}

// ファイルを読み込む関数
func LoadConfFile(
	fileName string,
	fileType string,
	filePath string,
) {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		errMsg := "Failed to load " + fileName + "." + fileType
		fmt.Println(errMsg)
		os.Exit(0)
	}
}

