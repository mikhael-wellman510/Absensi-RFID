package main

import "attendance-api/cmd/app"

func main() {
	apps := app.App{}

	apps.ConnectDb()
	apps.Routes()
	apps.Run()
}
