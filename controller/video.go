package controller

import (
	"github.com/gin-gonic/gin"
	"go/help"
	"go/live"
	"strconv"
)

func UserRoom(c *gin.Context){
	roomID, _ := strconv.Atoi( c.Query("roomID") )
	userID := 1

	client := live.CreateClient(userID, nil)
	live.LiveManager.Clients[userID] = client

	if room, ok := live.LiveManager.Rooms[roomID]; ok {
		room.Clients[userID] = client
	} else {

	}

	c.HTML(200,"userroom.html",gin.H{
		"roomID": roomID,
		"userID": userID,
	})
}

func CreateRoom(c *gin.Context) {
	roomID, err := strconv.Atoi( c.Query("roomID") )
	if err != nil {
		help.Log.Info("create room Atoi err:", err.Error())
	}

	room := live.CreateRoom(roomID, nil)
	live.LiveManager.Rooms[roomID] = room

	c.HTML(200,"liveroom.html",gin.H{
		"roomID": roomID,
	})
}