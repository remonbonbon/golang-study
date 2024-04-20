package users

import (
	"example.com/golang-study/common"
	"example.com/golang-study/domain/model"
	"example.com/golang-study/domain/repository"
)

func FindUser(repo repository.UsersRepository, id string) (*model.User, error) {
	u, ok, err := repo.FindUser(id)
	if err != nil {
		return nil, common.SystemError("ユーザーが取得に失敗しました。", err)
	}
	if !ok {
		return nil, common.NotFound("ユーザーが見つかりません。", nil)
	}
	return u, nil
}
