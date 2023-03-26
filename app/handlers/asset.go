package handlers

import (
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/response"
	serviceasset "market/app/service/asset"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
	"strconv"
)

type Asset struct{}

func (l *Asset) AssetList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VAssetList)
	offset := utils.GetPages(params.Page, params.PageSize)
	imgs, total, err := model.NewListenImg(vars.DBMysql).ListenImgList(params.Name, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": imgs})
}

func (l *Asset) VAsset(ctx *gin.Context, v string) {
	if len(v) != 32 {
		return
	}
	filePath := model.NewListenImg(vars.DBMysql).FindImgByCode(v)
	ctx.File(vars.BasePath + filePath)
}

func (l *Asset) AssetDel(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	err = model.NewListenImg(vars.DBMysql).DeleteById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *Asset) AssetUpload(ctx *gin.Context) {
	savePath := vars.BasePath + vars.YmlConfig.GetString("FileUploadSetting.UploadPath")
	if ok, file := serviceasset.Upload(ctx, savePath); ok {
		if err := model.NewListenImg(vars.DBMysql).ListenImgCreate(file); err == nil {
			response.Success(ctx, nil)
			return
		}
	}
	response.Fail(ctx, "文件上传失败")
}
