package controllers
import(
	"net/http"

	"../common"
	"../models"
	"github.com/astaxie/beego/logs"
)
func fovourControll(w http.ResponseWriter, r *http.Request) {
	body, err := GetBodyData(r)
	if err != nil {
		logs.Error(err)
		ErrorResp(w, r, err.Error(), http.StatusInternalServerError) //500
		return
	}
	logs.Info(body)
	data, ok := body["data"].(map[string]interface{})
	if !ok {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest) //400
		return
	}

	action, ok1 := data["action"].(int) //0表示取消收藏，1表示添加收藏，2查看收藏状态
	uid, ok2 := data["uid"].(int) //用户id
	gid, ok3 := data["gid"].(int) //货品id
	if !ok1 || !ok2 || !ok3 {
		logs.Warn("读取表单参数错误")
		ErrorResp(w, r, "读取表单参数错误", http.StatusBadRequest)
		return
	}
	
	var respData interface{}
	var msg string
	if action == 0 {
		err = models.RemoveFavour(gid, uid)
		if err != nil {
			msg = "取消收藏失败"
			logs.Error("取消收藏失败，err:", err)
		}else {
			msg = "取消收藏成功"
			respData = struct{ success bool }{ true }
		}
        
	}else if action == 1 {
		err = models.AddFavour(gid, uid)
		if err != nil {
			msg = "添加收藏失败"
			logs.Error("添加收藏失败，err:", err)
		}else {
			msg = "添加收藏成功"
			respData = struct{ success bool }{ true }
		}	
	}else if action == 2 {
		status, err := models.SeeFavourStatus(gid, uid)
		if err != nil {
			msg = "查看收藏状态失败"
			logs.Error("查看收藏状态失败，err:", err)
		}else {
			msg = "查看收藏状态成功"
			respData = struct{ status bool }{ status }//true表示已收藏
		}	
	}else {
		if err := NaturalResp(w, r, "", "请求信息有误", 0); err != nil {
			logs.Error(err)
			return
		}
	}
	if err != nil {
		ErrorResp(w, r, msg, http.StatusInternalServerError)
		return
	}
	if err := NaturalResp(w, r, respData, msg, 0); err != nil {
		logs.Error(err)
		return
	}
}