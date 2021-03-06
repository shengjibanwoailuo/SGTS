package main

import (
	"net/http"

	"./controllers"
	"./env"
	"github.com/astaxie/beego/logs"
)

func main() {
	env.InitEnv()
	//res, _ := models.GetGoodsByType(0,1,"")
	//logs.Info(res)
	mux := http.NewServeMux()
	mux.HandleFunc("/go-helloWorld", controllers.GoHelloWorld)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/logout", controllers.Logout)
	mux.HandleFunc("/register", controllers.Register)
	// 获取用户信息
	mux.HandleFunc("/getUserInfo", controllers.GetUserInfo)
	// 收藏夹
	mux.HandleFunc("/getFolder", controllers.GetFolder)
	mux.HandleFunc("/getFavourGoods", controllers.GetFavourGoods)
	// 货品种类
	mux.HandleFunc("/getGoodsType", controllers.GetGoodsType)
	// 主页
	mux.HandleFunc("/getGoodsByType", controllers.GetGoodsByType)

	// 货品详情
	mux.HandleFunc("/getGoodInfo", controllers.GetGoodInfo)
	mux.HandleFunc("/getGoodImg", controllers.GetGoodImg)
	// 发布货品
	mux.HandleFunc("/uploadGood", controllers.UploadGood)
	// 货品管理：获取货品列表
	mux.HandleFunc("/getGoodsByCategory", controllers.GetGoodsByCategory)
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		logs.Error(err)
	}
}
