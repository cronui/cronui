package config

import (
	"fmt"
)

type Server struct {
	Host string `default:"127.0.0.1"`
	Port int    `default:"4000"`
}

func (c *Server) GetListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
