package config

import (
	"context"
	"github.com/rogerogers/config-manager/kitex_gen/api"
)

type ConfigImpl struct {
}

func (c *ConfigImpl) CreateConfig(ctx context.Context, req *api.CreateConfigReq) (res *api.CreateConfigReply, err error) {
	return &api.CreateConfigReply{Code: 200}, nil
}
