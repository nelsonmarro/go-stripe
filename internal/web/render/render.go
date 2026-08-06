// Package render provides functions to render and patch web page components.
package render

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/nelsonmarro/go-stripe/templates/layout"
	"github.com/nelsonmarro/go-stripe/utils"
	"github.com/starfederation/datastar-go/datastar"
)

// PatchSPA handles partial updates for Single Page Application navigation using Datastar.
// It patches the main content, the page header, and the sidebar navigation.
func PatchSPA(w http.ResponseWriter, r *http.Request, title string, content templ.Component) error {
	sse := datastar.NewSSE(w, r)

	// 1. Patch Main Content (wrapped to preserve ID)
	if err := sse.PatchElementTempl(layout.ContentWrapper(content), datastar.WithSelector("#content")); err != nil {
		return err
	}

	// 2. Patch Page Header (Breadcrumb/Title)
	if err := sse.PatchElementTempl(layout.PageHeader(title), datastar.WithSelector("#page-header")); err != nil {
		return err
	}

	// 3. Patch Sidebar Navigation (Active State)
	collapsed := false
	if val := r.Context().Value(utils.SidebarCollapsedKey); val != nil {
		if b, ok := val.(bool); ok {
			collapsed = b
		}
	}
	if err := sse.PatchElementTempl(layout.SidebarNavigation(title, collapsed), datastar.WithSelector("#sidebar-component")); err != nil {
		return err
	}

	return nil
}
