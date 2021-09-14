package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Forget(c*gin.Context)  {
	param, _ := c.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(param, &m)


}