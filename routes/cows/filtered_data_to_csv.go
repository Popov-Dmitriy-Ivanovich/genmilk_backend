package cows

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Переменные функции
// Путь к файлу
const PathToCSVFile = "./static/csv/"
const formatToDate = "2006-01-02"

// var (
// 	cellName string
// 	ListName = "List1"
// )

func ToCSVFile(fsc []FilterSerializedCow) (string, error) {
	// Создаем csv файл с отложеным закрытием
	now := time.Now()
	fullPath := PathToCSVFile + "filtered_data_" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(uint64(len(fsc)), 16) + ".csv"
	
	cswFile, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("Ошибка создания файла: %v", err)
	}
	defer cswFile.Close()

	writer := csv.NewWriter(cswFile)
	writer.Comma = ';'
	defer writer.Flush()
	
	// // Записываем заголовки
	headers := getHeaders()
	err = writer.Write(headers)
	if err != nil {
		return "", fmt.Errorf("Ошибка записи заголовков: %v", err)
	}

	
	// Записываем данные
	for _, data := range fsc {
		var strOfFile = []string{}
		// Объявим функция для уменьшения размера кода
		// Функция записи ошибочной строки
		writeErrorRequiredData := func() error {
			err = writer.Write([]string{"Отсутсвуют обязательные данные"})
			if err != nil {
			    return err
			}
			return nil
		}
		
		
		if data.InventoryNumber == nil || *data.InventoryNumber == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		if data.FarmGroupName == "" {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}
		if data.BirthDate.IsZero() {
			err = writeErrorRequiredData()
			if err != nil {
				return "", err
			}
			continue
		}

		// Поля Genotyped и Approved будут существовать в любом случае
		// ===== //
		// Записываем данные
		
		if data.RSHNNumber != nil{ // РСХН всегда хранит номер
		    strOfFile = append(strOfFile,*data.RSHNNumber)
		}else {
			strOfFile = append(strOfFile,"")
		}
		if data.InventoryNumber != nil {
		    strOfFile = append(strOfFile,*data.InventoryNumber)
		}else {
			strOfFile = append(strOfFile,"")
		}
		strOfFile = append(strOfFile,data.Name)
		strOfFile = append(strOfFile,data.FarmGroupName) 
		strOfFile = append(strOfFile,data.BirthDate.Time.Format(formatToDate))
		
		
		strOfFile = append(strOfFile,strconv.FormatBool(data.Genotyped))
		strOfFile = append(strOfFile,strconv.FormatBool(data.Approved))
		

		if data.DepartDate != nil {
			strOfFile = append(strOfFile,data.DepartDate.Time.Format(formatToDate))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.BreedName != nil { // Проверка на пустой указатель
			strOfFile = append(strOfFile,*data.BreedName)
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.BirkingDate != nil {
			strOfFile = append(strOfFile,data.BirkingDate.Time.Format(formatToDate))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.GenotypingDate != nil {
			strOfFile = append(strOfFile,data.GenotypingDate.Time.Format(formatToDate))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.InbrindingCoeffByFamily != nil {
			strOfFile = append(strOfFile,fmt.Sprintf("%f",*data.InbrindingCoeffByFamily))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.InbrindingCoeffByGenotype != nil {
			strOfFile = append(strOfFile,fmt.Sprintf("%f",*data.InbrindingCoeffByGenotype))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.ExteriorRating != nil {
			strOfFile = append(strOfFile,fmt.Sprintf("%f",*data.ExteriorRating))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.SexName != nil {
			strOfFile = append(strOfFile,*data.SexName)
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.HozName != nil {
			strOfFile = append(strOfFile,*data.HozName)
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.DeathDate != nil {
			strOfFile = append(strOfFile,data.DeathDate.Time.Format(formatToDate))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.IsDead != nil {
			strOfFile = append(strOfFile,strconv.FormatBool(*data.IsDead))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.IsTwins != nil {
			strOfFile = append(strOfFile,strconv.FormatBool(*data.IsTwins))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.IsStillBorn != nil {
			strOfFile = append(strOfFile,strconv.FormatBool(*data.IsStillBorn))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.IsAborted != nil {
			strOfFile = append(strOfFile,strconv.FormatBool(*data.IsAborted))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.IsGenotyped != nil {
			strOfFile = append(strOfFile,strconv.FormatBool(*data.IsGenotyped))
		} else {
			strOfFile = append(strOfFile,"")
		}
		if data.CreatedAt != nil {
			strOfFile = append(strOfFile,data.CreatedAt.Time.Format(formatToDate))
		} else {
			strOfFile = append(strOfFile,"")
		}
		

		// Запись строки в csv файл
		if err = writer.Write(strOfFile); err != nil {
			return "", err
		}
	}
	return fullPath, nil
}
