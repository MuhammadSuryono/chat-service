package dto

type PaginationDto struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
