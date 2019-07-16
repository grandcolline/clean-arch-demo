package entity

import "github.com/google/uuid"

// User ユーザエンティティ
type User struct {
	UUID  string // ユーザUUID
	Name  string // ユーザ名
	Email string // メールアドレス
}

// SetNewUUID はユニークなユーザUUIDを発行します.
func (e *User) SetNewUUID() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	e.UUID = id.String()
	return nil
}
