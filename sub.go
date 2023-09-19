package main

import (
	"context"
	"crypto/sha256"
	_ "crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func sub() {
	loggingSettings("test.log")
	fmt.Println("OK")
	jwtGenerater()
	redisConnect()
	for i := 0; i < 1000; i++ {
		redisConnect()
	}
	//
	h := hashConvert("password", "passpass")
	fmt.Println(h)
}

var JwtString string

type UserData struct {
	ID       string
	Password string
}

// JWTの生成と検証
func jwtGenerater() {
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": "mao",
		//"password": "12345",
		//"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Header: map[string]interface {}{"alg":"HS256", "typ":"JWT"}
	fmt.Printf("Header: %#v\n", token.Header)
	// Claims: jwt.MapClaims{"exp":1634051243, "user_id":12345678}
	fmt.Printf("Claims: %#v\n", token.Claims)

	// トークンに署名を付与
	//tokenString, _ := token.SignedString([]byte("SECRET_KEY"))
	tokenString, _ := token.SignedString([]byte("prsk_key"))
	//
	fmt.Println("tokenString:", tokenString)
	JwtString = tokenString
	//
	fmt.Println("----------")
	//
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		//
		//if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		//}

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//return []byte("SECRET_KEY"), nil
		return []byte("prsk_key"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
		fmt.Printf("user_id: %v\n", string(claims["user_id"].(string)))
		//fmt.Printf("password: %v\n", string(claims["password"].(string)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
		fmt.Printf("alg: %v\n", token.Header["alg"])
	} else {
		fmt.Println("認証エラー")
		fmt.Println(err)
	}
	//
	fmt.Println("----------")
	// unix時刻を日時に変換し有効時間を確認
	unix := int64(claims["exp"].(int64))
	dtFromUnix := time.Unix(unix, 0)
	//
	//nowtime := time.Now()
	nowtime := time.Now().Add(time.Hour * 2)
	// 時刻比較、現在時刻を超えているか、超えていたらfalse
	diff := dtFromUnix.After(nowtime)
	//
	fmt.Println(nowtime)
	fmt.Println(dtFromUnix)
	fmt.Println(diff)
	//
	fmt.Println("----------")
	nowunix := time.Now().Unix()
	//nowunix := time.Now().Add(time.Hour * 2).Unix()
	fmt.Println("制限時刻:", unix)
	fmt.Println("現在時刻:", nowunix)
	if unix < nowunix {
		fmt.Println("制限時刻を超えています")
	} else {
		fmt.Println("制限時刻を超えていません")
	}
}

// Redisとの接続
func redisConnect() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})
	//
	u := UserData{
		ID:       "123456",
		Password: "pass",
	}
	jsonData, err := json.Marshal(&u)
	if err != nil {
		fmt.Println("変換エラー")
	}
	//rdb.Set(ctx, "mykey1", "hogehoge", 0)
	//rdb.Set(ctx, "session-uuid1", JwtString, 0)
	//rdb.Set(ctx, uuid, JwtString, 0)
	//
	uuid := UuidGenerate()
	//
	//err = rdb.Set(ctx, "session-uuid2", jsonData, 0).Err()
	err = rdb.Set(ctx, uuid, jsonData, 0).Err()
	if err != nil {
		fmt.Println("登録エラー")
	}
	//
	//ret, err := rdb.Get(ctx, "session-uuid1").Result()
	ret, err := rdb.Get(ctx, uuid).Result()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Result: ", ret)
	//log.Println("Result: ", ret)

	// 登録されているKeyをすべて取得、表示する
	cmd := rdb.Keys(ctx, "*")
	res, err := cmd.Result()
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Printf("%v\n", res)
		fmt.Println(len(res))
		// logファイルに出力する用
		log.Println("Redisデータ登録数:", len(res))
		//for i := 0; i < len(res); i++ {
		//	fmt.Println(res[i])
		//}
	}
}

// UUID を作成する
func UuidGenerate() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	uu := u.String()
	//fmt.Println(uu)
	return uu
}

// ハッシュ化+ソルト+ストレッチング
func hashConvert(password string, solt string) string {
	str := password + solt
	//s := []byte(password)
	//sha512 := sha512.Sum512(s)
	//sha512 := sha512.Sum512([]byte(password))
	sha := sha256.Sum256([]byte(str))
	//b := sha512[:]
	//h := hex.EncodeToString(b)
	str2 := hex.EncodeToString(sha[:])
	// ソルト+ストレッチング
	str2 = str2 + solt
	sha2 := sha256.Sum256([]byte(str2))
	h := hex.EncodeToString(sha2[:])
	//fmt.Println(h)
	return h
}

// logファイルに記録
func loggingSettings(fileName string) {
	logfile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// logをマイクロ秒まで表示させる
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(multiLogFile)
}
