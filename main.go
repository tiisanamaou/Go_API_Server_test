package main

import (
	"encoding/json"
	"golang-server/router"

	//"log"
	"net/http"
	"strconv"
)

type (
	SampleHandler struct {
		Data string
	}
	SampleResponse struct {
		Status     string `json:"status"`
		Message    string `json:"message"`
		ReturnCode string `json:"returnCode"`
		UserData   string `json:"userData"`
	}
)

func main() {
	// Add Method GET
	http.HandleFunc("/get", router.GetAPI)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)

	// httpHandlerの準備
	//mux := &http.ServeMux{}

	// httpHandlerの設定。第1引数に設定したURLへ接続すると第2引数のHandler処理が実行されるようになる
	//mux.Handle("/api", NewSampleHandler())

	// httpサーバー起動処理。引数にはポート番号とhttpHandlerを設定する
	//if err := http.ListenAndServe(":8080", mux); err != nil {
	//	log.Fatal(err)
	//}
}

// SampleHandlerの構造体にinterfaceのhttp.Handlerを設定して返す関数
// interfaceのhttp.HandlerにはServeHTTP関数が含まれており、後の処理ListenAndServe関数から呼び出される
func NewSampleHandler() http.Handler {
	return &SampleHandler{"テストテキスト"}
}

// http.Handlerのinterfaceで定義されているServeHTTP関数を作成する。
// ServeHTTP関数はListenAndServe関数内で呼び出される
func (h *SampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// リターンコードの設定
	returnCode := 200

	// httpResponseの内容を設定
	res := &SampleResponse{
		Status:     "OK",
		Message:    h.Data,
		ReturnCode: strconv.Itoa(returnCode),
		UserData:   "ユーザーネーム",
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// CROSエラーが出ないようにする設定
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// ステータスコードを設定、200
	w.WriteHeader(http.StatusOK)

	// httpResponseの内容を書き込む
	buf, _ := json.Marshal(res)
	_, _ = w.Write(buf)
}
