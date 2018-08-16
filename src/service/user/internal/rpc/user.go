package rpc

import (
	userpbgw "github.com/linhlc888/sample/src/proto/gateway/user"
	"github.com/linhlc888/sample/src/service/user/model/user"
	"golang.org/x/net/context"
)

//UserService use this struct to implement rpc function
type UserService struct{}

//Login allow user login with register email and password
func (u *UserService) Login(ctx context.Context, in *userpbgw.LoginRequest) (*userpbgw.Response, error) {
	return &userpbgw.Response{
		Error:   0,
		Message: "Login",
	}, nil
}

//Register allow user register new account
func (u *UserService) Register(ctx context.Context, in *userpbgw.RegisterRequest) (*userpbgw.Response, error) {
	isExisted, err := user.CheckExistingEmail(in.Email)
	/*
		if err != nil {
			return &userpbgw.Response{
				Error:   2222,
				Message: fmt.Sprintf("Error: %s", err),
			}, nil
		}
	*/
	if isExisted {
		return &userpbgw.Response{
			Error:   2222,
			Message: "Given email is existing",
		}, nil
	}
	newuser := user.User{
		Fullname: in.Fullname,
		Email:    in.Email,
		Password: in.Password,
	}
	err = user.Insert(&newuser)
	if err != nil {
		return &userpbgw.Response{
			Error:   2222,
			Message: "Register Error",
		}, nil
	}
	return &userpbgw.Response{
		Error:   0,
		Message: "Register Sucessfull",
	}, nil
}
