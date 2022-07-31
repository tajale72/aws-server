package router

import (
	controller "interview/internal/controller"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	origin := c.Request.Header.Get("User-Agent")
	name := c.Param("name")
	res, err := controller.GetUser(name)
	if err != nil {
		log.Println("error from controller", err)
		c.JSON(http.StatusBadRequest, err)
	} else {
		log.Println("origin", origin)
		c.JSON(http.StatusOK, res)
	}
}

func InsertUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("error reading the request body\n", err)
		c.JSONP(http.StatusBadRequest, err)
	}
	res, err := controller.InsertUser(body)
	if err != nil {
		log.Println("error inserting the data\n", err)
		c.JSONP(http.StatusBadRequest, err)
	} else {
		log.Println("inserted users")
		c.JSONP(http.StatusOK, res)
	}

}

func GetAllUser(c *gin.Context) {
	res, err := controller.GetAllUser()
	if err != nil {
		log.Println("error from controller", err)
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, res)
}
