package handler

import (
	"chatplus/core"
	"chatplus/store/model"
	"chatplus/utils"
	"chatplus/utils/resp"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConfigHandler struct {
	BaseHandler
}

func NewConfigHandler(app *core.AppServer, db *gorm.DB) *ConfigHandler {
	return &ConfigHandler{BaseHandler: BaseHandler{App: app, DB: db}}
}

// Get 获取指定的系统配置
func (h *ConfigHandler) Get(c *gin.Context) {
	key := c.Query("key")
	fmt.Println("这是获取系统配置的key ", key)
	var config model.Config
	res := h.DB.Where("marker", key).Find(&config)
	if res.Error != nil {
		fmt.Println("这是获取配置的error ", res.Error)
		resp.ERROR(c, res.Error.Error())
		return
	}

	var value map[string]interface{}
	err := utils.JsonDecode(config.Config, &value)
	if err != nil {
		fmt.Println("这是转换json 报错 ", err)
		fmt.Println(value)
		resp.ERROR(c, err.Error())
		return
	}

	resp.SUCCESS(c, value)
}
