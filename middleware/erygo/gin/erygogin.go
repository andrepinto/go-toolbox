package erygogin

import (
	"fmt"

	"github.com/andrepinto/erygo"
	"github.com/gin-gonic/gin"
)

// Gonic -- aborts gin HTTP request with StatusHTTP
// and provides json representation of error
func Gonic(err *erygo.Err, ctx *gin.Context) {
	//	ctx.Error(err)
	ctx.AbortWithStatusJSON(err.StatusHTTP, err)
}

// Recovery -- gin middleware to catch panics and wrap it to cherry error.
// If panic caught it aborts HTTP request with defaultErr.
func Recovery(defaultErr erygo.ErrConstruct, logger erygo.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			var errToReturn *erygo.Err
			if r := recover(); r != nil {
				if erygoErr, ok := r.(*erygo.Err); ok {
					errToReturn = erygoErr
				} else {
					errToReturn = defaultErr().Log(fmt.Errorf("%v", r), logger)
				}
				Gonic(errToReturn, ctx)
			}
		}()

		ctx.Next()
	}
}
