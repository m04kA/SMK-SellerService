package domain

// PaginationResult результат с пагинацией
type PaginationResult struct {
	Page  int
	Limit int
	Total int
}
