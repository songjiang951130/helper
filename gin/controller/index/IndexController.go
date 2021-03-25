package index

import (
	"helper/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	user := service.GetIndexInfo()
	c.HTML(http.StatusOK, "index.html", user)
}
