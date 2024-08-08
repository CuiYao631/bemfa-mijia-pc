package user

import (
	"changeme/app/http"
	"context"
	"log"
)

/*
用户模块
[邮箱注册],[邮箱登录],[修改密码],[手机注册],[手机登录],[设置新的AppID和secretKey],[获取所有appID],[删除appID]
*/

// User struct
type User struct {
}

func NewUser() User {
	return User{}
}

// RegisterByEmail 注册邮箱
func (u *User) RegisterByEmail(ctx context.Context, email, password string) {

}

// LoginByEmail 登录邮箱
func (u *User) LoginByEmail(ctx context.Context, email, password string) error {

	form, err := http.PostForm(ctx, "https://pro.bemfa.com/v1/login", map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return err
	}
	log.Println(form)
	return nil
}
