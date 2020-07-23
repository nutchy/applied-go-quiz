package main

import (
	"flag"
	"fmt"

	"chayanon.t/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello")

	port := flag.String("port", "8080", "default port: 8080")
	state := flag.String("state", "sit", "default state: sit")

	flag.Parse()
	fmt.Println(port, state)

	//Initialize router api
	e := echo.New()
	if err := handler.InitRoutes(e); err != nil {
		fmt.Errorf("new routes error:", err)
		return
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
}
