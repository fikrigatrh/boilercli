package dto

import (
	"boilerplate/internal/model"
	"boilerplate/utils"
	"github.com/guregu/null"
	"math"
)

type BankResponse struct {
	ID          string      `json:"id"`
	Code        string      `json:"code"`
	Name        string      `json:"name"`
	Description null.String `json:"description"` // Nullable field
	IsActive    bool        `json:"isActive"`
	Currency    string      `json:"currency"`
	Type        string      `json:"type"`
	Country     string      `json:"country"`
}

func NewResponseBank(bank model.Bank) BankResponse {
	return BankResponse{
		ID:          bank.ID,
		Code:        bank.Code,
		Name:        bank.Name,
		Description: bank.Description,
		IsActive:    bank.IsActive,
		Currency:    bank.Currency,
		Country:     bank.Country,
		Type:        bank.Type,
	}
}

type BankListResponse []BankResponse

func (l *BankListResponse) NewResponseBankList(list model.BankList) {
	var temp BankResponse
	for _, each := range list {
		temp.ID = each.ID
		temp.Code = each.Code
		temp.Name = each.Name
		temp.Description = each.Description
		temp.Type = each.Type
		*l = append(*l, temp)
	}
}

type ResponseListBank struct {
	Banks    BankListResponse `json:"banks"`
	Metadata utils.Metadata   `json:"metadata"`
}

func NewResponseListBank(list model.BankList, filter model.Filter) (res ResponseListBank) {
	res.Metadata.PageSize = filter.Pagination.PageSize
	res.Metadata.Page = filter.Pagination.Page
	if len(list) > 0 {
		res.Metadata.TotalData = list[0].Count
		res.Metadata.TotalPage = int(math.Ceil(float64(list[0].Count) / float64(filter.Pagination.PageSize)))
	}
	res.Banks = responseBankListFilter(list)
	return
}

func responseBankListFilter(list model.BankList) BankListResponse {
	var (
		bank BankResponse
	)

	res := make(BankListResponse, 0)
	if len(list) > 0 {
		for _, eachLoan := range list {
			bank.ID = eachLoan.ID
			bank.Name = eachLoan.Name
			bank.Description = eachLoan.Description
			bank.IsActive = eachLoan.IsActive
			bank.Code = eachLoan.Code
			bank.Currency = eachLoan.Currency
			bank.Type = eachLoan.Type
			res = append(res, bank)
		}
	}

	return res
}
