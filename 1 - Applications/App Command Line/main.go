package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello There - Start Application ...")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}