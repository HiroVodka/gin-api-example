package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"api-tutorial/internal/db"
	"api-tutorial/internal/handler"
)

func main() {
	// DB接続
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Ginの設定
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		c.Next()
	})

	// ハンドラー初期化
	userHandler := handler.NewUserHandler(dbConn)

	// ルーティング設定
	users := r.Group("/users")
	{
		users.GET("", userHandler.ListUsers)
		users.GET("/:id", userHandler.GetUser)
		// 他のエンドポイントも同様に設定
	}

	r.Run(":4649")
}
