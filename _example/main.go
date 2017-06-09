package main

import (
	"fmt"

	"github.com/digineo/goldflags"
)

const appName = "demonstator"

func main() {
	fmt.Println(goldflags.Banner(appName))
}
