package example

import (
	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"boilerplate/internal/repository/query"
	"context"
	"fmt"
	"github.com/saucon/sauron/v2/pkg/log"
	"time"
)

func (b *BankRepo) ResolveByFilter(ctx context.Context, filter model.Filter) (model.BankList, error) {
	var (
		whereClause string
		result      model.BankList
		err         error
	)

	clauses, args, err := repository.ComposeFilter(filter)
	if err != nil {
		b.log.Error(log.LogData{
			Description: "[ResolveByFilter.bankrepo] composeFilter",
			StartTime:   time.Now(),
			Err:         err,
		})
		return result, err
	}

	whereClause += clauses

	baseQuery := fmt.Sprintf("%s %s",
		query.Bank.Select,
		whereClause,
	)

	err = b.Infra.DbPsql.DB.Raw(baseQuery, args...).Scan(&result).WithContext(ctx).Error
	if err != nil {
		b.log.Error(log.LogData{
			Description: "[ResolveByFilter.bankrepo] read raw query",
			StartTime:   time.Now(),
			Err:         err,
		})
		return result, err
	}

	return result, nil
}
