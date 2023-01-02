package database 

import(
"fmt"
)
type Database interface {
	post()
}

type Model struct {
}

func Getdatabase() *Model {
	return &Model{}

}

func (p *Model) Post() {
	fmt.Println("model post")

}