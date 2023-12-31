package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseData struct {
	UserID   string `json:"UserID"`
	UserRank int    `json:"UserRank"`
	UserName string `json:"UserName"`
}

func GetAPI(w http.ResponseWriter, r *http.Request) {
	//Validate request
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// CROSエラーが出ないようにする設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// ステータスコードを設定
	w.WriteHeader(http.StatusOK) //200

	// Add Response
	responseData := ResponseData{
		UserID:   "0002",
		UserRank: 15,
		UserName: "_sana_",
	}
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	fmt.Fprintln(w, string(jsonData))
}
