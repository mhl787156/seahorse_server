package users

import (
	"github.com/mhl787156/seahorse_server/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	oID := bson.ObjectIdHex(c.Param("_id"))
	user := models.User{}
	err := db.C(models.CollectionUsers).FindId(oID).One(&user)
	if err != nil {
		c.Error(err)
	}
	log.Print(user.Name)
	c.JSON(200, user)
}
