package main

import (
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/models"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/admin"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/analitics"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/auth"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/breeds"
	checkmilks "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/check_milks"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/cows"
	dailymilks "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/daily_milks"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/districts"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/farms"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/gui"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/lactations"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/load"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/monogenetic_illnesses"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/partners"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/regions"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/sexes"
	"github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/updates"
	user_create "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/routes/user"
	"fmt"
	"text/template"

	// "net/http"
	_ "github.com/Popov-Dmitriy-Ivanovich/genmilk_backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

func seq(start, end int) []int {
	var nums []int
	for i := start; i <= end; i++ {
		nums = append(nums, i)
	}
	return nums
}

// @title           GenMilk API
// @version         1.0
// @description     Сваггер сгенерирован из кода, поэтому может содержать неточности. По мере возможности они будут описаны далее
// @description     Все даты передаются как строки
// @description     Большая часть рутов не возвращает вложенные объекты
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      genmilk.ru
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	models.GetDb()

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"seq": seq,
	})

	r.LoadHTMLGlob("templates/*")

	apiGroup := r.Group("/api")
	routes.WriteRoutes(apiGroup, &routes.Api{}, &regions.Regions{}, &farms.Farms{}, &breeds.Breeds{}, &checkmilks.CheckMilks{},
		&cows.Cows{}, &dailymilks.DailyMilk{}, &districts.Districts{}, &lactations.Lactations{}, &sexes.Sexes{}, &analitics.Analitics{},
		&monogenetic_illnesses.MonogeneticIllneses{}, &gui.Gui{}, &load.Load{}, &auth.Auth{}, &updates.Update{}, &partners.Partners{},
		&admin.Admin{}, &user_create.User{})

	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup.Static("/static", "static")
	r.Run()

	fmt.Println("Hell the world")
}
