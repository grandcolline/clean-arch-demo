package main

import (
	"github.com/grandcolline/clean-arch-demo/driver"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"
)

func main() {
	defer mysql.CloseConn()
	driver.Router.Run()
}
