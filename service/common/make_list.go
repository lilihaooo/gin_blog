package common

import (
	"blog_gin/pkg/req"
	"gorm.io/gorm"
)

func MakeList[T any](model T, db *gorm.DB, pageReq req.PaginationReq) (list []T, count int64, err error) {
	offset := req.GetOffset(&pageReq)
	if offset < 0 {
		offset = 0
	}

	err = db.Model(model).Select("id").Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(pageReq.PageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

//global.DB.Where(condition).Order(sort)
