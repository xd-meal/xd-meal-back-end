package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserMongo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Email      string             `json:"email" bson:"email"`
	PassWord   string             `json:"password" bson:"password"` //登录密码
	Type       int                `json:"type" bson:"type"`         //1-管理员 2-外部用户 3-微信用户
	Depart     string             `json:"depart" bson:"depart"`
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
	Unique     string             `json:"unique" bson:"unique"`         //导入用户填邮箱，微信用户填userid
}

func (d UserMongo) CreateRow() interface{} {
	return createRow(d, "meal", "user")
}

func (d UserMongo) UpdateAll(filter, update bson.M) int64 {
	return UpdateAll(filter, update, "meal", "user")
}

func (d UserMongo) FindOne(filter bson.M) bson.M {
	return FindOneSelected(filter, "meal", "user")
}
