package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/database"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/helpers"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/models"
)

type PhotoController interface {
	GetByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type photoController struct {
	photoHelper helpers.PhotoHelper
	jwtHelper   helpers.JWTHelper
}

func NewPhotoController(photoHelper helpers.PhotoHelper, jwtHelper helpers.JWTHelper) PhotoController {
	return &photoController{
		photoHelper: photoHelper,
		jwtHelper:   jwtHelper,
	}
}

func (c *photoController) GetByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var photo models.Photo = c.photoHelper.GetByID(id)
	if (photo == models.Photo{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", photo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *photoController) Insert(context *gin.Context) {
	var photoCreateData database.PhotoCreateData
	errData := context.ShouldBind(&photoCreateData)
	if errData != nil {
		res := helpers.BuildErrorResponse("Failed to proccess req", errData.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			photoCreateData.UserID = uint16(convertUserID)
		}
		result := c.photoHelper.Insert(photoCreateData)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *photoController) Update(context *gin.Context) {
	var photoUpdateData database.PhotoUpdateData
	errData := context.ShouldBind(&photoUpdateData)
	if errData != nil {
		res := helpers.BuildErrorResponse("Failed to proccess req", errData.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtHelper.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoHelper.IsAllowedToEdit(userID, photoUpdateData.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			photoUpdateData.UserID = id
		}
		result := c.photoHelper.Update(photoUpdateData)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) Delete(context *gin.Context) {
	var photo models.Photo
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed you get id", "No params were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	photo.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtHelper.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoHelper.IsAllowedToEdit(userID, photo.ID) {
		c.photoHelper.Delete(photo)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtHelper.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
