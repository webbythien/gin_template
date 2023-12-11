package controller

import (
	"admin-panel/v1/app/model"
	"admin-panel/v1/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(ctx *gin.Context) {
	var req model.SendMessageReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	//jsonData, err := json.Marshal(req)
	//if err != nil {
	//	log.Fatal("Fail convert to JSON")
	//}
	//if err := cache.RedisPublishMessage(req.RoomID, jsonData); err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}
	err := service.SendMessageTask(req.Message, req.RoomID)
	if err != nil {
		panic("fail to send message")
	}
	//emitter := cache.New(cache.Options{
	//	Host:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
	//	Password: config.RedisPass,
	//})
	//emitter.In(req.RoomID).Emit("message", req.Message)

	//fmt.Println(req)
	return

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
