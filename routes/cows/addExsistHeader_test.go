package cows

import (
	"reflect"
	"testing"
	"time"

	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/filters/cows_filter"
	// "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"
)

func Test_addExistsHeaderToFile(t *testing.T) {
	var ep uint = 25
	Num := ""
	rshn := "RSHN"
	isTrue := true
	now := time.Now()
	ordByDesc := false
	// date := models.DateOnly{Time: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)}
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Format("1800-01-21")
	test := []cows_filter.CowsFilter {
		{
			BirkingDateFrom: nil,
			BirkingDateTo: nil,
			BirthDateFrom: nil,
			BirthDateTo: nil,
			BreedId: nil,
			CalvingDateFrom: nil,
			CalvingDateTo: nil,
			ControlMilkingDateFrom: nil,
			ControlMilkingDateTo: nil,
			CreatedAtFrom: nil,
			CreatedAtTo: nil,
			DepartDateFrom: nil,
			DepartDateTo: nil,
			EntitiesOnPage: &ep,
			ExteriorFrom: nil,
			ExteriorTo: nil,
			HasAnyIllnes: nil,
			HozId: nil,
			IllDateFrom: nil,
			IllDateTo: nil,
			InbrindingCoeffByFamilyFrom: nil,
			InbrindingCoeffByFamilyTo: nil,
			InbrindingCoeffByGenotypeFrom: nil,
			InbrindingCoeffByGenotypeTo: nil,
			InseminationDateFrom: nil,
			InseminationDateTo: nil,
			IsAborted: nil,
			IsDead: nil,
			IsGenotyped: nil,
			IsIll: nil,
			IsStillBorn: nil,
			IsTwins: nil,
			MonogeneticIllneses: []uint{},
			OrderBy: &rshn,
			OrderByDesc: &ordByDesc,
			SearchQuery: &Num,
		},
		{
			BirkingDateFrom: nil,
			BirkingDateTo: nil,
			BirthDateFrom: nil,
			BirthDateTo: nil,
			BreedId: []uint{1},
			CalvingDateFrom: nil,
			CalvingDateTo: nil,
			ControlMilkingDateFrom: nil,
			ControlMilkingDateTo: nil,
			CreatedAtFrom: nil,
			CreatedAtTo: nil,
			DepartDateFrom: &date,
			DepartDateTo: &date,
			EntitiesOnPage: &ep,
			ExteriorFrom: nil,
			ExteriorTo: nil,
			HasAnyIllnes: nil,
			HozId: nil,
			IllDateFrom: nil,
			IllDateTo: nil,
			InbrindingCoeffByFamilyFrom: nil,
			InbrindingCoeffByFamilyTo: nil,
			InbrindingCoeffByGenotypeFrom: nil,
			InbrindingCoeffByGenotypeTo: nil,
			InseminationDateFrom: nil,
			InseminationDateTo: nil,
			IsAborted: nil,
			IsDead: &isTrue,
			IsGenotyped: nil,
			IsIll: nil,
			IsStillBorn: nil,
			IsTwins: nil,
			MonogeneticIllneses: []uint{},
			OrderBy: &rshn,
			OrderByDesc: &ordByDesc,
			SearchQuery: &Num,
		},
	}
	

	tempT := reflect.TypeOf(FilterSerializedCow{})
	v := reflect.ValueOf(FilterSerializedCow{})
	for _,val := range test {
		ctrOfGotArr := 0
		tempHeader := []string{}
		got := addExistsHeaderToFile(&val)
		t.Log(got)
		for i:=1; i < tempT.NumField(); i++ {
			if v.Field(i).Kind() == reflect.Slice {
				continue
			}
			if got[ctrOfGotArr] {
				tempHeader = append(tempHeader, tempT.Field(i).Name)
				ctrOfGotArr++
			}

		}
		t.Log(tempHeader)
	}
}
