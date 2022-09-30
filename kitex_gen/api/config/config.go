// Code generated by Kitex v0.4.2. DO NOT EDIT.

package config

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	api "github.com/rogerogers/config-manager/kitex_gen/api"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return configServiceInfo
}

var configServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Config"
	handlerType := (*api.Config)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateConfig": kitex.NewMethodInfo(createConfigHandler, newCreateConfigArgs, newCreateConfigResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.2",
		Extra:           extra,
	}
	return svcInfo
}

func createConfigHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.CreateConfigReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.Config).CreateConfig(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateConfigArgs:
		success, err := handler.(api.Config).CreateConfig(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateConfigResult)
		realResult.Success = success
	}
	return nil
}
func newCreateConfigArgs() interface{} {
	return &CreateConfigArgs{}
}

func newCreateConfigResult() interface{} {
	return &CreateConfigResult{}
}

type CreateConfigArgs struct {
	Req *api.CreateConfigReq
}

func (p *CreateConfigArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.CreateConfigReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateConfigArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateConfigArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateConfigArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateConfigArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateConfigArgs) Unmarshal(in []byte) error {
	msg := new(api.CreateConfigReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateConfigArgs_Req_DEFAULT *api.CreateConfigReq

func (p *CreateConfigArgs) GetReq() *api.CreateConfigReq {
	if !p.IsSetReq() {
		return CreateConfigArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateConfigArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateConfigResult struct {
	Success *api.CreateConfigReply
}

var CreateConfigResult_Success_DEFAULT *api.CreateConfigReply

func (p *CreateConfigResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.CreateConfigReply)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateConfigResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateConfigResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateConfigResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateConfigResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateConfigResult) Unmarshal(in []byte) error {
	msg := new(api.CreateConfigReply)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateConfigResult) GetSuccess() *api.CreateConfigReply {
	if !p.IsSetSuccess() {
		return CreateConfigResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateConfigResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.CreateConfigReply)
}

func (p *CreateConfigResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateConfig(ctx context.Context, Req *api.CreateConfigReq) (r *api.CreateConfigReply, err error) {
	var _args CreateConfigArgs
	_args.Req = Req
	var _result CreateConfigResult
	if err = p.c.Call(ctx, "CreateConfig", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}