package handler

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"go_srvs/user_srv/global"
	"go_srvs/user_srv/model"
	"go_srvs/user_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UserServer struct{}

func Model2Response(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		PassWord: user.PassWord,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	fmt.Printf("%v", userInfoRsp)
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, request *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(request.Pn), int(request.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := Model2Response(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: request.Mobile}).First(&user)

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRsp := Model2Response(user)
	return &userInfoRsp, nil

}

func (s *UserServer) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户找不到")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userInfoRes := Model2Response(user)
	return &userInfoRes, nil
}

func (s *UserServer) CreateUser(ctx context.Context, request *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: request.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已经存在")
	}
	user.Mobile = request.Mobile
	user.NickName = request.NickName
	// 密码加密
	options := &password.Options{16, 100, 32, sha256.New}
	salt, encodedPwd := password.Encode("generic password", options)
	user.PassWord = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoRsp := Model2Response(user)
	return &userInfoRsp, nil
}

func (s UserServer) UpdateUser(ctx context.Context, request *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	result := global.DB.First(&user, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	birthDay := time.Unix(int64(request.BirthDay), 0)
	user.NickName = request.NickName
	user.Birthday = &birthDay
	user.Gender = request.Gender
	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) CheckPassword(ctx context.Context, request *proto.CheckPasswordInfo) (*proto.CheckResponse, error) {
	options := &password.Options{16, 100, 32, sha256.New}
	pwdInfo := strings.Split(request.EncryptedPassword, "$")
	check := password.Verify(request.PassWord, pwdInfo[2], pwdInfo[3], options)
	return &proto.CheckResponse{Success: check}, nil
}
