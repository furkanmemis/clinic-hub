package main

import (
	"fmt"
	"clinic-hub/handler"
	"clinic-hub/services"
	"net/http"
)

func main() {

	fmt.Println("Multi Tenant w/Go")

	services.TenantInitilization()
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
		fmt.Printf("%s created.\n", route)
		http.HandleFunc(route, handleFunc)
	}

	fmt.Println("Server running with 8080")
	http.ListenAndServe(":8080", nil)
}
