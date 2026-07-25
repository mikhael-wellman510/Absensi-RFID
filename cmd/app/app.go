package app

import (
	"attendance-api/config"
	"attendance-api/internal/adapters/controllers"
	"attendance-api/internal/adapters/middleware"
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/enums"
	"attendance-api/internal/seeders"
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

	seeders.SeedSuperAdmin(db)
}

func (app *App) Routes() {

	router := gin.Default()

	baseUrl := fmt.Sprintf("%s/v%d", constants.ApiPrefix, constants.ApiVersion)

	auditLogRepository := repositories.NewAuditLogRepository(app.Db)
	auditLogUseCase := usecases.NewAuditLogService(auditLogRepository)

	userSessionRepository := repositories.NewUserSessionRepository(app.Db)
	userSessionUseCase := usecases.NewUserSessionService(userSessionRepository)

	authRepo := repositories.NewAuthRepository(app.Db)
	authUseCase := usecases.NewAuthService(authRepo, auditLogUseCase, userSessionUseCase)
	authController := controllers.NewAuthController(authUseCase)

	// Public Routes
	publicRoutes := router.Group(baseUrl)
	authRoutesPublic := publicRoutes.Group("/auth")
	{
		authRoutesPublic.POST("/login", authController.Login)
		authRoutesPublic.POST("/refreshToken", authController.RefreshToken)
		authRoutesPublic.POST("/forgotPassword", authController.ForgotPassword)
		authRoutesPublic.POST("/resetPassword", authController.ResetPassword)
		authRoutesPublic.GET("/verifyEmail", authController.VerifyEmail)
	}

	// Protected Routes
	protectedRoutes := router.Group(baseUrl)
	protectedRoutes.Use(middleware.AuthMiddleware(authRepo))

	authRoutesProtected := protectedRoutes.Group("/auth")
	{
		authRoutesProtected.GET("/getMe", authController.GetMe)
		authRoutesProtected.POST("/logout", authController.Logout)
		authRoutesProtected.POST("/logoutAll", authController.LogoutAll)
		authRoutesProtected.POST("/createUser", authController.CreateUser)
	}

	gradeRepo := repositories.NewGradeRepository(app.Db)
	gradeUseCase := usecases.NewGradeService(gradeRepo)
	gradeController := controllers.NewGradeController(gradeUseCase)

	gradeRoutesProtected := protectedRoutes.Group("/grade")
	{
		gradeRoutesProtected.POST("/createGrade", middleware.RequireRoles(enums.SuperAdmin, enums.Teacher), gradeController.CreateGrade)
		gradeRoutesProtected.GET("/findGradeById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.Teacher), gradeController.FindGradeById)
	}

	schoolRepo := repositories.NewSchoolRepository(app.Db)
	schoolUseCase := usecases.NewSchoolService(schoolRepo)
	schoolController := controllers.NewSchoolController(schoolUseCase)

	schoolRoutesProtected := protectedRoutes.Group("/school")
	{
		schoolRoutesProtected.POST("/createSchool", middleware.RequireRoles(enums.SuperAdmin), schoolController.CreateSchool)
		schoolRoutesProtected.GET("/findSchoolById/:id", middleware.RequireRoles(enums.SuperAdmin), schoolController.FindSchoolById)
	}

	userRepo := repositories.NewUserRepository(app.Db)
	userUseCase := usecases.NewUserService(userRepo)
	userController := controllers.NewUserController(userUseCase)

	userRoutesProtected := protectedRoutes.Group("/user")
	{
		userRoutesProtected.POST("/createUser", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), userController.CreateUser)
		userRoutesProtected.GET("/findUserById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), userController.FindUserById)
	}

	parentRepo := repositories.NewParentRepository(app.Db)
	parentUseCase := usecases.NewParentService(parentRepo, userUseCase)
	parentController := controllers.NewParentController(parentUseCase)

	parentRoutesProtected := protectedRoutes.Group("/parent")
	{
		parentRoutesProtected.POST("/createParent", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), parentController.CreateParent)
		parentRoutesProtected.GET("/findParentById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), parentController.FindParentById)
	}

	teacherRepo := repositories.NewTeacherRepository(app.Db)
	teacherUseCase := usecases.NewTeacherService(teacherRepo, userUseCase, schoolUseCase)
	teacherController := controllers.NewTeacherController(teacherUseCase)

	teacherRoutesProtected := protectedRoutes.Group("/teacher")
	{
		teacherRoutesProtected.POST("/createTeacher", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), teacherController.CreateTeacher)
		teacherRoutesProtected.GET("/findTeacherById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), teacherController.FindTeacherById)
	}

	studentRepo := repositories.NewStudentRepository(app.Db)
	studentUseCase := usecases.NewStudentService(studentRepo, schoolUseCase)
	studentController := controllers.NewStudentController(studentUseCase)

	studentRoutesProtected := protectedRoutes.Group("/student")
	{
		studentRoutesProtected.POST("/createStudent", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), studentController.CreateStudent)
		studentRoutesProtected.GET("/findStudentById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), studentController.FindStudentById)
	}

	studentParentRepo := repositories.NewStudentParentRepository(app.Db)
	studentParentUseCase := usecases.NewStudentParentService(studentParentRepo, studentUseCase, parentUseCase)
	studentParentController := controllers.NewStudentParentController(studentParentUseCase)

	studentParentRoutesProtected := protectedRoutes.Group("/studentparent")
	{
		studentParentRoutesProtected.POST("/createStudentParent", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), studentParentController.CreateStudentParent)
		studentParentRoutesProtected.GET("/findStudentParentById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), studentParentController.FindStudentParentById)
	}

	academicYearRepo := repositories.NewAcademicYearRepository(app.Db)
	academicYearUseCase := usecases.NewAcademicYearService(academicYearRepo, schoolUseCase)
	academicYearController := controllers.NewAcademicYearController(academicYearUseCase)

	academicYearRoutesProtected := protectedRoutes.Group("/academicyear")
	{
		academicYearRoutesProtected.POST("/createAcademicYear", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), academicYearController.CreateAcademicYear)
		academicYearRoutesProtected.GET("/findAcademicYearById/:id", middleware.RequireRoles(enums.SuperAdmin, enums.SchoolAdmin), academicYearController.FindAcademicYear)

	}

	//emailUseCase := usecases.NewEmailService()
	//emailController := controllers.NewEmailController(emailUseCase)
	//
	//emailRoutes := router.Group(fmt.Sprintf("%s/email", baseUrl))
	//emailRoutes.POST("/sendEmail", emailController.SendEmail)

	app.Router = router
}

func (app *App) Run() {
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	err := app.Router.Run(port)
	if err != nil {
		return
	}

}
