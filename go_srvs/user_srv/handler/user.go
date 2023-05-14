package handler

import (
	"context"
	"go_srvs/user_srv/global"
	"go_srvs/user_srv/model"
	"go_srvs/user_srv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct{}

func Model2Response(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.PassWord,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
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
