package main

import(

	"fmt"
	"hexagonalapi/api"
	"hexagonalapi/service"
	"hexagonalapi/database"
)


func main() {
	ret := api.Createobject(&service.Middleware{})

		ret1 := Instance(&database.Model{})

		ret2 :=Getdatabase()  

	ret.Post()
	fmt.Println(ret,ret1)

}