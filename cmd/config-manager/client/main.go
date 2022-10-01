package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rogerogers/config-manager/kitex_gen/api"
	"github.com/rogerogers/config-manager/kitex_gen/api/config"

	"github.com/cloudwego/kitex/client/callopt"
)

func main() {

	url := callopt.WithURL("http://localhost:8888")
	client, err := config.NewClient("config")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.CreateConfig(context.Background(), &api.CreateConfigReq{
		Namespace: "default",
		Group:     "config",
		Filename:  "config.yaml",
	}, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
