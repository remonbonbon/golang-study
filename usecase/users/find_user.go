package users

import (
	"example.com/golang-study/common"
	"example.com/golang-study/domain/model"
	"example.com/golang-study/infra/repository"
)

func FindUser(id string) (*model.User, error) {
	repo := repository.NewUsersRepository()

	u, ok, err := repo.FindUser(id)
	if err != nil {
		return nil, common.SystemError("ユーザーが取得に失敗しました。", err)
	}
	if !ok {
		return nil, common.NotFound("ユーザーが見つかりません。", nil)
	}
	return u, nil
}
