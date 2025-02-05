package cows_filter

import (
	"errors"

	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/filters"
)

type ByDeath struct {
}

func (f ByDeath) Apply(fm filters.FilteredModel) error {
	query := fm.GetQuery()
	bodyData, ok := fm.GetFilterParameters()["object"].(CowsFilter)
	if !ok {
		return errors.New("wrong object provided in filter filed object")
	}
	if bodyData.IsDead != nil && *bodyData.IsDead {
		query = query.Where("death_date IS NOT NULL")
	}
	if bodyData.IsDead != nil && !*bodyData.IsDead {
		query = query.Where("death_date IS NULL")
	}
	fm.SetQuery(query)
	return nil
}
