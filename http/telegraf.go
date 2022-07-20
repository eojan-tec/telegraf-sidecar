package http

import (
	tfmodel "codeup.aliyun.com/xinnong/telegraf-sidecar/model"
	"codeup.aliyun.com/xinnong/telegraf-sidecar/tfapi"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/middleware"
	"net/http"
)

// GetFile
// @Summary 获取文件内容
// @Description 获取文件内容
// @Tags telegraf
// @Accept  json
// @Produce  json
// @Param path query string true "相对路径"
// @Success 200 {object} tfapi.ContentResponse
// @Router /api/telegraf/getFile [get]
func GetFile(c *gin.Context) {
	// 定义请求和响应
	req := &tfapi.PathRequest{}
	resp := &tfapi.ContentResponse{Code: model.Success}
	// 处理请求：绑定请求
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 验证请求：验证请求的数据是否合规
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 响应请求：获取文件内容
	content, err := tfmodel.GetFile(req.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	} else {
		resp.Content = content
	}
	c.JSON(http.StatusOK, resp)
}

// GetDir
// @Summary 获取文件目录下的文件路径
// @Description 获取文件目录下的文件路径
// @Tags telegraf
// @Accept  json
// @Produce  json
// @Param path query string true "相对路径"
// @Success 200 {object} tfapi.ContentResponse
// @Router /api/telegraf/getDir [get]
func GetDir(c *gin.Context) {
	// 定义请求和响应
	req := &tfapi.PathRequest{}
	resp := &tfapi.FilesResponse{Code: model.Success}
	// 处理请求：绑定请求
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 验证请求：验证请求的数据是否合规
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 响应请求：获取文件目录下的文件路径
	filesInfo, err := tfmodel.GetDir(req.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	} else {
		resp.Files = tfmodel.FilesInfoToPB(filesInfo)
	}
	c.JSON(http.StatusOK, resp)
}

// Update
// @Summary 更新文件内容
// @Description 更新文件内容
// @Tags telegraf
// @Accept  json
// @Produce  json
// @Param content body tfapi.ContentRequest true "文件和文件内容"
// @Success 200 {object} tfapi.CommonResponse
// @Router /api/telegraf/update [put]
func Update(c *gin.Context) {
	// 定义请求和响应
	req := &tfapi.ContentRequest{}
	resp := &tfapi.CommonResponse{Code: model.Success}
	// 处理请求：绑定请求
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 验证请求：验证请求的数据是否合规
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 响应请求：更新文件
	err = tfmodel.Update(req.Path, req.Content)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Delete
// @Summary 删除文件
// @Description 删除文件
// @Tags telegraf
// @Accept  json
// @Produce  json
// @Param path query string true "相对路径"
// @Success 200 {object} tfapi.CommonResponse
// @Router /api/telegraf/delete [delete]
func Delete(c *gin.Context) {
	// 定义请求和响应
	req := &tfapi.PathRequest{}
	resp := &tfapi.CommonResponse{Code: model.Success}
	// 处理请求：绑定请求
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 验证请求：验证请求的数据是否合规
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 响应请求：删除文件
	err = tfmodel.Delete(req.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Touch
// @Summary 更新文件的访问时间和修改时间
// @Description 更新文件的访问时间和修改时间
// @Tags telegraf
// @Accept  json
// @Produce  json
// @Param path body tfapi.PathRequest true "相对路径"
// @Success 200 {object} tfapi.CommonResponse
// @Router /api/telegraf/touch [post]
func Touch(c *gin.Context) {
	// 定义请求和响应
	req := &tfapi.PathRequest{}
	resp := &tfapi.CommonResponse{Code: model.Success}
	// 处理请求：绑定请求
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 验证请求：验证请求的数据是否合规
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 响应请求：更新文件的访问时间和修改时间
	err = tfmodel.Touch(req.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Exist
// @Summary 检查文件是否存在
// @Description 检查文件是否存在
// @Tags telegraf
// @Accept  json
// @Produce  json
// @Param path query string true "相对路径"
// @Success 200 {object} tfapi.CommonResponse
// @Router /api/telegraf/exist [get]
func Exist(c *gin.Context) {
	// 定义请求和响应
	req := &tfapi.PathRequest{}
	resp := &tfapi.CommonResponse{Code: model.Success}
	// 处理请求：绑定请求
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 验证请求：验证请求的数据是否合规
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	// 响应请求：检查文件是否存在
	err = tfmodel.Exist(req.Path)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
