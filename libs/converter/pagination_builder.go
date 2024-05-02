package converter

import (
	"fmt"
	"math"
)

type PaginationBuilder struct {
	Page     int         `json:"page"`
	Limit    int         `json:"-"`
	Offset   int         `json:"-"`
	Total    string      `json:"total"`
	PerPage  int         `json:"perPage"`
	LastPage int         `json:"lastPage"`
	Data     interface{} `json:"data"`
}

func NewPaginationQuery(page int, limit int) PaginationBuilder {
	var offset = 0
	if page != 0 {
		offset = (page - 1) * limit
	}

	return PaginationBuilder{
		Page:    page,
		Limit:   limit,
		PerPage: limit,
		Offset:  offset,
	}
}

func (p PaginationBuilder) ToPaginationDataResult(countData int64, data interface{}) *PaginationBuilder {
	var totalPage = float64(countData) / float64(p.Limit)
	var round = math.Round(totalPage)
	if round >= totalPage {
		p.LastPage = int(round)
	} else {
		p.LastPage = int(round) + 1
	}

	p.Total = fmt.Sprintf("%v", countData)
	p.Data = data
	return &p
}
