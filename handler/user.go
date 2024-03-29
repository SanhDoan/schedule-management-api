package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management-api/common"
	"schedule-management-api/model"
	"schedule-management-api/repository"
)

type UserHandler struct {
	repo repository.UserRepo
}

func (u UserHandler) GetList(c echo.Context) (err error) {
	var users []model.User
	err = u.repo.GetListUsers(&users)
	if err != nil {
		return RespondToClient(c, common.ERROR_GET_ROW_FROM_DB, common.MSG_GET_ROW_FROM_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, users)
}

func (u UserHandler) Update(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	user := new(model.User)
	if err = c.Bind(&user); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}
	user.ID = id
	err = u.repo.UpdateUser(user)

	if err != nil {
		return RespondToClient(c, common.ERROR_UPDATE_ROW_IN_DB, common.MSG_UPDATE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, user)
}

func (u UserHandler) Delete(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	err = u.repo.DeleteUser(id)
	if err != nil {
		return RespondToClient(c, common.ERROR_DELETE_ROW_IN_DB, common.MSG_DELETE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, true)
}