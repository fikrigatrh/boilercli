package dto

import (
	"boilerplate/internal/model"
	"github.com/guregu/null"
	"strconv"
	"strings"
)

type BankListRequest struct {
	Code     null.String `json:"code"`
	Name     null.String `json:"name"`
	Page     null.String `json:"page" example:"1"`
	Size     null.String `json:"size" example:"50"`
	Currency null.String `json:"currency" example:"HKD"`
	Type     null.String `json:"type" example:"bank"`
}

func (b *BankListRequest) ToFilter() model.Filter {
	var (
		page int = 1
		size int = 10
	)

	if b.Page.String != "" {
		page, _ = strconv.Atoi(b.Page.ValueOrZero())
	}
	if b.Size.String != "" {
		size, _ = strconv.Atoi(b.Size.ValueOrZero())
	}

	filter := model.Filter{
		Pagination: model.Pagination{
			Page:     page,
			PageSize: size,
		},
		Sorts: []model.Sort{
			{
				Field: "code",
				Order: model.SortAsc,
			},
		},
	}

	if b.Name.String != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "name",
			Operator: model.OperatorLike,
			Value:    b.Name.ValueOrZero(),
		})
	}
	if b.Code.String != "" {
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "code",
			Operator: model.OperatorEqual,
			Value:    b.Code.ValueOrZero(),
		})
	}
	if b.Currency.Valid {
		b.Currency = null.StringFrom(strings.ToUpper(b.Currency.String))
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "currency",
			Operator: model.OperatorEqual,
			Value:    b.Currency.ValueOrZero(),
		})
	}
	if b.Type.Valid {
		b.Currency = null.StringFrom(strings.ToUpper(b.Currency.String))
		filter.FilterFields = append(filter.FilterFields, model.FilterField{
			Field:    "type",
			Operator: model.OperatorEqual,
			Value:    b.Type.ValueOrZero(),
		})
	}

	return filter
}

type BankRequest struct {
	Code null.String `json:"code"`
	Name null.String `json:"name"`
}
