package db

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/pkg/conf"
	"gorm.io/gorm"
)

type UserPO struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *UserPO) TableName() string {
	return conf.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*UserPO, error) {
	res := make([]*UserPO, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*UserPO) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*UserPO, error) {
	res := make([]*UserPO, 0)
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}
