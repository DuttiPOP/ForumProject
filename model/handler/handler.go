package handler

import (
	"ForumProject/model/constants"
	"ForumProject/model/service"
	"github.com/gin-contrib/sessions"
	gormSessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Handler struct {
	validate *validator.Validate
	services *service.Service
	db       *gorm.DB
}

func NewHandler(services *service.Service, db *gorm.DB) *Handler {
	return &Handler{
		services: services,
		validate: validator.New(),
		db:       db,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	store := gormSessions.NewStore(h.db, true, []byte("2ee2df817db37e42682f8051dc5c2721"))
	router.Use(sessions.Sessions("session_id", store))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	api.Use(h.authMiddleWare())

	me := api.Group("/me")
	{
		me.GET("/", h.getMyProfile)
		me.PUT("/", h.updateMyProfile)
		post := me.Group("/post")
		{
			post.GET("/all", h.getMyPosts)
		}
	}

	user := api.Group("/user")
	{
		user.GET("/:"+constants.UserIDKey, h.getUserById)
		post := user.Group("/:" + constants.UserIDKey + "/post")
		{
			post.GET("/all", h.getUserPosts)
		}
	}

	post := api.Group("/post")
	{

		post.POST("/", h.createPost)
		post.GET("/:"+constants.PostIDKey, h.getPostById)
		post.PUT("/:"+constants.PostIDKey, h.updatePost)
		comment := post.Group("/:" + constants.PostIDKey + "/comment")
		{
			comment.GET("/all", h.getCommentsByPostId)
			comment.POST("/", h.createComment)
			comment.PUT("/:"+constants.CommentIDKey, h.updateComment)
		}
	}

	return router
}
