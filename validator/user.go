package validator

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserCreate struct {
	Username string `validate:"email,max=50"`
	Password string `validate:"min=6,max=20"`
	Fullname string `validate:"required"`
}

func CreateUserValidator(c *fiber.Ctx) error {
	body := new(UserCreate)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	validate := validator.New()
	err := validate.Struct(body)
	errs := validationErrorFormat(err)
	if len(errs) > 0 {
		return c.Status(400).JSON(errs)
	}
	return c.Next()
}
