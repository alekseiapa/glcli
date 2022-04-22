package main

import "github.com/alekseiapa/glcli/cmd"

var version = "dev"

func main() {
	cmd.Version = version
	cmd.Execute()
}
