package dtouser

import (
	"go-answer/internal/apps/answer/object/objcommon"
	"go-answer/internal/apps/answer/object/objuser"
)

type UserCreateResp struct {
	ID uint `json:"id"` // 数据自增id
}

type UserDetailResp struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objuser.UserBaseInfo
	objcommon.OperatorBaseInfo
}

type UserPageListItem struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objuser.UserBaseInfo
	objcommon.OperatorBaseInfo
}

type UserPageListResp struct {
	List  []UserPageListItem `json:"list"`  // 数据列表
	Total int64              `json:"total"` // 数据总条数
}
