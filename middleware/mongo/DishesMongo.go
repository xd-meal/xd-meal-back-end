package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DishesMongo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"` //菜品名称
	Dsc        string             `json:"dsc" bson:"dsc"`
	Supplier   string             `json:"supplier" bson:"supplier"`     //供应商
	TypeA      int32              `json:"typeA" bson:"typeA"`           //1-lunch 2-dinner
	TypeB      int32              `json:"typeB" bson:"typeB"`           //1-自助餐 2-简餐
	MealDay    string             `json:"mealDay"  bson:"mealDay"`      //用餐日
	MealNum    int                `json:"mealNum" bson:"mealNum"`       //菜品编号
	CreateTime time.Time          `json:"createTime" bson:"createTime"` //创建时间
	UpdateTime time.Time          `json:"updateTime" bson:"updateTime"` //创建时间
	Status     int32              `json:"status" bson:"status"`         //0-开启 1-关闭
}

func (d DishesMongo) CreateRow() interface{} {
	return createRow(d, "meal", "dishes")
}

func (d DishesMongo) FindAll(filter bson.M) []bson.M {
	return FindAllSelected(filter, "meal", "dishes")
}

//可选的自助餐
func (d DishesMongo) GetOptionalBuffet() []bson.M {
	switches := Switches{}.FindOne(bson.M{"name": "order"})
	filter2 := bson.M{"mealDay": bson.M{"$gte": switches["startMealDay"], "$lte": switches["endMealDay"]}, "typeB": int32(1)}
	return d.FindAll(filter2)
}
