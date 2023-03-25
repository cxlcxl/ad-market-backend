package validator

import (
	"github.com/gin-gonic/gin"
	"market/app/handlers"
	"market/app/response"
	"market/app/validator/v_data"
	"market/app/vars"
	"strconv"
)

func (v BsValidator) VAssetList(ctx *gin.Context) {
	var params v_data.VAssetList
	bindData(ctx, &params, (&handlers.Asset{}).AssetList)
}

func (v BsValidator) VAssetUpload(ctx *gin.Context) {
	f, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, "文件上传失败："+err.Error())
		return
	}

	// 超过系统设定的最大值：32M，tmpFile.Size 的单位是 bytes 和我们定义的文件单位 KB 比较，就需要将我们的单位*1024(即2的10次方)，一步到位就是 << 10
	limitSize := vars.YmlConfig.GetInt64("FileUploadSetting.Size")
	if f.Size > limitSize<<10 {
		response.Fail(ctx, "文件超过限制的 "+strconv.Itoa(int(limitSize))+" KB")
		return
	}

	if _, err := f.Open(); err != nil {
		response.Fail(ctx, "读取文件失败")
		return
	}

	(&handlers.Asset{}).AssetUpload(ctx)
}

func (v BsValidator) VAsset(ctx *gin.Context) {
	bindRouteData(ctx, "code", (&handlers.Asset{}).VAsset)
}

func (v BsValidator) VAssetDel(ctx *gin.Context) {
	bindRouteData(ctx, "id", (&handlers.Asset{}).AssetDel)
}
