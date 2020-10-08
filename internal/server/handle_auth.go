package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/cronui/cronui/internal/entity"
)

func (s *Server) handleAuthToken(c *fiber.Ctx) error {
	r := new(apiAuthTokenRequest)
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := r.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var user entity.User
	if err := s.db.Where("name = ?", r.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(jwt.MapClaims{
				"field": "username",
				"message": "User with this username does not exists",
			})
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(r.Password)); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(jwt.MapClaims{
			"field": "password",
			"message": "Wrong password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}

type apiAuthTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *apiAuthTokenRequest) validate() error {
	if r.Username == "" {
		return fmt.Errorf("field 'username' cannot be empty")
	}
	if r.Password == "" {
		return fmt.Errorf("field 'password' cannot be empty")
	}
	return nil
}
