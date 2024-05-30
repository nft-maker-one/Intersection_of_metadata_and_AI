package handler

import (
	"chatplus/core"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Locs Locations = Locations{
	Loc: []Location{{"广东", []float64{113.266887, 23.133306}, ItemStyle{Color: "#00EEFF"}, 0, ApiConf{}}},
}

func init() {
	// 定义处理程序，接收客户端发来的数据
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// 读取客户端发来的数据
			location := Location{}
			bufData, err := io.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte("接收信息失败，请重新发送"))
				return
			}
			err = json.Unmarshal(bufData, &location)
			if err != nil {
				w.Write([]byte("解码失败"))
				return
			}
			if r.Header.Get("Status") == "kill" {
				for k, v := range Locs.Loc {
					if v.Name == location.Name {
						Locs.Loc = append(Locs.Loc[:k], Locs.Loc[k+1:]...)
						w.Write([]byte("删除成功"))
						break
					}
				}
			} else {
				Locs.Loc = append(Locs.Loc, location)
				w.Write([]byte("添加成功"))
			}

		})

		// 启动 HTTP 服务器，监听所有网络接口上的 8080 端口
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

}

func NewMap(app *core.AppServer, db *gorm.DB) *MapHandler {
	return &MapHandler{BaseHandler: BaseHandler{App: app, DB: db}}
}

type ItemStyle struct {
	Color string `json:"color"`
}

type Location struct {
	Name      string    `json:"name"`
	Value     []float64 `json:"value"`
	ItemStyle ItemStyle `json:"itemStyle"`
	Price     int       `json:"price"`
	Api       ApiConf   `json:"api"`
}

type ApiConf struct {
	Platform string `yaml:"platform" json:"platform"`
	ApiKey   string `yaml:"apiKey" json:"apiKey"`
	Version  string `yaml:"version" json:"version"`
	ProxyUrl string `yaml:"proxyUrl" json:"proxyUrl"`
}

type Locations struct {
	Loc    []Location `json:"loc"`
	Status bool       `json:"status"`
}

type MapHandler struct {
	BaseHandler
	Locations
}

func (l *MapHandler) Send(c *gin.Context) {
	c.JSON(http.StatusOK, Locs)
}
