package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccessToken struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Token      string             `json:"token" bson:"token"`
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
	UpdateTime time.Time          `json:"updateTime" bson:"updateTime"` //更新时间
}

func (d AccessToken) CreateRow() interface{} {
	return createRow(d, "meal", "accessToken")
}

func (d AccessToken) UpdateAll(filter, update bson.M) int64 {
	return UpdateAll(filter, update, "meal", "accessToken")
}

func (d AccessToken) FindAll(filter bson.M) []bson.M {
	return FindAllSelected(filter, "meal", "accessToken")
}

func (d AccessToken) FindOne(filter bson.M) bson.M {
	return FindOneSelected(filter, "meal", "accessToken")
}
