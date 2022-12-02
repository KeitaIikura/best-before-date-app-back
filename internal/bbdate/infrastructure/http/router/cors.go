package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type (
	CORSRouter interface {
		Use(middleware ...gin.HandlerFunc)
		GET(path string, handlers ...gin.HandlerFunc)
		POST(path string, handlers ...gin.HandlerFunc)
		PUT(path string, handlers ...gin.HandlerFunc)
		DELETE(path string, handlers ...gin.HandlerFunc)
	}
	CORSRouterImpl struct {
		g           *gin.Engine
		corsOrigins string
		routes      map[string]string
	}
)

// 管理画面と加盟店用画面で許可したい範囲が異なるたんめ originをstringで受け付ける
func NewCorsRouterImpl(g *gin.Engine, allowOrigin string) CORSRouter {
	g.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == allowOrigin
		},
		MaxAge: 12 * time.Hour,
	}))
	return &CORSRouterImpl{
		g:           g,
		corsOrigins: allowOrigin,
		routes:      map[string]string{},
	}
}

func (r *CORSRouterImpl) Use(middleware ...gin.HandlerFunc) {
	r.g.Use(middleware...)
}

func (r *CORSRouterImpl) GET(path string, handlers ...gin.HandlerFunc) {
	r.optionsOnce(path, r.optionHandler)
	r.g.GET(path, handlers...)
}

func (r *CORSRouterImpl) POST(path string, handlers ...gin.HandlerFunc) {
	r.optionsOnce(path, r.optionHandler)
	r.g.POST(path, handlers...)
}

func (r *CORSRouterImpl) PUT(path string, handlers ...gin.HandlerFunc) {
	r.optionsOnce(path, r.optionHandler)
	r.g.PUT(path, handlers...)
}

func (r *CORSRouterImpl) DELETE(path string, handlers ...gin.HandlerFunc) {
	r.optionsOnce(path, r.optionHandler)
	r.g.DELETE(path, handlers...)
}

func (r *CORSRouterImpl) optionHandler(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if r.corsOrigins == origin || r.corsOrigins == "*" {
		c.AbortWithStatus(204)
		return
	}
	c.AbortWithStatus(403)
}

func (r *CORSRouterImpl) optionsOnce(path string, handlers ...gin.HandlerFunc) {
	if _, exist := r.routes[path]; exist {
		return
	}
	r.routes[path] = path
	r.g.OPTIONS(path, r.optionHandler)
}
