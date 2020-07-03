package WuflyGo

import (
	"encoding/json"
	"net/http"
	"sync"
)

type HttpContext struct {
	Req        *http.Request
	Resp       *responseWriter
	store      map[string]interface{}
	storeMutex *sync.RWMutex
}

func (ctx *HttpContext) JSON(data interface{}) {
	ctx.Req.Header.Set("Context-type", "application/json")
	jsons, _ := json.Marshal(data)
	_, _ = ctx.Resp.Write(jsons)
}

// NewContext 创建context
func NewContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	ctx := &HttpContext{}
	ctx.Init(w, r)
	return ctx
}

// Init 初始化context
func (ctx *HttpContext) Init(w http.ResponseWriter, r *http.Request) {
	ctx.storeMutex = new(sync.RWMutex)
	ctx.Resp = &responseWriter{w, 0}
	ctx.Req = r
	ctx.storeMutex.Lock()
	ctx.store = nil
	ctx.storeMutex.Unlock()
}

// Set 设置context的值
func (ctx HttpContext) Set(key string, val interface{}) {
	ctx.storeMutex.Lock()
	if ctx.store == nil {
		ctx.store = make(map[string]interface{})
	}
	ctx.store[key] = val
	ctx.storeMutex.Unlock()
}

// Get 获取context的值
func (ctx HttpContext) Get(key string) interface{} {
	ctx.storeMutex.RLock()
	val := ctx.store[key]
	ctx.storeMutex.RUnlock()
	return val
}
