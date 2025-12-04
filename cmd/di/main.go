package main

import (
	"fmt"
)

// ログ出力
func LogOutput(message string) {
	fmt.Println(message)
}

// 簡易的なデータストア
type SimpleDataStore struct {
	userData map[string]string
}

// SimpleDataStore method
// userIDを受け取り、名前を返す
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	// カンマokイディオム
	name, ok := sds.userData[userID]
	return name, ok
}

// SimpleDataStoreのインスタンスを生成するファクトリ
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "花子",
			"2": "太郎",
			"3": "パット",
		},
	}
}

// DataStore Interface
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// Logger Interface
type Logger interface {
	Log(message string)
}

// 関数型LoggerAdapterを定義
type LoggerAdapter func(message string)

// LoggerAdapterに付随するメソッド
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// ビジネスロジック
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)

	if !ok {
		return "", fmt.Errorf("未登録のユーザーです")
	}
	return name + "さん、こんにちは", nil
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)

	name, err := logic.SayHello("4")
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println(name)
}
