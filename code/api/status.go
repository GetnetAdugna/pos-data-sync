package api

import (
	"net/http"
	"time"

	"serveos-datasync/code/services/mqtt"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func InitStatusRouter(router *gin.Engine) {
	startTime = time.Now()
	atomic.StoreInt32(&mqtt.MqttConnected, 1)

	router.GET("/status", getStatus)
}

type StatusResponse struct {
	RunningSince               string `json:"runningSince"`
	RunningFor                 string `json:"runningFor"`
	MQTTConnectionStatus       string `json:"mqttConnectionStatus"`
	PublishTopic               string `json:"publishTopic"`
	SubscribeTopic             string `json:"subscribeTopic"`
	MQTTDisconnectsLastHour    int    `json:"mqttDisconnectsLastHour"`
	MQTTDisconnectsLast24Hours int    `json:"mqttDisconnectsLast24Hours"`
}

func getStatus(c *gin.Context) {
	runningFor := time.Since(startTime)

	response := StatusResponse{
		RunningSince:               startTime.Format(time.RFC3339),
		RunningFor:                 runningFor.String(),
		MQTTConnectionStatus:       mqtt.GetMQTTConnectionStatus(),
		PublishTopic:               "your/publish/topic",
		SubscribeTopic:             "your/subscribe/topic",
		MQTTDisconnectsLastHour:    mqtt.GetMQTTDisconnectsLastHour(),
		MQTTDisconnectsLast24Hours: mqtt.GetMQTTDisconnectsLast24Hours(),
	}

	c.JSON(http.StatusOK, response)
}
