package serviceasset

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/utils"
	"market/app/vars"
	"os"
	"path"
	"strings"
	"time"
)

func Upload(ctx *gin.Context, savePath string) (ok bool, lessonImg *model.LessonImg) {
	newSavePath, newReturnPath := generateYearMonthPath(savePath)
	// 获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	file, _ := ctx.FormFile("file") // file 使用默认的上传名
	// 保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	var saveErr error
	saveFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), file.Filename)
	saveName := utils.MD5(saveFileName)
	saveFileName = saveName + path.Ext(saveFileName)

	if saveErr = ctx.SaveUploadedFile(file, newSavePath+saveFileName); saveErr == nil {
		// 上传成功,返回资源的相对路径，这里请根据实际返回绝对路径或者相对路径
		img := &model.LessonImg{
			ImgUrl: strings.ReplaceAll(newReturnPath+saveFileName, vars.BasePath, ""),
			FCode:  saveName,
			Name:   file.Filename,
			State:  1,
		}
		return true, img
	}

	return false, nil
}

// 文件上传可以设置按照 xxx年-xx月 格式存储
func generateYearMonthPath(savePathPre string) (string, string) {
	returnPath := vars.BasePath + vars.YmlConfig.GetString("FileUploadSetting.ReturnPath")
	curYearMonth := time.Now().Format("200601")
	newSavePathPre := savePathPre + curYearMonth
	newReturnPathPre := returnPath + curYearMonth
	// 相关路径不存在，创建目录
	if _, err := os.Stat(newSavePathPre); err != nil {
		if err = os.MkdirAll(newSavePathPre, os.ModePerm); err != nil {
			return "", ""
		}
	}
	return newSavePathPre + "/", newReturnPathPre + "/"
}
