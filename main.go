package main

import (
	"image-processor/framework/cmd"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	cmd.Run()
}
