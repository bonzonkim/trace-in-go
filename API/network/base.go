package network

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// func newRootSpan(name string, c *gin.Context) opentracing.Span {
// 	tracer := opentracing.GlobalTracer()
// 	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
// 	sendSpan := tracer.StartSpan(name, ext.RPCServerOption(spanCtx))

// defer sendSpan.Finish()
// return sendSpan
// func (s *Router) send(c *gin.Context) {
// 	//	newRootSpan("send_root_span", c)
// 	tracer := opentracing.GlobalTracer()
// 	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
// 	sendSpan := tracer.StartSpan("send_root_span", ext.RPCServerOption(spanCtx))

// 	defer sendSpan.Finish()

//		c.JSON(http.StatusOK, "Success")
//	}
func (s *Router) send(c *gin.Context) {
	fmt.Println("=================Send=================")

	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan("send_span", ext.RPCServerOption(spanCtx))

	defer sendSpan.Finish()

	fmt.Println("=================Send1=================")
	c.JSON(http.StatusOK, "Success Sample Span")
	fmt.Println("=================Send2=================")
}

func (s *Router) sendWithTag(c *gin.Context) {
}

func (s *Router) sendWithChild(c *gin.Context) {
}
