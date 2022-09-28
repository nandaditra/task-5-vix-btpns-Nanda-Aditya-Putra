package database

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserUpdateData struct {
	ID        uint64                `json:"id" form:"id" binding:"required"`
	Username  string                `json:"username" form:"username" binding:"required"`
	Email     string                `json:"email" form:"email" binding:"required"`
	Password  string                `json:"password" form:"password" binding:"required" validate:"min:6"`
	CreatedAt timestamppb.Timestamp `json:"created_at" form:"created_at"`
	UpdatedAt timestamppb.Timestamp `json:"updated_at" form:"updated_at"`
}
