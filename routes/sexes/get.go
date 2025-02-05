package sexes

import (
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes"

	"github.com/gin-gonic/gin"
)

// Get
// @Summary      Get list of sexes
// @Description  Возращает список полов
// @Tags         Sexes
// @Produce      json
// @Success      200  {array}   models.Sex
// @Failure      422  {object}   string
// @Failure      404  {object}   string
// @Router       /sexes [get]
func (s *Sexes) Get() func(*gin.Context) {
	return routes.GenerateGetFunctionByFilters[models.Sex](true)
}
