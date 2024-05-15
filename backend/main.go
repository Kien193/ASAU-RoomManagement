package main

import "ASAU-user-api/cmd"

func main() {
	server := cmd.ApiServer{}
	server.Run()
}
