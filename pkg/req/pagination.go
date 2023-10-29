package req

type PaginationReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

// NewPaginationReq 初始化分页请求, 设置默认值
func NewPaginationReq() *PaginationReq {
	return &PaginationReq{
		PageSize: 5, // 设置默认值为5
	}
}

func GetOffset(p *PaginationReq) int {
	offset := (p.Page - 1) * p.PageSize
	if offset < 0 {
		return 0
	}
	return offset
}
