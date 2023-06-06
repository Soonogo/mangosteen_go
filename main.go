package main

import (
	"mangosteen/cmd"
	"mangosteen/config"
)

func main() {
	config.LoadAppConfig()

	cmd.Run()
}
