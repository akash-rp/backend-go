package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func modifyResponse(res *http.Response) error {
	res.Header.Del("Access-Control-Allow-Origin")
	res.Header.Del("X-Kratos-Authenticated-Identity-Id")
	return nil
}

func getLoginAuthFlow(c *gin.Context) {
	target := "127.0.0.1:4433"
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.URL.Path = "/self-service/login/browser"
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func loginUser(c *gin.Context) {
	target := "127.0.0.1:4433"
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.URL.Path = "/self-service/login"
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func logoutUser(c *gin.Context) {
	target := "127.0.0.1:4433"
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.URL.Path = "/self-service/logout"
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func getLogoutAuthFlow(c *gin.Context) {
	target := "127.0.0.1:4433"
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.URL.Path = "/self-service/logout/browser"
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func getUserDetails(c *gin.Context) {
	target := "127.0.0.1:4433"
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.URL.Path = "/sessions/whoami"
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
	proxy.ServeHTTP(c.Writer, c.Request)
}
