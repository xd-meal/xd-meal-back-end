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
	Type       int                `json:"type" bson:"type"`
	Depart     string             `json:"depart" bson:"depart"`
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
}

func (d UserMongo) CreateRow() interface{} {
	return createRow(d, "meal", "user")
}

func (d UserMongo) UpdateAll(filter, update bson.M) interface{} {
	return UpdateAll(filter, update, "meal", "user")
}

func (d UserMongo) FindOne(filter bson.M) bson.M {
	return FindOneSelected(filter, "meal", "user")
}
