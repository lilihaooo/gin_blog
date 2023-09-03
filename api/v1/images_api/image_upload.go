package images_api

import (
	"blog_gin/global"
	"blog_gin/pkg/res"
	"blog_gin/service"
	"blog_gin/service/image_ser"
	"github.com/gin-gonic/gin"
	"os"
)

// ImageUpload 上传多图片, 返回图片的Url
func (ImagesApi) ImageUpload(c *gin.Context) {
	// 接收参数
	form, err := c.MultipartForm()
	if err != nil {
		res.Fail(c, res.INVALID_PARAMS, "无法解析上传的表单数据")
		return
	}
	files := form.File["images"]
	uploadDir := global.Config.Upload.Path
	// 检查并创建存储上传文件的目录
	if _, err = os.Stat(uploadDir); os.IsNotExist(err) {
		if err = os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			res.Fail(c, res.FAIL_OPER, "无法创建上传目录")
			return
		}
	}

	var resList []image_ser.FileUploadResponse
	for _, file := range files {
		resList = append(resList, service.AppService.ImageService.ImageUploadService(c, file))

	}
	res.OkWithData(c, resList)
}
