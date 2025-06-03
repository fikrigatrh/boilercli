package usecase

import (
	"boilerplate/internal/dto"
	errorUc "boilerplate/internal/error"
	"boilerplate/utils"
	"context"
	"fmt"
	"github.com/saucon/sauron/v2/pkg/log"
)

func (u *usecase) ResolveBank(ctx context.Context, request dto.BankListRequest) (dto.ResponseListBank, error) {
	filter := request.ToFilter()

	bankList, err := u.exampleRepo.ResolveByFilter(ctx, filter)
	if err != nil {
		u.log.Error(log.LogData{
			Err:         err,
			Description: fmt.Sprintf("error while resolving bank list, filter: %v", filter),
		})
		err = utils.MakeError(errorUc.InternalServerError)
		return dto.ResponseListBank{}, err
	}

	res := dto.NewResponseListBank(bankList, filter)
	return res, nil
}
