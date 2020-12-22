package main

import (
	_ "201221ManyPrograms/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
    fmt.Println("hello world")
}

