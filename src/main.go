package main

import (
	"fmt"
)

func main() {
	fmt.Println("start app...")
	init, err := NewInitialize()
	if err != nil {
		return
	}
	e := RouterSetting(init.Controller)
	e.Logger.Fatal(e.Start(":80"))
}
