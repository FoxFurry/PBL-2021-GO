package docs

import (
	"foxy/internal/domain/dto"
	_ "github.com/pdrum/swagger-automation/api"
)

// swagger:route POST /user/register User registerUser
// Registers a new user in database
// responses:
//   200: userRegisterSuccessResponseWrapper
//	 409: userRegisterAlreadyExistsWrapper

// Request body should consist of full name, email and password of the new user
// swagger:parameters registerUser
type userRegisterRequestWrapper struct {
	// in:body
	Body dto.UserRegister
}

// Response body contains single integer: ID of newly created user
// swagger:response userRegisterSuccessResponseWrapper
type userRegisterSuccessResponseWrapper struct {
	// in:body
	Body struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}
}

// If user already exists, this response will be sent
// swagger:response userRegisterAlreadyExistsWrapper
type userRegisterAlreadyExistsWrapper struct {
	// in: body
	Body struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
}

//*************************************************************************************************
//*************************************************************************************************

// swagger:route POST /user/login User loginUser
// User authentication endpoint
// responses:
//   200: userLoginSuccessResponseWrapper
//	 401: userLoginWrongCredentialsWrapper

// Request body should consist of email and password in order to authenticate
// swagger:parameters loginUser
type userLoginRequestWrapper struct {
	// in:body
	Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

// Response body contains single integer: ID of authenticated user
// swagger:response userLoginSuccessResponseWrapper
type userLoginSuccessResponseWrapper struct {
	// in:body
	Body struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}
}

// If credentials are wrong, this response will be sent
// swagger:response userLoginWrongCredentialsWrapper
type userLoginWrongCredentialsWrapper struct {
	// in: body
	Body struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
}
