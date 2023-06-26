package server

import (
	"net/http"
	"strings"

	"github.com/AnimeKaizoku/StickersApi/src/entryPoints/stickerHandlers"
	"github.com/gin-gonic/gin"
)

func LoadHandlers() {
	// revoke token handlers
	bindHandler(stickerHandlers.AddPackHandler, "addPack")
	bindHandler(stickerHandlers.GetPackInfoHandler, "getPack")
	bindHandler(stickerHandlers.SearchPackHandler, "searchPack")

	bindHandler(func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("addPack, getPack, searchPack"))
	}, "Docs")

	bindNoRoot()
}

func bindHandler(handler gin.HandlerFunc, paths ...string) {
	var currentPath string
	for _, path := range paths {
		// ServerEngine.GET(path, handler)
		// ServerEngine.POST(path, handler)

		currentPath = strings.ToLower(path)
		getHandlers[currentPath] = handler
		postHandlers[currentPath] = handler
	}
}

func bindNoRoot() {
	ServerEngine.NoRoute(noRootHandler)
}

func noRootHandler(c *gin.Context) {
	path := strings.TrimSpace(c.Request.URL.Path)
	path = strings.Trim(path, "/")
	path = strings.ToLower(path)
	switch c.Request.Method {
	case http.MethodGet:
		h := getHandlers[path]
		if h != nil {
			h(c)
			return
		}
		c.Redirect(http.StatusPermanentRedirect, DocsPath)
		return
	case http.MethodPost:
		h := postHandlers[path]
		if h != nil {
			h(c)
			return
		}
		c.Redirect(http.StatusPermanentRedirect, DocsPath)
		return
	default:
		goto redirect_req
	}

redirect_req: /* redirect requests to /docs */
	c.Redirect(http.StatusPermanentRedirect, DocsPath)
}
