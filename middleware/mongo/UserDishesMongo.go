package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserDishes struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Uid       string             `json:"uid" bson:"uid"`               //用户id
	DishId    string             `json:"dishId" bson:"dishId"`         //菜单id
	MealDay   string             `json:"mealDay"  bson:"mealDay"`      //用餐日
	MealNum   int                `json:"mealNum" bson:"mealNum"`       //菜品编号
	OrderTime time.Time          `json:"updateTime" bson:"updateTime"` //订餐时间
	Status    int                `json:"status" bson:"status"`         //0-订餐 1-取餐
	BadEval   int                `json:"BadEval" bson:"BadEval"`       //差评
}

func (d UserDishes) CreateRow() interface{} {
	return createRow(d, "meal", "userDishes")
}

func (d UserDishes) UpdateAll(filter, update bson.M) interface{} {
	return UpdateAll(filter, update, "meal", "userDishes")
}
