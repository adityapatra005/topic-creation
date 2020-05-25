package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"pub_sub/components"
	"pub_sub/kafka"
	"pub_sub/topicCreation"

	"github.com/gin-gonic/gin"
)

func DonePost() gin.HandlerFunc {
	return func(c *gin.Context) {

		var body components.Request_body
		c.BindJSON(&body)
		e, err := json.Marshal(body)
		key := body.Key
		fmt.Println(err)

		c.JSON(200, gin.H{
			"status": "Transaction Successful",
		})
		err1 := kafka.Push(context.Background(), []byte(key), []byte(e))
		topicCreation.Topic(body.Topic_name)
		fmt.Println(err1)

	}
}
