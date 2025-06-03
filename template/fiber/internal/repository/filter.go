package repository

import (
	"boilerplate/internal/model"
	"errors"
	"fmt"
	"strings"
)

func ComposeFilter(filter model.Filter) (clauseStr string, args []interface{}, err error) {
	args = make([]interface{}, 0)
	clause := make([]string, 0)

	if len(filter.FilterFields) > 0 {
		var (
			whereQueries []string
			whereArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				whereQueries = append(whereQueries, fmt.Sprintf("%s = ?", filterField.Field))
				whereArgs = append(whereArgs, filterField.Value)
			case model.OperatorOr:
				whereQueries = append(whereQueries, fmt.Sprintf("(%s = ? OR %s = ?)", strings.Split(filterField.Field, ",")[0], strings.Split(filterField.Field, ",")[1]))
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = errors.New(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorRange:
				whereQueries = append(whereQueries, fmt.Sprintf("%s BETWEEN ? AND ?", filterField.Field))
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = errors.New(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIn:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = errors.New(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, "?")
				}
				whereQueries = append(whereQueries, fmt.Sprintf("%s IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIsNull:
				value, ok := filterField.Value.(bool)
				if !ok {
					err = errors.New(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				if value {
					whereQueries = append(whereQueries, fmt.Sprintf("%s IS NULL", filterField.Field))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("%s IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				whereQueries = append(whereQueries, fmt.Sprintf("%s != ?", filterField.Field))
				whereArgs = append(whereArgs, filterField.Value)
			}
		}

		clause = append(clause, fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND ")))
		args = append(args, whereArgs...)
	}

	if len(filter.FilterFields) > 0 {
		var sortQuery []string
		query := ""
		cond := ""
		sortField := ""
		for _, sort := range filter.Sorts {
			sortField = sort.Field
			if sort.Field != "" {
				if sort.Condition != "" {
					sortQuery = append(sortQuery, fmt.Sprintf("%s", sort.Field))
					cond = sort.Condition
				} else {
					sortQuery = append(sortQuery, fmt.Sprintf("%s %s", sort.Field, sort.Order))
				}
			}
		}
		if sortField != "" {
			if cond == model.ConditionIfNull {
				query = fmt.Sprintf(" ORDER BY ifnull(%s) %s", strings.Join(sortQuery, ","), filter.Sorts[0].Order)
			} else {
				query = fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
			}
		}

		clause = append(clause, query)
	} else {
		var sortQuery []string
		sortField := ""
		cond := ""
		query := ""
		for _, sort := range filter.Sorts {
			sortField = sort.Field
			if sort.Field != "" {
				if sort.Condition != "" {
					sortQuery = append(sortQuery, fmt.Sprintf("%s", sort.Field))
					cond = sort.Condition
				} else {
					sortQuery = append(sortQuery, fmt.Sprintf("%s %s", sort.Field, sort.Order))
				}
			}
		}
		if sortField != "" {
			if cond == model.ConditionIfNull {
				query = fmt.Sprintf(" ORDER BY ifnull(%s) %s", strings.Join(sortQuery, ","), filter.Sorts[0].Order)
			} else {
				query = fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
			}
		}

		clause = append(clause, query)
	}

	if filter.Pagination.PageSize > 0 {
		query := ""
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
		clause = append(clause, query)
	}

	clauseStr = strings.Join(clause, "")

	return
}
