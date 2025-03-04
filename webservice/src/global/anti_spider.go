package global

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type _CheckAntiSpiderFunc func(context *gin.Context) bool

func checkUserAgent(context *gin.Context) bool {
	ua := context.GetHeader(StrUserAgent)
	if len(ua) == 0 {
		return false
	}
	forbiddenUAs := []string{"Googlebot", "Baiduspider", "MJ12bot", "AhrefsBot", "YandexBot", "Python"}
	for _, bot := range forbiddenUAs {
		if strings.Contains(strings.ToLower(ua), strings.ToLower(bot)) {
			return false
		}
	}
	return true
}

func checkReferer(context *gin.Context) bool {
	referer := context.GetHeader(StrUserAgent)
	if len(referer) > 0 && !strings.Contains(referer, StrCurrentDomain) {
		return false
	}
	return true
}

func AntiSpiderMiddleware() gin.HandlerFunc {
	var checkAntiSpiderFuncList = []_CheckAntiSpiderFunc{checkUserAgent, checkReferer}
	return func(context *gin.Context) {
		for _, checkFunc := range checkAntiSpiderFuncList {
			if !checkFunc(context) {
				context.AbortWithStatus(http.StatusForbidden)
				return
			}
		}
	}
}
