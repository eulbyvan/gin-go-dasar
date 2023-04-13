/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu Apr 13 2023 9:49:00 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// buat sebuah router gin baru
	r := gin.Default()

	// definisikan routes
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s", name)})
	})

	r.POST("/hai", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hai, world!"})
	})

	// jalankan server
	r.Run(":8080")
}