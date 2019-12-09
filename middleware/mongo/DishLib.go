package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DishLib struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`             //菜品名称
	Supplier   string             `json:"supplier" bson:"supplier"`     //供应商
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间

}
