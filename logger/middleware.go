package logger

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

// XRI 唯一请求ID
const XRI = "X-Request-ID"

var snowflakeNode *snowflake.Node

func init() {
	// todo 水平扩展实例的时候node节点下面需要修改
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(fmt.Sprintf("init snowflake node fail, err:%s",err))
	}

	snowflakeNode = node
}

// RequestID 生成唯一请求ID中间件
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		xri := snowflakeNode.Generate()
		c.Set(XRI, xri.Int64())
		c.Header(XRI, strconv.Itoa(int(xri.Int64())))
		c.Next()
	}
}

// getRequestID 获取请求ID
func getRequestID(ctx *gin.Context) (requestID int64) {
	value := ctx.Value(XRI)
	if value != nil {
		v, ok := value.(int64)
		if !ok {
			Error(nil, "request Id assertion failed", "value", value)
		} else {
			requestID = v
		}
	}
	return
}
