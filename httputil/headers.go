package httputil

import (
	"context"
	"net/textproto"
	"strings"

	"github.com/andrepinto/erygo"
	"github.com/andrepinto/go-toolbox/middleware/erygo/gin"
	"github.com/gin-gonic/gin"
)

type contextKey int

const (
	ClientContextKey contextKey = iota
	ProductContextKey
	RequestIDContextKey

	AcceptLanguageContextKey
)

//Extrernal headers
const (
	UserClientHeader    = "User-Client"
	AuthorizationHeader = "User-Token"
)

//Internal headers
const (
	RequestIDXHeader = "X-Request-ID"
)

var hdrToKey = map[string]interface{}{
	textproto.CanonicalMIMEHeaderKey(RequestIDXHeader): RequestIDContextKey,
}

//RequireHeaders ...
func RequireHeaders(errToReturn erygo.ErrConstruct, headers ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var notFoundHeaders []string
		for _, v := range headers {
			if ctx.GetHeader(textproto.CanonicalMIMEHeaderKey(v)) == "" {
				notFoundHeaders = append(notFoundHeaders, v)
			}
		}
		if len(notFoundHeaders) > 0 {
			err := errToReturn()
			for _, notFoundHeader := range notFoundHeaders {
				err.AddDetailF("required header %s was not provided", notFoundHeader)
			}
			erygogin.Gonic(err, ctx)
		}
	}
}

func PrepareContext(ctx *gin.Context) {
	for hn, ck := range hdrToKey {
		if hv := ctx.GetHeader(hn); hv != "" {
			rctx := context.WithValue(ctx.Request.Context(), ck, hv)
			ctx.Request = ctx.Request.WithContext(rctx)
		}
	}

	acceptLanguages := ctx.GetHeader("Accept-Language")
	acceptLanguagesToContext := make([]string, 0)
	for _, language := range strings.Split(acceptLanguages, ",") {
		language = strings.Split(strings.TrimSpace(language), ";")[0] // drop quality values
		acceptLanguagesToContext = append(acceptLanguagesToContext, language)
	}
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), AcceptLanguageContextKey, acceptLanguagesToContext))
}
