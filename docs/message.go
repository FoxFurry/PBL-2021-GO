package docs

import (
	"foxy/internal/domain/dto"
	"foxy/internal/domain/entity"
)

// swagger:route POST /message/create/{userID} Message createMessageRequestWrapper
// Create a message by a specific user
// responses:
//   200: createMessageSuccessResponseWrapper

// Create message request should include roomID in which message will be posted and stringified data
// swagger:parameters createMessageRequestWrapper
type createMessageRequestWrapper struct {
	// in:body
	Body dto.MessageCreate
}

// Response contains ID of newly created message
// swagger:response createMessageSuccessResponseWrapper
type createMessageSuccessResponseWrapper struct {
	// in:body
	Body struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}
}

//*************************************************************************************************
//*************************************************************************************************

// swagger:route GET /message/{roomID} Message getMessageRequestWrapper
// Get all message from specific room
// responses:
//   200: getMessageSuccessResponseWrapper

// Empty body
//swagger:parameters getMessageRequestWrapper
type getMessageRequestWrapper struct {
	// in:body
}

// Response contains array of messages from specific room
// swagger:response getMessageSuccessResponseWrapper
type getMessageSuccessResponseWrapper struct {
	// in:body
	Body struct {
		Data []entity.Message `json:"data"`
	}
}
