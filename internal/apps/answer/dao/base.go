package dao

import (
	"go-answer/pkg/storages"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Base struct {
	Tx *gorm.DB
}

// DB 获取DB
func (base *Base) DB(ctx *gin.Context) (db *gorm.DB) {
	if base.Tx != nil {
		return base.Tx.WithContext(ctx)
	}

	db = storages.DBDemo.WithContext(ctx)
	return
}
