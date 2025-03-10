// @title			Task API
// @version			1.0
// @description	 	Esta é uma API de gerenciamento de tarefas.
// @host			localhost:8080
// @BasePath		/
package main

import (
	"context"
	"go-crud-api/models"
	"net/http"
	"time"

	_ "go-crud-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
func createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary 	Lista todas as tarefas
// @Description Lista todas as tarefas
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} map[string]interface{}
// @Failure 	400 {object} map[string]string
// @Failure 	500 {object} map[string]string
// @Router 		/tasks [get]
func listTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

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
func updateTask(c *gin.Context) {
	idParam := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, bson.M{"_id": taskID}, bson.M{"$set": task})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

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
func deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": taskID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

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

	router.POST("/tasks", createTask)
	router.GET("/tasks", listTasks)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.Run(":8080")
}
