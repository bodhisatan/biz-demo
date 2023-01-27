package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/biz-demo/book-shop/app/user/infras/db"
	"github.com/cloudwego/biz-demo/book-shop/app/user/infras/redis"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (s *UserService) CreateUser(req *user.CreateUserReq) error {
	users, err := db.QueryUser(s.ctx, req.GetUserName())
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}

// Cache Aside模式：
// 读：先读cache，不存在则回源，回源如果读到则更新cache；
// 写：更新数据库时，删除cache
func (s *UserService) MGetUser(req *user.MGetUserReq) ([]*user.User, error) {
	ret := make([]*user.User, 0)
	idNotCached := make([]int64, 0)

	userInfoStr, err := redis.MGet(req.GetIds())
	// 降级
	if err != nil || userInfoStr == nil {
		idNotCached = req.Ids
	} else {
		for index, item := range userInfoStr {
			if item == "" {
				idNotCached = append(idNotCached, req.GetIds()[index])
			} else {
				ret = append(ret, s.getDtoFromString(item))
			}
		}
	}

	users, err := db.MGetUsers(s.ctx, idNotCached)
	if err != nil {
		return nil, err
	}

	for _, userModel := range users {
		userCur := &user.User{
			UserId:   int64(userModel.ID),
			UserName: userModel.UserName,
		}
		ret = append(ret, userCur)

		str, _ := sonic.MarshalString(userCur)

		redis.Upsert(int64(userModel.ID), str)
	}
	return ret, nil
}

func (s *UserService) CheckUser(req *user.CheckUserReq) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}

func (s *UserService) getDtoFromString(userInfo string) *user.User {
	ret := &user.User{}
	_ = sonic.UnmarshalString(userInfo, ret)
	return ret
}
