package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)


var getMyName = func(c *fiber.Ctx) error {
    return c.SendString("Hello, World! I am MD Rafiqul Islam")
}

func main(){
    app:= fiber.New()

    //  Routes
    app.Get("/",getMyName).Name("Home Page")
    
    app.Get("/loop/:times",func(c *fiber.Ctx) error{
        str_to_int,_ := strconv.Atoi(c.Params("times")) 
        for i :=1; i<= str_to_int; i++{
            fmt.Println(i)
        }
        return c.SendString("Loop is running "+c.Params("times"))
    })



    a := app.Group("/api")
    a.Name("API")
    a.Get("/hello",getMyName).Name("Api Hello")



    app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))



    //  Start Server 
    fmt.Println("Server running at http://localhost:3000")
    data, _ := json.MarshalIndent(app.Stack(), "", "  ")
    fmt.Println(string(data))
    app.Listen(":3000")
}