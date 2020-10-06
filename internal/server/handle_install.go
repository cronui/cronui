package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/cronui/cronui/internal/entity"
)

func (s *Server) handleInstallGet(c *fiber.Ctx) error {
	installed, err := s.isInstalled()
	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"installed": installed,
	})

	return nil
}

func (s *Server) handleInstallPost(c *fiber.Ctx) error {
	installed, err := s.isInstalled()
	if err != nil {
		return err
	}
	if installed {
		return errors.New("user is already defined")
	}



	return errors.New("TODO implement")
}

func (s *Server) isInstalled() (bool, error) {
	var count int64
	if err := s.db.Model(&entity.User{}).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
