package main

import (
	"fmt"

	"musubi/src/config"
)

func main() {
	var conf = config.New()
	fmt.Println(conf.BlockDir)

}
