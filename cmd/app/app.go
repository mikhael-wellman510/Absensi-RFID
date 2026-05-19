package app

import (
	"attendance-api/config"
	"attendance-api/internal/adapters/controllers"
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/usecases"
	"attendance-api/migrations"
	"attendance-api/pkg/constants"
	"attendance-api/pkg/drivers/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Db     *gorm.DB
	Router *gin.Engine
}

func (app *App) ConnectDb() {
	db := sql.OpenDbConnection()
	app.Db = db
	err := migrations.Migrate(db)

	if err != nil {
		log.Fatal("Migrations Failed : -> ", err)
	}

}

func (app *App) Routes() {

	router := gin.Default()

	baseUrl := fmt.Sprintf("%s/v%d", constants.ApiPrevix, constants.ApiVersion)

	schoolRepo := repositories.NewSchoolRepository(app.Db)
	schoolUseCase := usecases.NewSchoolService(schoolRepo)
	schoolController := controllers.NewSchoolController(schoolUseCase)

	schoolRoutes := router.Group(fmt.Sprintf("%s/school", baseUrl))

	schoolRoutes.POST("/createSchool", schoolController.CreateSchool)

	app.Router = router
}

func (app *App) Run() {
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	app.Router.Run(port)

}
