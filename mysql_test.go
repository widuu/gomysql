package gomysql

import (
	"log"
	"testing"
)

func TestMysql(*testing.T) {
	c, err := SetConfig("./conf/conf.ini")
	if err != nil {
		log.Println(err)
	}
	//data := c.SetTable("user").FindOne()
	log.Println(c)
}
