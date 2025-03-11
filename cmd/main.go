// @title			Task API
// @version			1.0
// @description	 	Esta é uma API de gerenciamento de tarefas.
// @host			localhost:8080
// @BasePath		/
package main

import (
	"time"

	//_ "go-crud-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

// @Summary 	Cria uma nova tarefa
// @Description Cria uma nova terefa
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		task body models.Task true "Dados da Tarefa"
// @Success 	201 {object} map[string]interface{}
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/tasks [post]

// @Summary 	Lista todas as tarefas
// @Description Lista todas as tarefas
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} map[string]interface{}
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/tasks [get]

// @Summary 	Atualiza uma tarefa
// @Description Atualiza uma terefa
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID da Tarefa"
// @Param 		task body models.Task true "Dados atualizados da Tarefa"
// @Success 	200 {object} map[string]string
// @Failure 	400 {object} map[string]string
// @Failure 	404 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/tasks [put]

// @Summary 	Deleta uma tarefa
// @Description Deleta uma terefa
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID da Tarefa"
// @Success 	200 {object} map[string]string
// @Failure 	400 {object} map[string]string
// @Failure 	404 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/tasks [delete]

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
