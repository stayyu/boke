package handler

import (
	"boke/go/api-server/global"
	"boke/go/api-server/model"
	"boke/go/api-server/proto"
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

func ModelToRsponse(user model.Admin) proto.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清， 哪些字段是有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Role:     int32(user.Role),
	}
	return userInfoRsp
}

func (u *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.Admin
	user.NickName = req.NiceName
	user.Password = req.PassWord
	//密码加密
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	su := global.DB.Create(&user)
	if su.Error != nil {
		return nil, status.Errorf(codes.Internal, su.Error.Error())
	}
	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	//通过id查询用户
	var user model.Admin
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) GetUsername(ctx context.Context, req *proto.NameRequest) (*proto.UserInfoResponse, error) {
	//通过用户名查询用户
	var user model.Admin
	result := global.DB.Where(&model.Admin{NickName: req.NickName}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	//校验密码
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.CheckResponse{Suess: check}, nil
}
