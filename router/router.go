package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/novaladip/geldstroom-api-go/auth"
)

type Router struct {
	DB     *sql.DB
	R      *gin.Engine
	Secret string
}

// Initializing routes
func (r Router) Init() {
	auth := &auth.Authhentication{
		Db:     r.DB,
		Secret: r.Secret,
	}

	authRoutes := r.R.Group("/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}
}
