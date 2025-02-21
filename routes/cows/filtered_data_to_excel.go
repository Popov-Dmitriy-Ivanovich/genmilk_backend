package cows

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	// "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/filters/cows_filter"
	// "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"
	"github.com/xuri/excelize/v2"
)

// Переменные функции
// Путь к файлу
const PathToExcelFile = "./static/excel/"
// var forRelyOnFiltersParams

var (
	cellName string
	ListName = "Sheet1"
	EmptyValue = "Нет данных"
)

func ToExcelOld(fsc []FilterSerializedCow, SelecsId []uint64, hw []bool) (string, error) {
	f := excelize.NewFile()

	// Создаем новый лист
	index, err := f.NewSheet(ListName)
	if err != nil {
		return "", err
	}
	
	// Устанавливаем активный лист
	f.SetActiveSheet(index)

	// Записываем заголовки
	// headers := getHeaders()
	for i, columnName := range getHeaders(fsc[0], hw) {
		// if i == 0 {
		// 	continue
		// }
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		err = f.SetCellValue(ListName, cell, columnName)
		if err != nil {
			return "", err
		}
	}

	// dateStyle, err := f.NewStyle(&excelize.Style{NumFmt: 165})
	// if err != nil {
	// 	return "", err
	// }

	// Записываем данные
	for row, data := range fsc {
		colNum := 1
		ctr := 7 // индекс 
		// Функция инкрементирования ячейки
		Incr := func() {
			cellName, err = excelize.CoordinatesToCellName(colNum, row+2)
			colNum++
		}
		// Запись Id Селекса
		Incr()
		if SelecsId[row] != 0 {
			if err = f.SetCellInt(ListName, cellName, int(SelecsId[row])); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
				return "", err
			} else {
				Incr()
			}
		}

		// Поля Genotyped и Approved будут существовать в любом случае
		// ===== //
		// Записываем данные
		if data.RSHNNumber != nil{ // РСХН всегда хранит номер
		    if err = f.SetCellValue(ListName, cellName, *data.RSHNNumber); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.InventoryNumber != nil {
		    if err = f.SetCellValue(ListName, cellName, *data.InventoryNumber); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.Name != "" {
			if err = f.SetCellValue(ListName, cellName, data.Name); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.FarmGroupName != "" {
			if err = f.SetCellValue(ListName, cellName, data.FarmGroupName); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
				return "", err
			} else {
				Incr()
			}
		}	
		if !data.BirthDate.Time.IsZero() {
			if err = f.SetCellValue(ListName, cellName, data.BirthDate.Time); err != nil {
				// f.SetCellStyle(ListName, cellName, cellName, dateStyle)
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		if data.Genotyped {
			if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		
		if data.Approved {
			if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
				return "", err
			} else {
				Incr()
			}
		}else {
			if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
				return "", err
			} else {
				Incr()
			}
		}
		// Проверку придется начать с 8-ого эл-та т.е. с 7-ого индекса
		if hw[ctr] {
			if data.DepartDate != nil && !data.DepartDate.Time.IsZero(){
				if err = f.SetCellValue(ListName, cellName, data.DepartDate.Time); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.BreedName != nil { // Проверка на пустой указатель
				if err = f.SetCellValue(ListName, cellName, *data.BreedName); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.BirkingDate != nil && !data.BirkingDate.Time.IsZero(){
				if err = f.SetCellValue(ListName, cellName, data.BirkingDate.Time); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.GenotypingDate != nil && !data.GenotypingDate.Time.IsZero(){
				if err = f.SetCellValue(ListName, cellName, data.GenotypingDate.Time); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.InbrindingCoeffByFamily != nil {
				if err = f.SetCellValue(ListName, cellName, *data.InbrindingCoeffByFamily); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.InbrindingCoeffByGenotype != nil {
				if err = f.SetCellValue(ListName, cellName, *data.InbrindingCoeffByGenotype); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.ExteriorRating != nil {
				if err = f.SetCellValue(ListName, cellName, *data.ExteriorRating); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.SexName != nil {
				if err = f.SetCellValue(ListName, cellName, *data.SexName); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.HozName != nil {
				if err = f.SetCellValue(ListName, cellName, *data.HozName); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.DeathDate != nil && !data.DeathDate.Time.IsZero() {
				if err = f.SetCellValue(ListName, cellName, data.DeathDate.Time); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.IsDead != nil {
				if *data.IsDead {
					if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}else {
					if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.IsTwins != nil {
				if *data.IsTwins {
					if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}else {
					if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.IsStillBorn != nil {
				if *data.IsStillBorn {
					if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}else {
					if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.IsAborted != nil {
				if *data.IsAborted {
					if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}else {
					if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.IsGenotyped != nil {
				if *data.IsGenotyped {
					if err = f.SetCellValue(ListName, cellName, "да"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}else {
					if err = f.SetCellValue(ListName, cellName, "нет"); err != nil {
						return "", err
					} else {
						Incr()
					}
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
		ctr++
		if hw[ctr] {
			if data.CreatedAt != nil && !data.CreatedAt.Time.IsZero(){
				if err = f.SetCellValue(ListName, cellName, data.CreatedAt.Time); err != nil {
					return "", err
				} else {
					Incr()
				}
			} else {
				if err = f.SetCellValue(ListName, cellName, EmptyValue); err != nil {
					return "", err
				} else {
					Incr()
				}
			}
		}
	}

	// Сохраняем файл в cow_backend\frontend\static
	now := time.Now()
	fullPath := PathToExcelFile + "filtered_data_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(uint64(len(fsc)), 16) + ".xlsx"
	if err := f.SaveAs(fullPath); err != nil {
		return "", fmt.Errorf("Ошибка создания Excel файла Error: %v", err)
	} else {
		return fullPath, nil
	}

}



func getHeaders(fsc FilterSerializedCow, b []bool) []string { // Получаем заголовки таблицы
	var headers []string
	t := reflect.TypeOf(fsc)
	v := reflect.ValueOf(fsc)

	// Нужен Id Селекса
	headers = append(headers, "ID селекса")
	ctr := 0
	for i := 1; i < t.NumField(); i++ { // Не берем поле id
		if v.Field(i).Kind() == reflect.Slice {
			continue
		}
		if !b[ctr] {
			ctr++
			continue
		}
		ctr++
		
		field := t.Field(i)
		headers = append(headers, field.Name)
	}

	return headers
}
