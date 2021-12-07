package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/cqroot/openstack-swift-dashboard/models"
)

type targetInput struct {
	Name           string
	Endpoint       string
	ScrapeInterval string
	ScrapeTimeout  string
}

func GetTargetList(c *gin.Context) {
	targets, err := models.TargetList()
	if err != nil {
		log.Debug().Err(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, targets)
	}
}

func PutTarget(c *gin.Context) {
	var input targetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target := models.Target{
		Name:           input.Name,
		Endpoint:       input.Endpoint,
		ScrapeInterval: input.ScrapeInterval,
		ScrapeTimeout:  input.ScrapeTimeout,
	}
	models.UpdateTarget(&target)

	c.JSON(http.StatusOK, gin.H{"data": target})
}
