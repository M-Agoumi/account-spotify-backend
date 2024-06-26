package account

import (
	"fmt"
	"github.com/M-Agoumi/account-spotify-backend/middleware"
	"github.com/M-Agoumi/account-spotify-backend/service/jwtService"
	"net/http"
)

type Account struct {
}

func (h *Account) User(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*jwtService.UserClaim)
	if !ok {
		http.Error(w, "No user in context", http.StatusInternalServerError)
		return
	}

	// Use the claims for your logic.
	_, err := fmt.Fprintf(w, "Hello, %s!", claims.Email)
	if err != nil {
		return
	}
}
