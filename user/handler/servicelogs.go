package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/go-microservices/user/repository"
	"github.com/tshubham7/go-microservices/user/services"
)

type logs struct {
	lr repository.LogRepo
	l  *log.Logger
}

// LogHandler ...
type LogHandler interface {
	// list service logs
	List() gin.HandlerFunc
}

// NewLogHandler ...
func NewLogHandler(r repository.LogRepo, l *log.Logger) LogHandler {
	return &logs{r, l}
}

// List ...
func (lg logs) List() gin.HandlerFunc {
	sr := services.NewLogService(lg.lr, lg.l)

	return func(c *gin.Context) {

		lmt, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: limit",
				"error":   err.Error(),
			})
			return
		}

		ofs, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: offset",
				"error":   err.Error(),
			})
			return
		}

		list, err := sr.ListAll(int32(lmt), int32(ofs))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to fetch logs",
				"error":   err.Error(),
			})
			return
		}

		resp := ListResponse{
			Limit: int32(lmt), Offset: int32(ofs), Count: int32(len(list)), Result: list}

		c.JSON(http.StatusOK, resp)
	}
}
