package api

//hexagonal architecture uses two design patterns
//1. bridge pattern
//2. singleton
import (
	"fmt"
	"hexagonalapi/service"
)

type Api interface {
	Post()
}

type Restapi struct {
	heman service.Service //bridge pattern
}

func Createobject(h service.Service) *Restapi { //singleton
	return &Restapi{heman: h}

}

func (p *Restapi) Post() {
	fmt.Println("restapi post")
	//p.heman = &Middleware{}
	p.heman.Post()

}
