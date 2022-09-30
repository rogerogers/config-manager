package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rogerogers/config-manager/kitex_gen/api"
	"github.com/rogerogers/config-manager/kitex_gen/api/config"

	kclient "github.com/cloudwego/kitex/client"
)

func main() {
	client, err := config.NewClient("config", kclient.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.CreateConfig(context.Background(), &api.CreateConfigReq{
		Namespace: "default",
		Group:     "config",
		Filename:  "config.yaml",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
