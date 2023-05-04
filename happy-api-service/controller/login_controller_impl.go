package controller

import (
	"encoding/json"
	"fmt"
	"happy-api-service/config"
	"happy-api-service/helper"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginControllerImpl struct {
	Config   *util.Config
	Validate *validator.Validate
}

func NewLoginController(config *util.Config) LoginController {
	return &LoginControllerImpl{
		Config: config,
	}
}

func (controller *LoginControllerImpl) Login(fiberCtx *fiber.Ctx) error {
	reqBody := fiberCtx.Request().Body()

	loginCreateRequest := web.LoginCreateRequest{}
	err := json.Unmarshal(reqBody, &loginCreateRequest)
	helper.DoPanicIfError(err)

	if err := controller.Validate.Struct(loginCreateRequest); err != nil {
		webResponse := web.WebResponse[string]{
			Code:   fiber.StatusBadRequest,
			Status: "failed",
			Data:   "request data is invalid",
		}

		fiberCtx.Response().SetStatusCode(fiber.StatusBadRequest)
		return fiberCtx.JSON(webResponse)
	}

	httpUrlGetUser := fmt.Sprintf("%s/users/%s", controller.Config.HappyUserServiceUrl, loginCreateRequest.UserName)
	userSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetUser, nil)

	if userSvcRes.Status == "failed" {
		userSvcRes.Data = "username account is not registered"

		fiberCtx.Response().SetStatusCode(fiber.StatusNotFound)
		return fiberCtx.JSON(userSvcRes)
	}

	userSvcResData := web.UserServiceResponse{}
	helper.JoinResponse(&userSvcRes.Data, &userSvcResData)

	reqPasswordBytes := []byte(loginCreateRequest.Password)
	hashedPasswordBytes := []byte(userSvcResData.Password)

	if err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, reqPasswordBytes); err != nil {
		webResponse := web.WebResponse[string]{
			Code:   fiber.StatusBadRequest,
			Status: "failed",
			Data:   "password is unmatched",
		}

		fiberCtx.Response().SetStatusCode(fiber.StatusBadRequest)
		return fiberCtx.JSON(webResponse)
	}

	jwtExpTime := time.Now().Add(time.Hour * 24)
	jwtClaim := &config.JwtClaim{
		UserName: loginCreateRequest.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "happy-service",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
		},
	}

	jwtTokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	jwtToken, err := jwtTokenAlgorithm.SignedString([]byte(controller.Config.JwtSecretKey))

	if err != nil {
		webResponse := web.WebResponse[string]{
			Code:   fiber.StatusInternalServerError,
			Status: "failed",
			Data:   "something went wrong",
		}

		fiberCtx.Response().SetStatusCode(fiber.StatusInternalServerError)
		return fiberCtx.JSON(webResponse)
	}

	jwtCookie := &fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    jwtToken,
		HTTPOnly: true,
	}

	idUserCookie := &fiber.Cookie{
		Name:     "id_user",
		Path:     "/",
		Value:    strconv.Itoa(int(userSvcResData.Id)),
		HTTPOnly: true,
	}

	fiberCtx.Cookie(jwtCookie)
	fiberCtx.Cookie(idUserCookie)

	webResponse := web.WebResponse[web.LoginWebResponse]{
		Code:   200,
		Status: "success",
		Data: web.LoginWebResponse{
			Id:       userSvcResData.Id,
			FullName: userSvcResData.FullName,
			UserName: userSvcResData.UserName,
			Email:    userSvcResData.Email,
		},
	}

	fiberCtx.Response().SetStatusCode(fiber.StatusOK)
	return fiberCtx.JSON(webResponse)
}
