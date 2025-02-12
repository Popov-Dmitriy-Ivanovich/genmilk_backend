package cows

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/filters/cows_filter"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"
	"github.com/gin-gonic/gin"
)
// Filter
// @Summary      Get excel filtered list of cows
// @Description  Возращает словарь с двумя ключами "N", "LST". По ключу "N" - общее кол-во найденных коров и их характеристики,
// @Description  по ключу "LST" - путь к файлу excel в директории ./static/excel/
// @Tags         Cows
// @Param        filter    body     cows_filter.CowsFilter  true  "applied filters"
// @Accept       json
// @Produce      json
// @Success      200  {array}   map[string]FilterSerializedCow
// @Failure      422  {object}  string
// @Failure      500  {object}  string
// @Router       /cows/filterExcel [post]
func (c *Cows) SendExcel() func(*gin.Context) {
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

		bodyData := cows_filter.CowsFilter{} // получаем структуру фильтров с сайта
		if len(jsonData) != 0 {
			err = json.Unmarshal(jsonData, &bodyData)
			if err != nil {
				c.JSON(422, err.Error())
				return
			}
		}

		if roleId == 1 {
			farmIdStr, exists := c.Get("FarmId") // получаем фермерское хозяйство
			if !exists {
				c.JSON(http.StatusInternalServerError, "FarmId не найден в контексте")
				return
			}

			farmIdUint64, err := strconv.ParseUint(farmIdStr.(string), 10, 0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Ошибка преобразования FarmId: "+err.Error())
				return
			}

			farmId := uint(farmIdUint64)
			bodyData.HozId = &farmId
		}

		if roleId == 2 {
			regionIdStr, exists := c.Get("RegionId")
			if !exists {
				c.JSON(http.StatusInternalServerError, "RegionId не найден в контексте")
				return
			}

			regionIdUint64, err := strconv.ParseUint(regionIdStr.(string), 10, 0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Ошибка преобразования RegionId: "+err.Error())
				return
			}

			regionId := uint(regionIdUint64)
			bodyData.RegionId = &regionId
		}

		db := models.GetDb()
		query := db.Model(&models.Cow{}).Preload("FarmGroup").Preload("Genetic").Where("approved <> -1")
		if nQuery, err := AddFiltersToQuery(bodyData, query); err != nil {
			c.JSON(422, err.Error())
			return
		} else {
			query = nQuery
		}
		// ====================================================================================================
		// ======================================= PAGINATION =================================================
		// ====================================================================================================
		recordsPerPage := uint64(50)
		pageNumber := uint64(1)
		if bodyData.EntitiesOnPage != nil {
			recordsPerPage = uint64(*bodyData.EntitiesOnPage)
		}

		if bodyData.PageNumber != nil {
			pageNumber = uint64(*bodyData.PageNumber)
		}
		// ====================================================================================================
		// ================================== Get final query result ==========================================
		// ====================================================================================================
		resCount := int64(0)
		if err := query.Count(&resCount).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		query = query.Limit(int(recordsPerPage)).Offset(int(recordsPerPage) * int(pageNumber-1)).Order("inventory_number")
		dbCows := []models.Cow{}
		if err := query.Debug().Find(&dbCows).Error; err != nil {
			c.JSON(500, err.Error())
			return
		}

		fsc := make([]FilterSerializedCow, 0, len(dbCows))
		for _, c := range dbCows {
			fsc = append(fsc, serializeByFilter(&c, &bodyData))
		}

		filePath, err := ToExcelOld(fsc)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		file := filePath[len(PathToExcelFile):] // Возврат пути файл в директории ./static/excel/
		resCount = int64(len(fsc)) // Число найденных коров и их характеристики
		// fmt.Print(query)
		c.JSON(200, gin.H{
			"N":   resCount,
			"LST": file,
			// "query": query,
		})
	}
}