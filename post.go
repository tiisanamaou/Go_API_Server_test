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
	Password string `json:"Password"`
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
		Password: "pass",
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
	fmt.Println(string(body))

	// Responseにデータを格納
	//buf, _ := json.Marshal(body)
	//_, _ = w.Write(buf)
	//_, _ = w.Write(jsonData)
	contentLen, err := w.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Content-length
	fmt.Println(contentLen)

	var b PostResponse
	if err := json.Unmarshal(body, &b); err != nil {
		return
	}

	//
	fmt.Println(b.Password)
	if b.Password != "Password" {
		fmt.Println("Password NG")
		return
	}
	fmt.Println("Password OK")
}
