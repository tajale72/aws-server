package router

// import (
// 	db "interview/db"
// 	"io/ioutil"
// 	"net/http"
// )

// type Router struct {
// 	db *db.Business
// }

// func (a *Router) InsertUser(c *gin.Context) {
// 	body, err := ioutil.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	id, err := a.db.InsertUser(body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "data not inserted")
// 	}
// 	c.JSONP(http.StatusAccepted, id)
// }

// func (a *Router) GetUser(c *gin.Context) {
// 	res, err := a.db.GetUser()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	c.JSONP(http.StatusAccepted, res)
// }
