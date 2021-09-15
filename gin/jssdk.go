package gin

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nilorg/sdk/random"
	wechat "github.com/zj77086536/go-wechat"
)

// NewHandleJssdkConfig 获取jssdk配置
func NewHandleJssdkConfig(appID string, client wechat.Clienter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Query("url")
		if uri == "" {
			ctx.JSON(400, "url不能为空")
			return
		}
		uriDecode, err := base64.StdEncoding.DecodeString(uri)
		if err != nil {
			ctx.JSON(400, "base64 decode error")
			return
		}
		timestamp := time.Now().Unix()
		noncestr := random.AZaz09(16)
		uriLayout := "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
		signatureParams := fmt.Sprintf(uriLayout, client.GetJsAPITicket(), noncestr, timestamp, string(uriDecode))
		h := sha1.New()
		io.WriteString(h, signatureParams)
		ctx.JSON(200, map[string]interface{}{
			"app_id":    appID,
			"noncestr":  noncestr,
			"timestamp": timestamp,
			"signature": fmt.Sprintf("%x", h.Sum(nil)),
		})
	}
}
