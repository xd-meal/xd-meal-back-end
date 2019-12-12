package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Switches struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Enable     int                `json:"enable" bson:"enable"`
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
	UpdateTime time.Time          `json:"updateTime" bson:"updateTime"` //更新时间
}

func (d Switches) CreateRow() interface{} {
	return createRow(d, "meal", "switches")
}

func (d Switches) UpdateAll(filter, update bson.M) interface{} {
	return UpdateAll(filter, update, "meal", "switches")
}

func (d Switches) FindAll(filter bson.M) []bson.M {
	return FindAllSelected(filter, "meal", "switches")
}

func (d Switches) FindOne(filter bson.M) bson.M {
	return FindOneSelected(filter, "meal", "switches")
}
