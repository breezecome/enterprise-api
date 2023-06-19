package user

import (
	flashModel "enterprise-api/app/models/flash"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListFlash(c *gin.Context) {
	var listFlashIn schemas.ListFlashIn
	if err := c.ShouldBindQuery(&listFlashIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	list, err := flashModel.List(true, listFlashIn.Page, listFlashIn.Max)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}
