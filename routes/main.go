// routes/router.go

package routes

import (
	"os"

	"github.com/chiennd172002/golang-swagger/controllers"
	docs "github.com/chiennd172002/golang-swagger/docs"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupV1Router() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()
	// CORS setup
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000", "https://parameter-store-fe-golang.up.railway.app", "http://localhost:" + os.Getenv("PORT")},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * 30 * time.Hour,
	// }))

	// Setup routes for the API version 1
	v1 := r.Group("/api/v1")
	{
		// Setup routes for the Lecturers
		v1.GET("/lecturers", controllers.ListLectures)
		v1.GET("/lecturers/:id", controllers.GetLecture)
		v1.POST("/lecturers", controllers.CreateLecture)
		v1.PUT("/lecturers/:id", controllers.UpdateLecture)
		v1.DELETE("/lecturers/:id", controllers.DeleteLecture)
		// Setup routes for the Students
		v1.GET("/students", controllers.ListStudents)
		v1.GET("/students/:id", controllers.GetStudent)
		v1.POST("/students", controllers.CreateStudent)
		v1.PUT("/students/:id", controllers.UpdateStudent)
		v1.DELETE("/students/:id", controllers.DeleteStudent)
		// Setup routes for the Courses
		v1.GET("/courses", controllers.ListCourses)
		v1.GET("/courses/:id", controllers.GetCourse)
		v1.DELETE("/courses/:id", controllers.DeleteCourse)

	}
	// Swagger setup
	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Title = "Golang swagger Backend API"
		docs.SwaggerInfo.Description = "This is a simple API for Golang swagger Backend."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else if gin.Mode() == gin.ReleaseMode {
		docs.SwaggerInfo.Title = "Golang swagger Backend API"
		docs.SwaggerInfo.Description = "This is a simple API for Golang swagger Backend."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
		docs.SwaggerInfo.Schemes = []string{"http"}
	}
	if os.Getenv("ENABLE_SWAGGER") == "true" {
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return r
}
