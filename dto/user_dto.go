package dto

import "prmlk.com/nextdebug/model"

type UserDto struct {
	Name string `json:"name"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
	}
}
