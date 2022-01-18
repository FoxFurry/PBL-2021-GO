package controller

import (
	"log"
	"strconv"

	"foxy/internal/domain/dto"
	"foxy/internal/domain/entity"
	"foxy/internal/http/httperr"
	"foxy/internal/http/httpresponse"
	"foxy/internal/http/middleware"
	"foxy/internal/infrastructure/jwtu"
	"foxy/internal/service"
	"github.com/gin-gonic/gin"
)

type IController interface {
	GetRouter() *gin.Engine
}

type foxyController struct {
	foxyService service.IService
}

func NewFoxyController() IController {
	return &foxyController{
		foxyService: service.NewFoxyService(),
	}
}

func (c *foxyController) GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	router.POST("/user/register", c.createUser)
	router.POST("/user/login", c.loginUser)
	router.GET("/user/:userID", middleware.TokenAuthMiddleware(), c.getUser)
	router.GET("/users", middleware.TokenAuthMiddleware(), c.getUsers)

	router.GET("/room/:userID", middleware.TokenAuthMiddleware(), c.getUserRooms)
	router.POST("/room/create/:userID", middleware.TokenAuthMiddleware(), c.createRoom)
	router.POST("/room/add/:userID", middleware.TokenAuthMiddleware(), c.addParticipant)

	router.POST("/message/create/:userID", middleware.TokenAuthMiddleware(), c.sendMessage)
	router.GET("/message/:roomID", middleware.TokenAuthMiddleware(), c.getAllMessages)

	return router
}

func (c *foxyController) createUser(ctx *gin.Context) {
	var userHolder dto.UserRegister

	err := ctx.ShouldBindJSON(&userHolder)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Creating user: %s %s %s", userHolder.FullName, userHolder.Password, userHolder.Email)

	newID, err := c.foxyService.RegisterUser(userHolder)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	token, _ := jwtu.CreateToken(newID)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"id":    newID,
		"token": token,
	})
}

func (c *foxyController) loginUser(ctx *gin.Context) {
	var userHolder dto.UserRegister

	err := ctx.ShouldBindJSON(&userHolder)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Authenticating user: %s %s %s", userHolder.FullName, userHolder.Password, userHolder.Email)

	userID, err := c.foxyService.Authorize(userHolder)

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	token, _ := jwtu.CreateToken(userID)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"id":    userID,
		"token": token,
	})
}

func (c *foxyController) getUsers(ctx *gin.Context) {
	users, err := c.foxyService.GetUsers()

	log.Printf("Getting all users")

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, users)
}

func (c *foxyController) getUser(ctx *gin.Context) {
	idParam := ctx.Param("userID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Getting user %d", id)

	user, err := c.foxyService.GetUser(uint(id))

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, user)
}

func (c *foxyController) getUserRooms(ctx *gin.Context) {
	idParam := ctx.Param("userID")
	id, err := strconv.Atoi(idParam)

	log.Printf("Getting rooms for user %d", id)

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	userRooms, err := c.foxyService.GetUsersRooms(uint(id))

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, userRooms)
}

func (c *foxyController) createRoom(ctx *gin.Context) {
	idParam := ctx.Param("userID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	var roomHolder dto.RoomCreate

	err = ctx.ShouldBindJSON(&roomHolder)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Creating a room by user %d: %s", id, roomHolder.Name)

	newID, err := c.foxyService.CreateRoom(uint(id), entity.Room{
		Name: roomHolder.Name,
	})

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"id": newID,
	})
}

func (c *foxyController) sendMessage(ctx *gin.Context) {
	idParam := ctx.Param("userID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	var messageCreateRequest dto.MessageCreate

	err = ctx.ShouldBindJSON(&messageCreateRequest)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Sending message by user %d for room %d: %s", id, messageCreateRequest.RoomID, messageCreateRequest.Data)

	newID, err := c.foxyService.SendMessage(uint(id), messageCreateRequest.RoomID, messageCreateRequest.Data)

	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"id": newID,
	})
}

func (c *foxyController) getAllMessages(ctx *gin.Context) {
	idParam := ctx.Param("roomID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Getting all messages for room %d", id)

	userRooms, err := c.foxyService.GetAllMessages(uint(id))

	if err != nil {
		httpresponse.RespondInternalError(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, userRooms)
}

func (c *foxyController) addParticipant(ctx *gin.Context) {
	var addParticipantRequest dto.AddParticipant

	err := ctx.ShouldBindJSON(&addParticipantRequest)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	log.Printf("Addint user %s to room %d", addParticipantRequest.ParticipantEmail, addParticipantRequest.RoomID)

	newParticipantID, err := c.foxyService.AddParticipant(addParticipantRequest.ParticipantEmail, addParticipantRequest.RoomID)
	if err != nil {
		httperr.Handle(ctx, err)
		return
	}

	httpresponse.RespondOK(ctx, newParticipantID)
}
