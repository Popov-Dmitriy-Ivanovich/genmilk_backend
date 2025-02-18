package cows

import (
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"

	"github.com/gin-gonic/gin"
)

func getPercents(x float64, min float64, max float64) float64 {
	return ((x-min)/(max-min) - 0.5) * 100
}

func GetGradePercents(grades models.Grade, statistics models.Statistics) models.Grade {
	//((x-min)/(max-min) - 0.5) * 100
	res := models.Grade{}
	if grades.GeneralValue != nil {
		res.GeneralValue = new(float64)
		*res.GeneralValue = getPercents(*grades.GeneralValue, statistics.MinGeneralValue, statistics.MaxGeneralValue)
	}
	if grades.EbvMilk != nil {
		res.EbvMilk = new(float64)
		*res.EbvMilk = getPercents(*grades.EbvMilk, statistics.MinEbvMilk, statistics.MaxEbvMilk)
	}
	if grades.EbvProtein != nil {
		res.EbvProtein = new(float64)
		*res.EbvProtein = getPercents(*grades.EbvProtein, statistics.MinEbvProtein, statistics.MaxEbvProtein)
	}
	if grades.EbvFat != nil {
		res.EbvFat = new(float64)
		*res.EbvFat = getPercents(*grades.EbvFat, statistics.MinEbvFat, statistics.MaxEbvFat)
	}
	if grades.EbvInsemenation != nil {
		res.EbvInsemenation = new(float64)
		*res.EbvInsemenation = getPercents(*grades.EbvInsemenation, statistics.MinEbvInsemenation, statistics.MaxEbvInsemenation)
	}
	if grades.EvbService != nil {
		res.EvbService = new(float64)
		*res.EvbService = getPercents(*grades.EvbService, statistics.MinEvbService, statistics.MaxEvbService)
	}
	return res
}

// Grades
// @Summary      Get grades
// @Description  Возращает словарь с двумя ключам "ByRegion" - оценки по региону и "ByHoz" - оценки по хозяйству
// @Tags         Cows
// @Param        id   path      int  true  "ID коровы для которой ищутся оценки"
// @Produce      json
// @Success      200  {object}   map[string]models.Grade
// @Failure      500  {object}   string
// @Router       /cows/{id}/grades [get]
func (c *Cows) Grades() func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := models.GetDb()
		cow := models.Cow{}
		if err := db.
			Preload("GradeRegion").
			Preload("GradeHoz").
			Preload("FarmGroup").
			Preload("FarmGroup.District").
			First(&cow, id).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}
		regStat := models.RegionalStatistics{}
		hozStat := models.HozStatistics{}

		if err := db.First(&regStat, map[string]interface{}{"region_id": cow.FarmGroup.District.RegionId}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		if err := db.First(&hozStat, map[string]any{"hoz_id": cow.FarmGroupId}).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		regionPercents := models.Grade{}
		hozPercents := models.Grade{}

		if cow.GradeRegion != nil {
			regionPercents = GetGradePercents(cow.GradeRegion.Grade, regStat.Statistics)
		}

		if cow.GradeHoz != nil {
			hozPercents = GetGradePercents(cow.GradeHoz.Grade, hozStat.Statistics)
		}

		c.JSON(200, gin.H{
			"ByRegion":         cow.GradeRegion,
			"ByHoz":            cow.GradeHoz,
			"ByRegionPercents": regionPercents,
			"ByHozPercents":    hozPercents,
		})
	}
}
