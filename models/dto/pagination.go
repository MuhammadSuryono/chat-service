package dto

type PaginationDto struct {
	Page         int    `form:"page"`
	Limit        int    `form:"limit"`
	LastDataDate string `form:"last_data_date"`
}
