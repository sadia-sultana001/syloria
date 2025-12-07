package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func NewMessage() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {

	n := next

	//for i := len(middlewares) - 1; i >= 0; i-- {
	//	mdW := middlewares[i]
	//	n = mdW(n)
	//}
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMidWare := range mngr.globalMiddleware {
		n = globalMidWare(n)
	}

	return n

}
