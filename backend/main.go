package main

import (
	"example.com/m/v2/initializers"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	// r := gin.Default()

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "pong"})
	// })

	// r.Run()
}
