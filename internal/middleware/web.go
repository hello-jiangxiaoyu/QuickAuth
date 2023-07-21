package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
)

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func StaticWebFile() gin.HandlerFunc {
	root := "./web/out"
	fs := &localFileSystem{
		FileSystem: gin.Dir(root, false),
		root:       root,
		indexes:    false,
	}
	fileserver := http.FileServer(fs)

	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		if strings.HasPrefix(urlPath, "/api/quick") {
			return
		}

		urlPath = amendPath(urlPath)
		if p := strings.TrimPrefix(urlPath, "/"); len(p) < len(urlPath) {
			name := path.Join(fs.root, p)
			stats, err := os.Stat(name)
			if err != nil {
				return
			}
			if stats.IsDir() {
				if !fs.indexes {
					index := path.Join(name, "index.html")
					if _, err = os.Stat(index); err != nil {
						return
					}
				}
			}
			req := c.Request.Clone(c)
			req.URL.Path = urlPath
			fileserver.ServeHTTP(c.Writer, req)
			c.Abort()
		}
	}
}

func amendPath(p string) string {
	if strings.HasPrefix(p, "/applications/") && len(p) > len("/applications/") {
		for i := len("/applications/"); i < len(p); i++ {
			if p[i] == '/' {
				return "/applications/[appId]" + p[i:]
			}
		}
		return "applications/[appId]/"
	}
	return p
}
