package std

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*-------------------------------------
 *配列の要素をスワップする
 *引数: a スライスのポインタ
       i 交換対象のインデックス
	   j 交換対象のインデックス
--------------------------------------*/
func Swap(a *[]int, i int, j int) {
	tmp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tmp
}

/*-------------------------------------
 *乱数生成用(汎用)
 *引数: r 乱数範囲
--------------------------------------*/
func GenRandomNum(r int) int64 {
	var res int64

	rand.Seed(time.Now().UnixNano())
	res = int64(rand.Intn(r))

	return res
}

/*-------------------------------------
 *エラーハンドリング
 *引数: err エラー
--------------------------------------*/
func Checkerr(err error) {
	if (err != nil) {
		panic(err)
	}
}

/*-------------------------------------
 *ファイル読み込み
 *引数: filename ファイル名
 *戻り値: 読み込んだファイルデータ(文字列)
--------------------------------------*/
func FileRead(filename string) string {
	file, err := os.Open(filename)
	Checkerr(err)

	data := make([]byte,1024)
	count, err := file.Read(data)
	Checkerr(err)

	return string(data[:count])
}

/*-------------------------------------
 *ファイル書き込み
 *引数: filename ファイル名
 *      data 書き込むデータ バイトデータ
--------------------------------------*/
func FileWrite(filename string, data []byte) {
	file, err := os.Create(filename)
	Checkerr(err)

	count, err := file.Write(data)
	Checkerr(err)

	fmt.Printf("success file write %d bytes\n",count)
}