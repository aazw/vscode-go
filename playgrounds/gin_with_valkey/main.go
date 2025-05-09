package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	store, _ := redis.NewStore(10, "tcp", "host.docker.internal:6379", "", "", []byte("secret")) // []byte("secret")がなくても起動するが、[]byte("secret")がないとCookieが設定されない
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", func(c *gin.Context) {

		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	router.Run("0.0.0.0:8080")
}
