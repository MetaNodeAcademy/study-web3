package db

import "gorm.io/gorm"

// QueryParams 分页查询参数
type QueryParams struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"pageSize" binding:"min=1"`
}

// PagedResult 分页结果
type PagedResult struct {
	// 总记录数
	Total int64 `json:"total"`
	// 当前页码
	Page int `json:"page"`
	// 每页记录数
	PageSize int `json:"pageSize"`
	// 总页数
	TotalPages int `json:"totaPages"`
	// 是否有下一页
	HasNextPage bool `json:"hasNextPage"`
	// 查询结果数据
	Data interface{} `json:"data"`
}

// Paginate 通用分页查询函数
func Paginate(db *gorm.DB, params QueryParams, dest interface{}) (*PagedResult, error) {
	var total int64
	offset := (params.Page - 1) * params.PageSize

	// 查询总数
	if err := db.Model(dest).Count(&total).Error; err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := int((total + int64(params.PageSize) - 1) / int64(params.PageSize))

	// 查询当前页数据
	query := db.Limit(params.PageSize).Offset(offset)
	if err := query.Find(dest).Error; err != nil {
		return nil, err
	}

	return &PagedResult{
		Total:       total,
		Page:        params.Page,
		PageSize:    params.PageSize,
		TotalPages:  totalPages,
		HasNextPage: params.Page < totalPages,
		Data:        dest,
	}, nil
}
