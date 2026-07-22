package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var tasks = []task{
	{ID: "1", Title: "Read a book", Priority: "Low", Status: "Needs Action", Description: ""},
	{ID: "2", Title: "Get groceries", Priority: "High", Status: "Needs Action", Description: "Fresh greens (spinach, lettuce)\nCooking veggies (onions, bell peppers, carrots, broccoli)\nFruits (bananas, apples, berries)\nHerbs & aromatics (garlic, ginger)"},
	{ID: "3", Title: "Solve Leetcode", Priority: "Medium", Status: "Needs Action", Description: ""},
}

func getAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusOK, newTask)
}

func findTaskById(c *gin.Context) {
	id := c.Param("id")
	_, task, err := helperFindTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task not Found"})
	}
	c.IndentedJSON(http.StatusOK, task)
}

func deleteTaskById(c *gin.Context) {
	id := c.Param("id")
	i, task, err := helperFindTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task not Found"})
	}
	tasks = append(tasks[:i], tasks[i+1:]...)
	c.IndentedJSON(http.StatusOK, task)
}

func completeTask(c *gin.Context) {
	id := c.Param("id")
	_, task, err := helperFindTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task not Found"})
	}
	task.Status = "Completed"
	c.IndentedJSON(http.StatusOK, task)
}

func modifyTask(c *gin.Context) {
	id := c.Param("id")
	_, task, err := helperFindTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not Found"})
		return
	}
	type TaskUpdate struct {
		Title       *string `json:"title"`
		Completed   *bool   `json:"completed"`
		Priority    *string `json:"priority"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
	}
	var newupdate TaskUpdate
	if err := c.BindJSON(&newupdate); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	task.Priority = *newupdate.Priority
	task.Status = *newupdate.Status
	task.Description = *newupdate.Description
	task.Title = *newupdate.Title

	c.IndentedJSON(http.StatusOK, task)
}

// HELPER HELPPP
func helperFindTask(id string) (int, *task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return i, &tasks[i], nil
		}
	}
	return -1, nil, errors.New("Task not found")
}

// ROUTER MAIN
func main() {
	router := gin.Default()
	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", findTaskById)
	router.POST("/tasks", addTask)
	router.DELETE("/tasks/:id", deleteTaskById)
	router.PATCH("/tasks/:id/complete", completeTask)
	router.PATCH("/tasks/:id/update", modifyTask)
	router.Run("localhost:6767")
}
