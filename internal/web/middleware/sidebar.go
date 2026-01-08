package middleware

import (
	"context"
	"net/http"

	"github.com/nelsonmarro/go-stripe/utils"
)

// SidebarStateMiddleware extracts the sidebar state from the "sidebar_state" cookie
// and injects it into the request context.
func SidebarStateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sidebar_state")
		
		// Logic:
		// Cookie "false" => Collapsed = true (The cookie name is sidebar_state, usually representing "is open")
		// Actually, let's re-read the prompt carefully:
		// "cookie named sidebar_state ... 'true' for expanded or 'false' for collapsed"
		// Code: "collapsed := cookie != nil && cookie.Value == "false""
		
		isCollapsed := false
		if err == nil && cookie.Value == "false" {
			isCollapsed = true
		}

		ctx := context.WithValue(r.Context(), utils.SidebarCollapsedKey, isCollapsed)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
