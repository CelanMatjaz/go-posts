package middleware

import (
	"net/http"

	utils "github.com/CelanMatjaz/go-posts/internal/utils"
)

func EnsureAuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    ctx := &utils.CustomContext{r, w}

		if !ctx.IsAuthenticated() {
            ctx.Redirect("/login")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
