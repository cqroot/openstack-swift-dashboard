package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"

	"github.com/cqroot/openstack-swift-dashboard/models"
)

func GetDiskList(c *gin.Context) {
	descStr := c.Query("desc")
	desc := false
	if strings.EqualFold(descStr, "true") {
		desc = true
	}

	disks, err := models.DiskList(
		cast.ToUint(c.Param("id")),
		cast.ToInt(c.Query("limit")),
		cast.ToInt(c.Query("offset")),
		desc,
	)
	if err != nil {
		log.Debug().Err(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, disks)
	}
}
