package std

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"os"
)

/*-------------------------------------
 *暗号化用の乱数を生成する
 *引数: r 乱数範囲
--------------------------------------*/
func GenCryptoNum(r int64) *big.Int {
	res, err := rand.Int(rand.Reader,big.NewInt(int64(r)))
	Checkerr(err)

	return res
}

/*-------------------------------------
 *SHA256ハッシュを生成する
 *引数: originalstr ハッシュ化対象文字列
 *戻り値: ハッシュ文字列
--------------------------------------*/
func Tohash(originalstr string) string {
	salt := os.Getenv("SALT")
	hashstr := []byte(originalstr + salt)

	hash_sha256 := sha256.Sum256(hashstr)

	return hex.EncodeToString(hash_sha256[:])
}

/*-------------------------------------
 *ハッシュ化文字列を照合する
 *引数: hash ハッシュ化文字列
 *     userid オリジン文字列
 *戻り値: 判定結果
         true 一致
		 false 一致せず
--------------------------------------*/
func IsMatch(hash, userid string) bool {
	salt := os.Getenv("SALT")
	hashstr := []byte(userid + salt)
	hash_sha256 := sha256.Sum256(hashstr)

	if hash == hex.EncodeToString(hash_sha256[:]) {
		return true
	}
	return false
}