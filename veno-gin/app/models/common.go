package models

import (
	"gorm.io/gorm"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: common
 * @Version: ...
 * @Date: 2022-10-31 17:49:25
 */

// 自增ID主键

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 软删除

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
