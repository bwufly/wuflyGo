package WuflyGo

import (
	"encoding/json"
	"net/http"
	"sync"
)

type HttpContext struct {
	Req *http.Request
	Resp *responseWriter
	store map[string]interface{}
	storeMutex *sync.RWMutex
}

func (ctx *HttpContext) JSON(data interface{})  {
	ctx.Req.Header.Set("Context-type","application/json")
	jsons,_ := json.Marshal(data)
	_,_ = ctx.Resp.Write(jsons)
}

// NewContext 创建context
func NewContext(w http.ResponseWriter,r *http.Request)*HttpContext  {
	ctx := &HttpContext{}
	ctx.Init(w,r)
	return ctx
}

// Init 初始化context
func (ctx *HttpContext) Init(w http.ResponseWriter,r *http.Request)  {
	ctx.storeMutex = new(sync.RWMutex)
	ctx.Resp = &responseWriter{w,0}
	ctx.Req = r
	ctx.storeMutex.Lock()
	ctx.store = nil
	ctx.storeMutex.Unlock()
}

