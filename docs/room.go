package docs

import (
	"foxy/internal/domain/dto"
	"foxy/internal/domain/entity"
)
import _ "github.com/pdrum/swagger-automation/api"

// swagger:route GET /room/{userID} Room getRoomsRequestWrapper
// Get all rooms for a specific user
// responses:
//   200: getRoomsSuccessResponseWrapper

// Empty body
// swagger:parameters getRoomsRequestWrapper
type getRoomsRequestWrapper struct {
	// in:body
}

// Response body contains an array of room objects
// swagger:response getRoomsSuccessResponseWrapper
type getRoomsSuccessResponseWrapper struct {
	// in:body
	Body struct {
		Data []entity.Room `json:"data"`
	}
}

//*************************************************************************************************
//*************************************************************************************************

// swagger:route POST /room/create/{userID} Room createRoomsRequestWrapper
// Create a room by a specific user
// User specified in URL will be admin of new room
// responses:
//   200: createRoomsSuccessResponseWrapper

// Create room request should only contain a name of new room
// swagger:parameters createRoomsRequestWrapper
type createRoomsRequestWrapper struct {
	// in:body
	Body dto.RoomCreate
}

// Response body contains an array of room objects
// swagger:response createRoomsSuccessResponseWrapper
type createRoomsSuccessResponseWrapper struct {
	// in:body
	Body struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}
}
