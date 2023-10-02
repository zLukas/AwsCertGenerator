package main

import (
	"fmt"
	"os"

	"github.com/zLukas/CloudTools/src/cert-generator/cmd"
)

func main() {
	enviroment := os.Getenv("ENVIROMENT")
	if enviroment == "LAMBDA" {

		cmd.RunLambda()
	} else {
		fmt.Print("running locally")
		cmd.RunLocal()
	}
}
