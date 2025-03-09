package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	Service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		Service: service,
	}
}

func (h *handler) GetTask(c *gin.Context) {
	tasks, err := h.Service.GetTaskService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to display tasks")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully display tasks", "data": tasks})
}

func (h *handler) GetTaskById(c *gin.Context) {
	stringId := c.Param("id")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		return
	}

	task, err := h.Service.GetTaskByIdService(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Task not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully display task", "data": task})
}

func (h *handler) CreateTask(c *gin.Context) {
	var newTask *TaskRequest

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	if err := h.Service.CreateTaskService(newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create task"})
}

func (h *handler) UpdateTask(c *gin.Context) {
	stringId := c.Param("id")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		return
	}

	var updatedTask TaskRequest

	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	if err := h.Service.UpdateTaskService(id, &updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update task"})
}

func (h *handler) DeleteTask(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		return
	}

	if err := h.Service.DeleteTaskService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete task"})
}
