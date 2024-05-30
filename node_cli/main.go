package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"node_cli/conf"
	"node_cli/position"
	exit "node_cli/utils"
	"os"
)

type ItemStyle struct {
	Color string `json:"color"`
}

type Location struct {
	Name      string    `json:"name"`
	Value     []float64 `json:"value"`
	ItemStyle struct {
		Color string `json:"color"`
	} `json:"itemStyle"`
	Price int          `json:"price"`
	Api   conf.ApiConf `json:"api"`
}

func main() {
	name := ""
	price := 0

	fmt.Println("节点使用规则")
	fmt.Println("使用 -price 指定节点每千字的费用")
	fmt.Println("使用 -name 指定创建节点的用户名")
	flag.StringVar(&name, "name", "", "set the name for the client node")
	flag.IntVar(&price, "price", 0, "set the price for api")
	flag.Parse()
	if name == "" || price == 0 {
		fmt.Println("请确实指定了价格与用户名")
		res := errors.New("错误的输入了价格或者用户名")
		panic(res)
	}

	loc := Location{}
	loc.Value = position.GetPos()
	loc.ItemStyle.Color = "#A6283F"
	loc.Name = name
	loc.Price = price
	loc.Api = *conf.InitConf("./conf.yaml")
	fmt.Println(loc)
	reqBytes, err := json.Marshal(loc)
	if err != nil {
		panic(err)
	}
	reqBody := bytes.NewBuffer(reqBytes)
	req, err := http.NewRequest("POST", "http://localhost:8080", reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Status", "alive")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	c := make(chan os.Signal, 1)
	exit.WaitElegantExit(c)
	req, err = http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(reqBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Status", "kill")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
