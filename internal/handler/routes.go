package handler

import (
	"fmt"
	"net/http"

	"chayanon.t/internal/merchant"
	"github.com/labstack/echo/v4"
)

// Route -
type Route struct {
	Group    string
	Path     string
	Method   string
	Endpoint echo.HandlerFunc
}

// InitRoutes -
func InitRoutes(e *echo.Echo) error {
	// merchantRepo := merchant.NewRepo()
	merchantService := merchant.NewService()
	merchantEndpoint := merchant.NewEndpoint(merchantService)

	// Declare routes
	routes := []Route{
		{
			Group:    "merchant",
			Path:     "/",
			Method:   http.MethodGet,
			Endpoint: merchantEndpoint.GetMerchantInfo,
		},
		{
			Group:    "merchant",
			Path:     "/information",
			Method:   http.MethodGet,
			Endpoint: nil,
		},
		{
			Group:    "merchant",
			Path:     "/:id/products",
			Method:   http.MethodGet,
			Endpoint: nil,
		},
		{
			Group:    "merchant",
			Path:     "/:id/product",
			Method:   http.MethodPost,
			Endpoint: nil,
		},
	}

	// http connection
	for _, rt := range routes {
		fmt.Printf("%+v\n", rt)
		e.Group(rt.Group).Add(rt.Method, rt.Path, rt.Endpoint)
	}
	return nil
}
