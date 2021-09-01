package app

import (
	"encoding/json"
	"go-projects/http-kafka-producer/structures"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func receiveMessage(c echo.Context) (err error) {

	msg := new(structures.TestMessage)

	if err := c.Bind(msg); err != nil {
		log.Error("MalformedMessage", err)
		return c.JSON(http.StatusBadRequest, "malformed Message")
	}

	msgJSON, err := json.Marshal(msg)

	if err != nil {
		log.Errorf()
	}

	err := util.PublishToKafka("testTopic", msgJSON, []string{"localhost:9092"})

	if err != nil {
		log.Errorf("failed to publish to kafka topic: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "malformed Message")
	}

	log.Info("message has been successfully published to kafka")

	return c.JSON(http.StatusOK, "success")

}
