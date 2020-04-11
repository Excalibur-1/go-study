package main

import (
	"fmt"
	"net/http"
)

//在golang的web开发中，一个handler响应一个http请求：
type Handler interface {
	ServerHttp(http.ResponseWriter, *http.Request)
}

//handler可以有多种实现方式
//一：HandlerFunc，HandlerFunc是对用户定义的处理函数的包装，
// 实现了ServerHttp方法，在方法内执行HandlerFunc对象
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f HandlerFunc) ServerHttp(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

//二：ServerMux
//ServerMux是管理http请求pattern和handler的，实现了ServerHttp方法，
// 在其内根据请求匹配HandlerFunc，并执行其ServerHttp方法
/*type ServerMux struct {
	mu sync.RWMutex
	m map[string]muxEntry
	es []muxEntry
	hosts bool
}

func (mux *ServerMux)ServerHttp(w http.ResponseWriter, r *http.Request)  {
	h ,_:=mux.Handler(r)
	h.ServerHttp(w,r)//调用的是HandlerFunc的ServerHttp
}
*/

//三：serverHandler
//serverHandler是对Server的封装，实现了ServerHttp方法，并在其内执行ServerMux的ServerHttp方法
/*type serverHandler struct {
	srv *http.Server
}

func (sh serverHandler)ServerHttp(w http.ResponseWriter, r *http.Request)  {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServerMux
	}
	...
	handler.ServeHTTP(w,r) // 调用的是ServerMux的ServeHTTP
}*/
type User struct {
	Name string
	Age int
}
type GenderType int32

const (
	_GenderType GenderType = 0
	Unknown     GenderType = 1
	Male        GenderType = 2
	Female      GenderType = 3
)
func main() {

	var user User
	fmt.Println(&user)
	var gender = Male
	fmt.Println(uint8(gender))
}