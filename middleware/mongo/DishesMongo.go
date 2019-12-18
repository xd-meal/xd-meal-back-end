package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DishesMongo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`             //菜品名称
	Supplier   string             `json:"supplier" bson:"supplier"`     //供应商
	TypeA      int                `json:"typeA" bson:"typeA"`           //1-lunch 2-dinner
	TypeB      int                `json:"typeB" bson:"typeB"`           //1-自助餐 2-简餐
	MealDay    string             `json:"mealDay"  bson:"mealDay"`      //用餐日
	MealNum    int                `json:"mealNum" bson:"mealNum"`       //菜品编号
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
	UpdateTime time.Time          `json:"updateTime" bson:"updateTime"` //创建时间
	Status     int                `json:"status" bson:"status"`         //0-开启 1-关闭
}

func (d DishesMongo) CreateRow() interface{} {
	return createRow(d, "meal", "dishes")
}

func (d DishesMongo) FindAll(filter bson.M) []bson.M {
	return FindAllSelected(filter, "meal", "dishes")
}
