package handler

import (
	"newsportal-backend/internal/adapter/handler/request"
	"newsportal-backend/internal/adapter/handler/response"
	"newsportal-backend/internal/core/domain/entity"
	"newsportal-backend/internal/core/service"
	validatorLib "newsportal-backend/lib/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var err error
var code string
var errorResp response.ErrorResponseDefault

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
}

// Login implements AuthHandler.
func (ah *authHandler) Login(c *fiber.Ctx) error {
	req := request.LoginRequest{}
	resp := response.SuccessAuthResponse{}

	if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER] Login - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(req); err != nil {
		code = "[HANDLER] Login - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqLogin := entity.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := ah.authService.GetUserByEmail(c.Context(), reqLogin)
	if err != nil {
		code = "[HANDLER] Login - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		if err.Error() == "invalid password" {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	resp.Meta.Status = true
	resp.Meta.Message = "Login Successful"
	resp.AccessToken = result.AccessToken
	resp.ExpiresAt = result.ExpiresAt

	return c.JSON(resp)
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}
