/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu Apr 13 2023 9:49:00 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// custom middleware
func customMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// tambah logic... (login, autentikasi, etc...)
		fmt.Println("eitss.. tidak semudah itu ferguso")
		// panggil next handler
		c.Next()
	}
}

func main() {
	// buat sebuah router gin baru
	r := gin.Default()

	// pakai custom middleware di main router
	r.Use(customMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "gokil"})
	})

	// definisikan router group /api/v1
	apiGroup := r.Group("/api/v1")

	// pakai middleware untuk router group
	apiGroup.Use(customMiddleware())

	// definisikan routes di dalam /api/v1
	apiGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "mantap"})
	})

	// jalankan server
	r.Run(":8080")
}