//
//  main.go
//  go-cli
//
//  Created by Kentaro Yano on 2022/08/26.
//  Copyright © 2022 InterPark. All rights reserved.
//

package main

import (
	"fmt"
	"go-cli/cmd"
)

func init() {
	fmt.Println("main.go init")
}

func main() {
	cmd.Execute()
}
