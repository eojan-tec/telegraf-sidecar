package provider

import (
	"context"
	tfmodel "github.com/eojan-tec/telegraf-sidecar/model"
	"github.com/eojan-tec/telegraf-sidecar/tfapi"
	"github.com/nooocode/pkg/model"
)

type TelegrafProvider struct {
	tfapi.UnimplementedTelegrafServer
}

func (u *TelegrafProvider) GetFile(ctx context.Context, in *tfapi.PathRequest) (*tfapi.ContentResponse, error) {
	resp := &tfapi.ContentResponse{Code: model.Success}
	content, err := tfmodel.GetFile(in.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Content = content
	}
	return resp, nil
}
func (u *TelegrafProvider) GetDir(ctx context.Context, in *tfapi.PathRequest) (*tfapi.FilesResponse, error) {
	resp := &tfapi.FilesResponse{Code: model.Success}
	filesInfo, err := tfmodel.GetDir(in.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Files = tfmodel.FilesInfoToPB(filesInfo)
	}
	return resp, nil
}
func (u *TelegrafProvider) Update(ctx context.Context, in *tfapi.ContentRequest) (*tfapi.CommonResponse, error) {
	resp := &tfapi.CommonResponse{Code: model.Success}
	err := tfmodel.Update(in.Path, in.Content)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
func (u *TelegrafProvider) Touch(ctx context.Context, in *tfapi.PathRequest) (*tfapi.CommonResponse, error) {
	resp := &tfapi.CommonResponse{Code: model.Success}
	err := tfmodel.Touch(in.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
func (u *TelegrafProvider) Delete(ctx context.Context, in *tfapi.PathRequest) (*tfapi.CommonResponse, error) {
	resp := &tfapi.CommonResponse{Code: model.Success}
	err := tfmodel.Delete(in.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
func (u *TelegrafProvider) Exist(ctx context.Context, in *tfapi.PathRequest) (*tfapi.CommonResponse, error) {
	resp := &tfapi.CommonResponse{Code: model.Success}
	err := tfmodel.Exist(in.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
