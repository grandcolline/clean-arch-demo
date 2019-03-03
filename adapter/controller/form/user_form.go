package form

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/grandcolline/clean-arch-demo/entity"
)

var emailRe = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User ユーザの入力フォーム
type User struct {
	Name  string
	Email string
}

// Set リクエストからフォームをセットする
func (u *User) Set(r *http.Request) (err error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		// TODO: エラーハンドリングは後で考える
		return
	}
	if err = r.Body.Close(); err != nil {
		// TODO: エラーハンドリングは後で考える
		return
	}
	if err = json.Unmarshal(body, &u); err != nil {
		// TODO: エラーハンドリングは後で考える
		return
	}
	return
}

// Validate バリデーションチェック
func (u *User) Validate() (ok bool, m []string) {
	ok = true
	// Nameのバリデーションチェック
	if len(u.Name) > 20 {
		m = append(m, "名前は20文字以内にしてください。")
		ok = false
	}
	// Emailのバリデーションチェック
	if u.Email != "" && !emailRe.MatchString(u.Email) {
		m = append(m, "メールアドレスが不正です。")
		ok = false
	}
	return
}

// Require 必須チェック
func (u *User) Require() (ok bool, m []string) {
	ok = true
	if u.Name == "" {
		m = append(m, "名前は必須項目です。")
		ok = false
	}
	if u.Email == "" {
		m = append(m, "メールアドレスは必須項目です。")
		ok = false
	}
	return
}

// ToEntity ユーザフォームをユーザエンティティに詰め直す
func (u *User) ToEntity() (e entity.User) {
	e.Name = u.Name
	e.Email = u.Email
	return
}
