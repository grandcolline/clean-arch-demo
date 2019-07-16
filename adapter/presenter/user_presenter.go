package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/grandcolline/clean-arch-demo/usecase"
)

// UserPresenter ユーザプレゼンター
type UserPresenter struct {
	writer http.ResponseWriter
}

// NewUserPresenter ユーザプレゼンターの作成
func NewUserPresenter(w http.ResponseWriter) usecase.UserOutputPort {
	return &UserPresenter{
		writer: w,
	}
}

// User ユーザレスポンスの構造体
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*
RenderUser はユーザをjsonでかえします。

Example Response:
	{
		"id": 1,
		"name": "John Lennon",
		"email": "john@example.com"
	}
*/
func (p *UserPresenter) RenderUser(u *entity.User) error {
	user := User{u.UUID, u.Name, u.Email}
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(p.writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	p.writer.Header().Set("Content-Type", "application/json")
	p.writer.Write(res)
	return nil
}

/*
RenderUserList はユーザのリストをjsonでかえします。

Example Response:
	[
		{
			"id": 1,
			"name": "John Lennon",
			"email": "john@example.com"
		},
		{
			"id": 2,
			"name": "Paul Mccartney",
			"email": "paul@example.com"
		}
	]
*/
func (p *UserPresenter) RenderUserList(us *[]entity.User) error {
	var users []User
	for _, u := range *us {
		users = append(users, User{u.UUID, u.Name, u.Email})
	}
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(p.writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	p.writer.Header().Set("Content-Type", "application/json")
	p.writer.Write(res)
	return nil
}
