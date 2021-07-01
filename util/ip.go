package util

import (
	"strings"
	"net"
	"github.com/gin-gonic/gin"
)

/**
获取客户端ip: 包括公网和内网，如果只获取公网，还需做一个ip地址检测
 */
func ClientIp(c *gin.Context) string {
	xForwardedFor := c.Request.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])

	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(c.Request.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
