package req

type PaginationReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func GetOffset(p *PaginationReq) int {
	return (p.Page - 1) * p.PageSize
}
