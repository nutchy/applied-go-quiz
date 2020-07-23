package merchant

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	sv *Service
}

// NewEndpoint -
func NewEndpoint(sv *Service) *Endpoint {
	return &Endpoint{sv: sv}
}

// GetMerchantInfo -
func (e Endpoint) GetMerchantInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"merchant": "foo",
	})
}
