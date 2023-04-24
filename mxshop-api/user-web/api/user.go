package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/global"
	"net/http"
	"time"

	"mxshop-api/user-web/global/reponse"
	"mxshop-api/user-web/proto"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换为http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{"msg": e.Message()})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "用户服务不可用"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "其他错误"})
			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context) {
	ip := global.ServerConfig.UserSrvInfo.Host
	port := global.ServerConfig.UserSrvInfo.Port
	// 拨号连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList 连接 [用户服务失败]", "msg", err.Error())
	}
	// 生成client并调用接口
	userSrvClient := proto.NewUserClient(userConn)

	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询用户列表失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {

		// 使用定义好的struct
		user := reponse.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			//Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)),
			Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2000-10-10"),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		// 使用map的方式
		//data := make(map[string]interface{})

		//data["id"] = value.Id
		//data["name"] = value.NickName
		//data["birthday"] = value.BirthDay
		//data["gender"] = value.Gender
		//data["mobile"] = value.Mobile

		result = append(result, user)
	}

	ctx.JSON(http.StatusOK, result)
}
