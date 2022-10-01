package main

import (
	"log"

	"github.com/rogerogers/config-manager/internal/config"
	api "github.com/rogerogers/config-manager/kitex_gen/api/config"
)

func main() {
	svr := api.NewServer(config.ConfigImpl{})
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
