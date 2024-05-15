package main

import "backend/cmd"

func main() {
	server := cmd.ApiServer{}
	server.Run()
}
