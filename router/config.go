package router

import (
	"github.com/TruongTrongThanh/ImageServer/router/routes"
)

// Init create router with routes in routes folder
func Init(root string) *Router {
	rtr := NewRouter(root)

	// Index routes
	rtr.RegisterRoute("/", "GET", routes.Index)
	rtr.RegisterRoute("/debug", "GET", routes.Debug)

	// Image routes
	rtr.RegisterRoute("/image", "GET", routes.GetImages)
	rtr.RegisterRoute("/image", "POST", routes.UploadImages)
	rtr.RegisterRoute("/image", "DELETE", routes.DeleteImage)

	return rtr
}
