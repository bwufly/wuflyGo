package WuflyGo

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	// DefaultAddress is used if no other is specified
	DefaultAddress = ":1688"
)

type WuflyGo struct {
	
}

func Classic() *WuflyGo {
	return &WuflyGo{}
}

type M  = map[string]string

var reqFunMap = make(map[string]func(ctx *HttpContext))

// Map 相对路由
func (w *WuflyGo) Map(relativePath string,handler func(ctx *HttpContext))  {
	reqFunMap[relativePath] = handler
}

// 启动服务
func (w *WuflyGo) Run(addr ...string)  {
	l := log.New(os.Stdout,"[wufly]",0)
	finalAddr := detectAddress(addr...)
	l.Printf("listening on %s",finalAddr)
	l.Fatal(http.ListenAndServe(finalAddr,w))
}

// detectAddress 获取监听地址
func detectAddress(addr ...string) string  {
	if len(addr)>0 {
		return addr[0]
	}
	if port := os.Getenv("PORT");port !=""{
		return ":"+port
	}
	return DefaultAddress
}

func (wu *WuflyGo) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(r.URL.Path)
	ctx := NewContext(w,r)
	fun,ok := reqFunMap[r.URL.Path]
	if ok {
		fun(ctx)
		return
	}
}
