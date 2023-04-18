package handler

import (
	"context"
	"gorm.io/gorm"
	"micro-server/user_srv/global"
	"micro-server/user_srv/model"
	"micro-server/user_srv/proto"
)

// Paginate 通用分页方法
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

// Model2Response 将model转为response
func Model2Response(user model.User) proto.UserInfoResponse {
	// 在grpc的message中，字段有默认值，不能随意赋值nil
	//
	userInfoRes := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRes.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRes
}

type UserServer struct{}

// GetUserList 获取用户列表
func (s *UserServer) GetUserList(ctx context.Context, req proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := Model2Response(user)
		rsp.Data = append(rsp.Data, &userInfoRsp) // 这里&userInfoRsp需要注意，要先赋值，再使用
	}

	return rsp, nil
}
