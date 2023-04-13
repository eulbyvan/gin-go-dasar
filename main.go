/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu Apr 13 2023 9:49:00 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// buat sebuah router gin baru
	r := gin.Default()

	// buat sebuah router group "/api/v1"
	apiRoutes := r.Group("/api/v1")

	// definisikan routes yang termasuk ke dalam group /api/v1
	apiRoutes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello from GET /api/v1"})
	})

	apiRoutes.POST("/users", func(ctx *gin.Context) {
		var requestBody struct {
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
		}
		err := ctx.ShouldBindJSON(&requestBody)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusCreated, gin.H{"code": "01", "status": "created", "message": "user created", "data": requestBody})
	})

	// jalankan server
	r.Run(":8080")
}