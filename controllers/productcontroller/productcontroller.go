package productcontroller

import (
	"net/http"

	"github.com/xvbnm48/go-session-jwt/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":    1,
			"name":  "Product 1",
			"price": 100,
		},
		{
			"id":    2,
			"name":  "Product 2",
			"price": 200,
		},
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
