package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func newRootSpan(name string, c *gin.Context) opentracing.Span {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sendSpan := tracer.StartSpan(name, ext.RPCServerOption(spanCtx))

	defer sendSpan.Finish()
	return sendSpan
 }

func (s *Router) send(c *gin.Context) {
	newRootSpan("send_root_span", c)
	c.JSON(http.StatusOK, "Success")
}

func (s *Router) defaultHandler(c *gin.Context) {
	c.JSON(404, gin.H{
		"error": "This path is not yet implemented",
	})
}





















func (s *Router) sendWithTag(c *gin.Context) {
}

func (s *Router) sendWithChild(c *gin.Context) {
}
