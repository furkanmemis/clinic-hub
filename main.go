package main

import (
	"clinic-hub/handler"
	middleware "clinic-hub/middlewares"
	"clinic-hub/services"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Multi Tenant w/Go")

	services.TenantInitilization()
	services.RoleInitilization("fuzei")
	services.AdminInitilization()

	fmt.Println("Authentication routes:")
	for route, handleFunc := range handler.AuthenticationHandlerMap {
		fmt.Printf("%s created.\n", route)
		http.HandleFunc(route, handleFunc)
	}

	fmt.Println("Tenant routes:")
	for route, handleFunc := range handler.TenantHandlerMap {
		fmt.Printf("%s created.\n", route)
		http.HandleFunc(route, handleFunc)
	}
	fmt.Println("User routes:")
	for route, handleFunc := range handler.UserHandlerMap {
		wrapped := middleware.JWTMiddleware(http.HandlerFunc(handleFunc))
		http.Handle(route, wrapped)
		fmt.Printf("%s created.\n", route)
	}

	fmt.Println("Server running with 8080")
	http.ListenAndServe(":8080", nil)
}
