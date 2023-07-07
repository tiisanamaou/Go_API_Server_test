package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostResponse struct {
	UserID   string `json:"UseID"`
	UserRank int    `json:"UserRank"`
	UserName string `json:"UserName"`
}

func GetAPI(w http.ResponseWriter, r *http.Request) {
	//Validate request
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// CROSエラーが出ないようにする設定
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Add Response
	postResponse := PostResponse{
		UserID:   "0002",
		UserRank: 15,
		UserName: "sana",
	}
	jsonData, err := json.Marshal(postResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

	fmt.Fprintln(w, string(jsonData))
}
