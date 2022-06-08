package router

import "github.com/singcl/gin-taro-api/internal/pkg/core"

type resource struct {
	kiko core.Kiko
}

type Server struct {
	Kiko core.Kiko
}

func NewHTTPServer() (*Server, error) {
	r := new(resource)

	kiko, err := core.New()

	if err != nil {
		panic(err)
	}

	r.kiko = kiko

	s := new(Server)
	s.Kiko = kiko

	return s, nil
}
