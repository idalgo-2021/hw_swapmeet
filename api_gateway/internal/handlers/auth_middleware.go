package handlers

import (
	auth "api_gateway/pkg/api/auth"

	"net/http"
)

func AuthMiddleware(authClient auth.AuthServiceClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken, err := extractBearerTokenFromRequest(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			_, err = authClient.ValidateToken(r.Context(), &auth.ValidateTokenRequest{AccessToken: accessToken})
			if err != nil {
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}

			// ctx := context.WithValue(r.Context(), "tokenClaims", resp)
			// next.ServeHTTP(w, r.WithContext(ctx))
			next.ServeHTTP(w, r)
		})
	}
}
