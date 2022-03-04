package dto

import "prmlk.com/nextdebug/model"

type UserDto struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		NickName: user.NickName,
		Avatar:   user.Avatar,
	}
}
