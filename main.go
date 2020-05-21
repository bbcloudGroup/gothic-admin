package main

import (
	"github.com/bbcloudGroup/gothic/bootstrap"
	"gothic-admin/boot"
)


func main() {

	app := boot.NewApp(bootstrap.GetArgs())
	app.Run()

}
