package image_ser

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/plugins/qi_niu"
	"blog_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type FileUploadResponse struct {
	File      string `json:"file"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (ImageService) ImageUploadService(c *gin.Context, file *multipart.FileHeader) (apiRes FileUploadResponse) {
	apiRes.File = file.Filename
	ext := strings.ToLower(filepath.Ext(file.Filename))
	// 检查文件类型是否在白名单中
	ok := utils.InListStr(ext, global.Config.WhiteList)
	if !ok {
		apiRes.IsSuccess = false
		apiRes.Msg = "非法文件"
		return
	}
	// 判断是否超出指定大小
	fileSizeInMB := float64(file.Size) / (1024 * 1024) // file.Size单位为字节将其转为MB
	formattedSize := fmt.Sprintf("%.2f", fileSizeInMB) //保留2位小数
	if fileSizeInMB > global.Config.Upload.Size {

		apiRes.IsSuccess = false
		apiRes.Msg = fmt.Sprintf("图片不能大于%.2fMB, 当前文件大小为%sMB", global.Config.Upload.Size, formattedSize)
		return
	}

	// 读取文件并将文件Md5加密(唯一, 同一文件的Md5值是相同的)
	fileObj, err := file.Open()
	if err != nil {
		global.Logrus.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	// 去数据库中查询这个文件是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {

		apiRes.IsSuccess = false
		apiRes.Msg = fmt.Sprintf("图片已存在")
		return
	}

	if global.Config.QiNiu.Enable {
		// 上传七牛
		filename := fmt.Sprintf("%s%s", "blog", file.Filename)
		path, err := qi_niu.UploadImage(byteData, filename, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Logrus.Error(err)
			apiRes.IsSuccess = false
			apiRes.Msg = err.Error()
			return
		}
		apiRes.File = path
		apiRes.IsSuccess = true
		apiRes.Msg = fmt.Sprintf("图片上传成功")
		// 上传成功写入数据库
		global.DB.Create(&models.BannerModel{
			Path: path,
			Hash: imageHash,
			Name: filename,
			Type: ctype.QiNiu,
		})
		return
	}

	// 上传本地
	uploadPath := filepath.Join(global.Config.Upload.Path, file.Filename)
	if err = c.SaveUploadedFile(file, uploadPath); err != nil {
		apiRes.File = file.Filename
		apiRes.IsSuccess = false
		apiRes.Msg = fmt.Sprintf("图片保存失败")

	}
	apiRes.IsSuccess = true
	apiRes.Msg = fmt.Sprintf("图片上传成功")
	// 上传成功写入数据库
	global.DB.Create(&models.BannerModel{
		Path: uploadPath,
		Hash: imageHash,
		Name: file.Filename,
		Type: ctype.Local,
	})
	return
}
