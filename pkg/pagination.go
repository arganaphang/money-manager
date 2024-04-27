package pkg

type Pagination struct {
	Page    uint `json:"page" form:"page"`
	PerPage uint `json:"per_page" form:"per_page"`
	Total   uint `json:"total" form:"total"`
}

type LimitOffset struct {
	Limit  uint
	Offset uint
}

func (p *Pagination) TransformToLimitOffset() LimitOffset {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	limit := p.PerPage
	offset := (p.Page - 1) * p.PerPage
	return LimitOffset{Limit: limit, Offset: offset}
}
