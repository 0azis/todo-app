package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
	"todo/internal/models"
	"todo/internal/pkg"
	"todo/internal/services"
)

type AuthControllers struct {
	UserServices services.UsersServices
}

func (ac AuthControllers) SignUp(c *fiber.Ctx) error {
	var credentials models.Credentials
	err := c.BodyParser(&credentials)
	if err != nil {
		return fiber.NewError(500, "Ошибка при получении данных")
	}
	isUserValid := models.IsUserValid{}.ValidCredentials(credentials)
	if isUserValid.Email != "" || isUserValid.Password != "" {
		return c.JSON(isUserValid)
	}

	hashedPassword, _ := pkg.Encode([]byte(credentials.Password))
	newUser := models.User{
		Email:    credentials.Email,
		Password: string(hashedPassword),
	}

	newUserID, err := ac.UserServices.InsertOne(newUser)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(500, "Ошибка при создании пользователя")
	}

	accessToken, err := pkg.CreateAccessToken(newUserID)
	if err != nil {
		return fiber.NewError(500, "Ошибка при создании пользователя")
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   accessToken,
		Expires: time.Now().Add(time.Minute * 10),
	})

	return fiber.NewError(200, "Успешная регистрация")
}

func (ac AuthControllers) SignIn(c *fiber.Ctx) error {
	var credentials models.Credentials
	err := c.BodyParser(&credentials)
	if err != nil {
		return fiber.NewError(500, "Ошибка при получении данных")
	}

	isUser, err := ac.UserServices.GetByEmail(credentials.Email)
	if err != nil {
		return fiber.NewError(500, "Ошибка при получении данных")
	}

	if err = pkg.Decode([]byte(isUser.Password), []byte(credentials.Password)); err != nil {
		return fiber.NewError(401, "Неправильные данные пользователя")
	}

	accessToken, err := pkg.CreateAccessToken(isUser.ID)
	if err != nil {
		return fiber.NewError(500, "Ошибка при создании пользователя")
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   accessToken,
		Expires: time.Now().Add(time.Minute * 10),
	})
	return fiber.NewError(200, "Успешный вход")
}
