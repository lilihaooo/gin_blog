package images_api

import (
	"blog_gin/global"
	"blog_gin/models"
	"blog_gin/models/ctype"
	"blog_gin/pkg/res"
	"blog_gin/plugins/qi_niu"
	"blog_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type fileUploadResponse struct {
	File      string `json:"file"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

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
	var resList []fileUploadResponse

	for _, file := range files {
		// 检查文件类型是否在白名单中
		ext := strings.ToLower(filepath.Ext(file.Filename))
		allowed := false
		whiteImageList := global.Config.WhiteList
		for _, allowedExt := range whiteImageList {
			if ext == allowedExt {
				allowed = true
				break
			}
		}
		if !allowed {
			resList = append(resList, fileUploadResponse{
				File:      file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("非法文件"),
			})
			continue
		}
		// 判断是否超出指定大小
		fileSizeInMB := float64(file.Size) / (1024 * 1024) // file.Size单位为字节将其转为MB
		formattedSize := fmt.Sprintf("%.2f", fileSizeInMB) //保留2位小数
		if fileSizeInMB > global.Config.Upload.Size {
			resList = append(resList, fileUploadResponse{
				File:      file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片不能大于%.2fMB, 当前文件大小为%sMB", global.Config.Upload.Size, formattedSize),
			})
			continue
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
			resList = append(resList, fileUploadResponse{
				File:      bannerModel.Path,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片已存在"),
			})
			continue
		}
		filename := fmt.Sprintf("%s%s", "image_", file.Filename)
		uploadPath := filepath.Join(uploadDir, filename)

		// 上传七牛
		if global.Config.QiNiu.Enable {
			path, err := qi_niu.UploadImage(byteData, filename, "blog")
			if err != nil {
				global.Logrus.Error(err)
				continue
			}

			resList = append(resList, fileUploadResponse{
				File:      path,
				IsSuccess: true,
				Msg:       fmt.Sprintf("图片上传成功"),
			})
			// 上传成功写入数据库
			global.DB.Create(&models.BannerModel{
				Path: path,
				Hash: imageHash,
				Name: filename,
				Type: ctype.QiNiu,
			})

			continue
		}

		// 上传图片

		if err = c.SaveUploadedFile(file, uploadPath); err != nil {
			resList = append(resList, fileUploadResponse{
				File:      file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片保存失败"),
			})
			continue
		}
		//imageURL := fmt.Sprintf("http://localhost:3030/uploads/%s", filename)
		resList = append(resList, fileUploadResponse{
			File:      file.Filename,
			IsSuccess: true,
			Msg:       fmt.Sprintf("图片上传成功"),
		})

		// 上传成功写入数据库
		global.DB.Create(&models.BannerModel{
			Path: uploadPath,
			Hash: imageHash,
			Name: filename,
			Type: ctype.Local,
		})
	}
	res.OkWithData(c, resList)
}
