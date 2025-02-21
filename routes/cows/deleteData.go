package cows

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/filters/cows_filter"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"
	"github.com/gin-gonic/gin"
)

func (c *Cows) toDeleteRows() func(*gin.Context) {
	return func(c *gin.Context) {
		roleId, exists := c.Get("RoleId") // номер роли при анутификации
		if !exists {
			c.JSON(http.StatusInternalServerError, "RoleId не найден в контексте")
			return
		}
		
		jsonData, err := io.ReadAll(c.Request.Body) // Читаем тело запроса
		if err != nil {
			c.JSON(500, err.Error())
			return
		}

		bodyData := cows_filter.CowsFilter{}
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData, &bodyData)
			if err != nil {
				c.JSON(422, err.Error())
				return
			}
		}
		
		if roleId != 4 { 
			c.JSON(401, err.Error())
			return
		}
		
		db := models.GetDb()
		query := db.Model(&models.Cow{}).Select("id").Preload("FarmGroup").Preload("Genetic").Where("approved <> -1") 
		if nQuery, err := AddFiltersToQuery(bodyData, query); err != nil {
			c.JSON(422, err.Error())
			return
		} else {
			query = nQuery
		}

		res := db.Where("id IN (?)", query).Delete(&models.Cow{})

		resCount := res.RowsAffected
		if err := res.Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"N":   resCount,
			// "LST": rc,
			// "query": query,
		})

	}
}