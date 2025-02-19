package cows

import (
	// "path/filepath"

	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/filters/cows_filter"

	//"time"
)

func addExistsHeaderToFile(filter *cows_filter.CowsFilter) []bool {
	var arrExistOfHeaders []bool = []bool{}

	// Обязательные поля
	// arrExistOfHeaders = append(arrExistOfHeaders, true) // ID
	arrExistOfHeaders = append(arrExistOfHeaders, true) // Поле РСХН номер
	arrExistOfHeaders = append(arrExistOfHeaders, true) // Поле InventoryNumber
	arrExistOfHeaders = append(arrExistOfHeaders, true) // Name
	arrExistOfHeaders = append(arrExistOfHeaders, true) // FarmGroupName
	arrExistOfHeaders = append(arrExistOfHeaders, true) // BirthDate
	arrExistOfHeaders = append(arrExistOfHeaders, true) // Genotyped
	arrExistOfHeaders = append(arrExistOfHeaders, true) // Approved

	// Наличие заголовка DepartDate
	if filter.DepartDateTo != nil && *filter.DepartDateTo != "" ||
		filter.DepartDateFrom != nil && *filter.DepartDateFrom != "" {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}

	// Наличие заголовка BreedName
	if len(filter.BreedId) != 0 {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
		}else {
			arrExistOfHeaders = append(arrExistOfHeaders, false)
		}
		
	// Наличие заголовка BirkingDate
	if filter.BirkingDateFrom != nil && *filter.BirkingDateFrom != "" ||
		filter.BirkingDateTo != nil && *filter.BirkingDateTo != "" {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Наличие заголовка GenotypingDate  
	if filter.GenotypingDateFrom != nil && *filter.GenotypingDateFrom != "" ||
	filter.GenotypingDateTo != nil && *filter.GenotypingDateTo != "" {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
		}else {
			arrExistOfHeaders = append(arrExistOfHeaders, false)
		}
		// Наличие заголовка InbrindingCoeffByFamily 
		if filter.InbrindingCoeffByFamilyFrom != nil || filter.InbrindingCoeffByFamilyTo != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Наличие заголовка InbrindingCoeffByGenotype
	if filter.InbrindingCoeffByGenotypeFrom != nil || filter.InbrindingCoeffByGenotypeTo != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}

	// Заголовк оценки экстерьера
	if filter.ExteriorFrom != nil || filter.ExteriorTo != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}

	// Заголовок SexName
	if len(filter.Sex) != 0 {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Заголовок hozName
	if filter.HozId != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, false) // Вычеркнули поле
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}

	// Заголовк поля даты и факта смерти
	if filter.IsDead != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Заголовок IsTwins
	if filter.IsTwins != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Заголовок факта мертворождения
	if filter.IsStillBorn != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Заголовок факта аборта
	if filter.IsAborted != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}
	// Заголовок факта генотипирования
	if filter.IsGenotyped != nil {
		arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}

	// Момент записи
	if filter.CreatedAtFrom != nil && *filter.CreatedAtFrom != "" ||
		filter.CreatedAtTo != nil && *filter.CreatedAtTo != "" {
			arrExistOfHeaders = append(arrExistOfHeaders, true)
	}else {
		arrExistOfHeaders = append(arrExistOfHeaders, false)
	}

	return arrExistOfHeaders
}
