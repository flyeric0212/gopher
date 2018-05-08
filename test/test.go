package main

import (
    "fmt"
	"github.com/lisijie/webcron/app/libs"
)

func main() {
    password := "admin123"
	fmt.Println(libs.Md5([]byte(password)))
}
