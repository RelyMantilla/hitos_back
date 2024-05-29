package main

import (
	"fmt"
	"hitos_back/controllers"
	"hitos_back/models"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// r := gin.New()

	models.Connectmodels()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	router := gin.Default()
	// config := cors.Config()
	// config.AllowOrigins = []string{"*"}
	// config.AllowCredentials = true
	// config.AllowAllOrigins = true

	// config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// return origin == "https://github.com"
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	// router.Use(cors.Default())

	// router.Use(middlewares.CORSMiddleware())

	// Api
	v1 := router.Group("/api")
	// v1.Use(middlewares.CORSMiddleware())
	{
		v1.POST("login", controllers.Login)
		v1.POST("register", controllers.Register)
	}

	admin := router.Group("/api")

	//admin.Use(middlewares.JwtAuthMiddleware())
	{
		admin.GET("getFamily", controllers.GetFamily)
		admin.POST("setFamily", controllers.SetFamily)
		admin.GET("getPillar", controllers.GetPillar)
		admin.POST("setPillar", controllers.SetPillar)
		admin.POST("setCompetence", controllers.SetCompetence)
		admin.GET("getCompetence", controllers.GetCompetence)
		admin.POST("setSkill", controllers.SetSkill)
		admin.GET("getSkill", controllers.GetSkill)
		admin.GET("person", controllers.GetPersonByID)
		admin.POST("getPerson", controllers.GetPerson)

		admin.GET("user", controllers.CurrentUser)

	}

	forum := router.Group("/api")

	//forum.Use(middlewares.JwtAuthMiddleware())
	{
		forum.GET("tags", controllers.GetTags)
		forum.GET("tag", controllers.GetTag)
		forum.POST("tag", controllers.SetTag)

		forum.GET("questions", controllers.GetQuestions)
		forum.GET("question", controllers.GetQuestion)
		forum.POST("question", controllers.SetQuestion)
		forum.POST("questionSolved", controllers.QuestionSolved)

		forum.GET("answers", controllers.GetAnswers)
		forum.GET("answer", controllers.GetAnswer)
		forum.POST("answer", controllers.SetAnswer)

		forum.GET("comments", controllers.GetComments)
		forum.GET("comment", controllers.GetComment)
		forum.POST("comment", controllers.SetComment)

		forum.POST("like", controllers.Like)

	}

	port := os.Getenv("port")
	_ = router.Run(":" + port)
}
