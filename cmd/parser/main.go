package main

import (
	"aurora/generator"
	"aurora/parser"
	"fmt"
	"log"
)

func main() {
	dsl := `DefineControllerMethod:
        Name: ShowUserProfile
        Path: /user/{userId}
        Method: GET
        Action: return c.Render(200, r.JSON(GetUserProfile(userId)))`

	method, err := parser.ParseControllerMethod(dsl)
	if err != nil {
		log.Fatal(err)
	}

	goCode := generator.GenerateBuffaloCode(method)
	fmt.Println(goCode)
}
