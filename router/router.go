package router

import (
	"final-project/controller"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.PUT("/", middleware.Authentication(), controller.UpdateDataUser)
		userRouter.DELETE("/", middleware.Authentication(), controller.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetAllPhoto)
		photoRouter.GET("/:PhotoId", controller.GetPhotoById)
		photoRouter.PUT("/:PhotoId",  controller.UpdatePhoto)
		photoRouter.DELETE("/:PhotoId",  controller.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.GET("/", controller.GetAllComment)
		commentRouter.GET("/:CommentId", controller.GetCommentById)
		commentRouter.PUT("/:CommentId",  controller.UpdateComment)
		commentRouter.DELETE("/:CommentId", controller.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.GET("/", controller.GetAllSocialMedia)
		socialMediaRouter.GET("/:SocialMediaId", controller.GetSocialMediaById)
		socialMediaRouter.PUT("/:SocialMediaId",  controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:SocialMediaId",  controller.DeleteSocialMedia)
	}
	return r
}