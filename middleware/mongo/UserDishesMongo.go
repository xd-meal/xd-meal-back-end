package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserDishesMongo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Uid       string             `json:"uid" bson:"uid"`               //用户id
	DishId    string             `json:"dishId" bson:"dishId"`         //菜单id
	Name      string             `json:"name" bson:"name"`             //菜品名称
	Supplier  string             `json:"supplier" bson:"supplier"`     //供应商
	MealDay   string             `json:"mealDay"  bson:"mealDay"`      //用餐日
	MealNum   int                `json:"mealNum" bson:"mealNum"`       //菜品编号
	OrderTime time.Time          `json:"updateTime" bson:"updateTime"` //订餐时间
	Status    int                `json:"status" bson:"status"`         //0-订餐 1-取餐
	BadEval   bool               `json:"badEval" bson:"badEval"`       //差评
}

func (d UserDishesMongo) CreateRow() interface{} {
	return createRow(d, "meal", "userDishes")
}

func (d UserDishesMongo) UpdateAll(filter, update bson.M) int64 {
	return UpdateAll(filter, update, "meal", "userDishes")
}

func (d UserDishesMongo) FindOne(filter bson.M) bson.M {
	return FindOneSelected(filter, "meal", "userDishes")
}

func (d UserDishesMongo) FindAll(filter bson.M) []bson.M {
	return FindAllSelected(filter, "meal", "userDishes")
}
