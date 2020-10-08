package server

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"

	"github.com/cronui/cronui/internal/entity"
)

func (s *Server) handleInstallGet(c *fiber.Ctx) error {
	installed, err := s.isInstalled()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"installed": installed,
	})
}

func (s *Server) handleInstallPost(c *fiber.Ctx) error {
	installed, err := s.isInstalled()
	if err != nil {
		return err
	}
	if installed {
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

	user := entity.User{
		Name:      r.Username,
		Pass:      string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.db.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id": user.ID,
	})
}

func (s *Server) isInstalled() (bool, error) {
	var count int64
	if err := s.db.Model(&entity.User{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
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
