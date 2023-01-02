package service
import(

	"fmt"
	"hexagonalapi/database"
)

type Service interface {
	Post()
}

type Middleware struct {
	xmen database. Database
}

func Instance(x database.Database) *Middleware {
	return &Middleware{xmen: x}

}

func (p *Middleware) Post() {
	p.xmen = &database.Model{}
	p.xmen.Post()
	fmt.Println("middleware post")


}