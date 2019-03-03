package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/grandcolline/clean-arch-demo/entity"
)

// UserPresenter ユーザプレゼンタ
type UserPresenter struct {
	Writer http.ResponseWriter
}

// func NewUserPresenter(w http.ResponseWriter) usecase.UserOutputPort {
// 	return &UserPresenter{
// 		w: w,
// 	}
// }

// User ユーザレスポンスの構造体
type User struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// RenderUser ユーザをjsonでかえす
func (p *UserPresenter) RenderUser(u *entity.User) error {
	user := User{u.ID, u.Name, u.Email}
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(p.Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.Write(res)
	return nil
}

// RenderUserList ユーザのリストをjsonでかえす
func (p *UserPresenter) RenderUserList(us *[]entity.User) error {
	var users []User
	for _, u := range *us {
		users = append(users, User{u.ID, u.Name, u.Email})
	}
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(p.Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.Write(res)
	return nil
}
