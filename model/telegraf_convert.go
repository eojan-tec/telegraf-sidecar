package model

import "github.com/eojan-tec/telegraf-sidecar/tfapi"

func FileInfoToPB(in FileInfo) *tfapi.FileInfo {
	return &tfapi.FileInfo{
		Path: in.Path,
	}
}
func FilesInfoToPB(in []FileInfo) []*tfapi.FileInfo {
	var pbFilesInfo []*tfapi.FileInfo
	for _, v := range in {
		pbFilesInfo = append(pbFilesInfo, FileInfoToPB(v))
	}
	return pbFilesInfo
}
