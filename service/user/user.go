package user

import (
	"chat-service/Exception"
	"chat-service/models/response"
	"chat-service/models/tables"
	"chat-service/repository/user"
	"chat-service/requests"
	"chat-service/system"
	"strings"
)

type userService struct {
}

var Service userService

func (receiver userService) CreateUser(request requests.AddUserRequest) {
	defer Exception.GetErrorJson(system.Context)

	if isExist := user.Repository.FindByUsernameIsExist(request.Username); isExist {
		system.Context.JSON(response.BAD_REQUEST_CODE, response.FailureResponse(response.BAD_REQUEST_STATUS, "Username was already", request))
		return
	}

	splitFullName := strings.Split(request.FullName, " ")
	lastName := ""
	for i, s := range splitFullName {
		if i != 0 {
			lastName += s + " "
		}
	}
	result := user.Repository.Save(tables.User{
		Email: request.Username,
		Name:  splitFullName[0],
	})

	system.Context.JSON(response.SUCCESS_CODE, response.SuccessResponse(true, "Success create", result))
}
