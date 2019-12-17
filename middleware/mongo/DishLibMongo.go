package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DishLib struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`             //菜品名称
	Supplier   string             `json:"supplier" bson:"supplier"`     //供应商
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
}

func (d DishLib) CreateRow() interface{} {
	return createRow(d, "meal", "dishLib")
}

func (d DishLib) UpdateAll(filter, update bson.M) interface{} {
	return UpdateAll(filter, update, "meal", "dishLib")
}

func (d DishLib) FindAll(filter bson.M) []bson.M {
	return FindAllSelected(filter, "meal", "dishLib")
}

func (d DishLib) FindOne(filter bson.M) bson.M {
	return FindOneSelected(filter, "meal", "dishLib")
}
