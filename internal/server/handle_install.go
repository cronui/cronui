package server

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) handleInstallGet(c *fiber.Ctx) error {
	count, err := s.db.User.Query().Count(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"installed": count > 0,
	})
}

func (s *Server) handleInstallPost(c *fiber.Ctx) error {
	count, err := s.db.User.Query().Count(c.Context())
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user is already defined")
	}

	r := new(apiPostInstallRequest)
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := r.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), 14)
	if err != nil {
		return err
	}

	user, err := s.db.User.Create().
		SetUsername(r.Username).
		SetPassword(string(hash)).
		Save(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": user.ID,
	})
}

type apiPostInstallRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *apiPostInstallRequest) validate() error {
	if r.Username == "" {
		return fmt.Errorf("field 'username' cannot be empty")
	}
	if len(r.Username) < 4 {
		return fmt.Errorf("field 'username' must contain at least 4 characters")
	}
	if r.Password == "" {
		return fmt.Errorf("field 'password' cannot be empty")
	}
	if len(r.Password) < 6 {
		return fmt.Errorf("field 'password' must contain at least 6 characters")
	}
	return nil
}
