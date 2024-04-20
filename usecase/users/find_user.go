package users

import (
	"example.com/golang-study/common"
	"example.com/golang-study/common/message"
	"example.com/golang-study/domain/model"
	"example.com/golang-study/domain/repository"
)

func FindUser(repo repository.UsersRepository, id string) (*model.User, error) {
	u, ok, err := repo.FindUser(id)
	if err != nil {
		return nil, common.SystemError(message.NotFound("ユーザー取得"), err)
	}
	if !ok {
		return nil, common.NotFound(message.NotFound("ユーザー"), nil)
	}
	return u, nil
}
