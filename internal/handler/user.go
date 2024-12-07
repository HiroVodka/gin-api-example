package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"api-tutorial/internal/db"
)

type UserHandler struct {
	q *db.Queries
}

func NewUserHandler(dbConn *sql.DB) *UserHandler {
	return &UserHandler{
		q: db.New(dbConn),
	}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.q.ListUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range users {
		users[i].Name = strings.ToValidUTF8(users[i].Name, "")
		users[i].Email = strings.ToValidUTF8(users[i].Email, "")
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.q.GetUser(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
} 