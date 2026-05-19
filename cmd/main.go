package main

import "attendance-api/cmd/app"

func main() {
	app := app.App{}

	app.ConnectDb()
	app.Routes()
	app.Run()
}
