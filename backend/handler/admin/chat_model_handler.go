package admin

import (
	"chatplus/core"
	"chatplus/core/types"
	"chatplus/handler"
	"chatplus/store/model"
	"chatplus/store/vo"
	"chatplus/utils"
	"chatplus/utils/resp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatModelHandler struct {
	handler.BaseHandler
}

func NewChatModelHandler(app *core.AppServer, db *gorm.DB) *ChatModelHandler {
	return &ChatModelHandler{BaseHandler: handler.BaseHandler{App: app, DB: db}}
}

func (h *ChatModelHandler) Save(c *gin.Context) {
	var data struct {
		Id          uint    `json:"id"`
		Name        string  `json:"name"`
		Value       string  `json:"value"`
		Enabled     bool    `json:"enabled"`
		SortNum     int     `json:"sort_num"`
		Open        bool    `json:"open"`
		Platform    string  `json:"platform"`
		Power       int     `json:"power"`
		MaxTokens   int     `json:"max_tokens"`  // 最大响应长度
		MaxContext  int     `json:"max_context"` // 最大上下文长度
		Temperature float32 `json:"temperature"` // 模型温度
		KeyId       int     `json:"key_id,omitempty"`
		CreatedAt   int64   `json:"created_at"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		resp.ERROR(c, types.InvalidArgs)
		return
	}

	item := model.ChatModel{
		Platform:    data.Platform,
		Name:        data.Name,
		Value:       data.Value,
		Enabled:     data.Enabled,
		SortNum:     data.SortNum,
		Open:        data.Open,
		MaxTokens:   data.MaxTokens,
		MaxContext:  data.MaxContext,
		Temperature: data.Temperature,
		KeyId:       data.KeyId,
		Power:       data.Power}
	var res *gorm.DB
	if data.Id > 0 {
		item.Id = data.Id
		res = h.DB.Select("*").Omit("created_at").Updates(&item)
	} else {
		res = h.DB.Create(&item)
	}
	if res.Error != nil {
		resp.ERROR(c, "更新数据库失败！")
		return
	}

	var itemVo vo.ChatModel
	err := utils.CopyObject(item, &itemVo)
	if err != nil {
		resp.ERROR(c, "数据拷贝失败！")
		return
	}
	itemVo.Id = item.Id
	itemVo.CreatedAt = item.CreatedAt.Unix()
	resp.SUCCESS(c, itemVo)
}

// List 模型列表
func (h *ChatModelHandler) List(c *gin.Context) {
	session := h.DB.Session(&gorm.Session{})
	enable := h.GetBool(c, "enable")
	if enable {
		session = session.Where("enabled", enable)
	}
	var items []model.ChatModel
	var cms = make([]vo.ChatModel, 0)
	res := session.Order("sort_num ASC").Find(&items)
	if res.Error != nil {
		resp.SUCCESS(c, cms)
		return
	}

	// initialize key name
	keyIds := make([]int, 0)
	for _, v := range items {
		keyIds = append(keyIds, v.KeyId)
	}
	var keys []model.ApiKey
	keyMap := make(map[uint]string)
	h.DB.Where("id IN ?", keyIds).Find(&keys)
	for _, v := range keys {
		keyMap[v.Id] = v.Name
	}
	for _, item := range items {
		var cm vo.ChatModel
		err := utils.CopyObject(item, &cm)
		if err == nil {
			cm.Id = item.Id
			cm.CreatedAt = item.CreatedAt.Unix()
			cm.UpdatedAt = item.UpdatedAt.Unix()
			cm.KeyName = keyMap[uint(item.KeyId)]
			cms = append(cms, cm)
		} else {
			logger.Error(err)
		}
	}
	resp.SUCCESS(c, cms)
}

func (h *ChatModelHandler) Set(c *gin.Context) {
	var data struct {
		Id    uint        `json:"id"`
		Filed string      `json:"filed"`
		Value interface{} `json:"value"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		resp.ERROR(c, types.InvalidArgs)
		return
	}

	res := h.DB.Model(&model.ChatModel{}).Where("id = ?", data.Id).Update(data.Filed, data.Value)
	if res.Error != nil {
		resp.ERROR(c, "更新数据库失败！")
		return
	}
	resp.SUCCESS(c)
}

func (h *ChatModelHandler) Sort(c *gin.Context) {
	var data struct {
		Ids   []uint `json:"ids"`
		Sorts []int  `json:"sorts"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		resp.ERROR(c, types.InvalidArgs)
		return
	}

	for index, id := range data.Ids {
		res := h.DB.Model(&model.ChatModel{}).Where("id = ?", id).Update("sort_num", data.Sorts[index])
		if res.Error != nil {
			resp.ERROR(c, "更新数据库失败！")
			return
		}
	}

	resp.SUCCESS(c)
}

func (h *ChatModelHandler) Remove(c *gin.Context) {
	id := h.GetInt(c, "id", 0)
	if id <= 0 {
		resp.ERROR(c, types.InvalidArgs)
		return
	}

	res := h.DB.Where("id = ?", id).Delete(&model.ChatModel{})
	if res.Error != nil {
		resp.ERROR(c, "更新数据库失败！")
		return
	}
	resp.SUCCESS(c)
}
