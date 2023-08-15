package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostResponse struct {
	UserID   string `json:"UserID"`
	UserRank int    `json:"UserRank"`
	UserName string `json:"UserName"`
}

func post(w http.ResponseWriter, r *http.Request) {
	//Validate request
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// ステータスコードを設定、200
	w.WriteHeader(http.StatusOK)

	// Add Response
	postResponse := PostResponse{
		UserID:   "0001",
		UserRank: 10,
		UserName: "maomao_post",
	}
	//fmt.Println(postResponse)
	jsonData, err := json.Marshal(postResponse)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	len := r.ContentLength
	body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
	r.Body.Read(body)         // byte 配列にリクエストボディを読み込む
	fmt.Fprintln(w, string(body))
	fmt.Println(string(body))
	//fmt.Fprintln(w, string(jsonData))
}
