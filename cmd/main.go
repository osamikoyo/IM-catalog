package main

import (
	"fmt"
	"github.com/osamikoyo/IM-catalog/internal/app"
)

func main() {
	application, err := app.Init()
	if err != nil {
		fmt.Println(err)
	}

	if err = application.Run();err != nil{
		fmt.Println(err)
	}
}