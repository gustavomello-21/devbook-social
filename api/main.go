package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavomello-21/devbook/api/src/config"
	"github.com/gustavomello-21/devbook/api/src/router"
)

func main() {
	fmt.Println("Rodando a API!!!")

	config.Load()

	r := router.Routes()

	err := http.ListenAndServe(config.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
