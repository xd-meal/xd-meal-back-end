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
	Dsc       string             `json:"dsc" bson:"dsc"`               //菜品明细
	Supplier  string             `json:"supplier" bson:"supplier"`     //供应商
	TypeA     int32              `json:"typeA" bson:"typeA"`           //1-lunch 2-dinner 3-breakfast
	MealDay   string             `json:"mealDay"  bson:"mealDay"`      //用餐日
	MealNum   int32              `json:"mealNum" bson:"mealNum"`       //菜品编号
	OrderTime time.Time          `json:"updateTime" bson:"updateTime"` //订餐时间
	Status    int32              `json:"status" bson:"status"`         //0-订餐 1-取餐
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

//可选的午餐和晚餐
func (d UserDishesMongo) GetUserDishesByOrdered(uid string) []bson.M {
	switches := Switches{}.FindOne(bson.M{"name": "order"})
	filter2 := bson.M{"uid": uid, "mealDay": bson.M{"$gte": switches["startMealDay"], "$lte": switches["endMealDay"]}, "typeA": bson.M{"$in": []int32{1, 2}}}
	return d.FindAll(filter2)
}

func (d UserDishesMongo) GetTotalByOrdered() []bson.M {
	switches := Switches{}.FindOne(bson.M{"name": "order"})
	pipeline := []bson.M{
		{
			"$match": bson.M{"mealDay": bson.M{"$gte": switches["startMealDay"], "$lte": switches["endMealDay"]}},
		},
		{
			"$group": bson.M{"_id": "$name", "supplier": bson.M{"$first": "$supplier"}, "total": bson.M{"$sum": 1}},
		},
	}
	return FindByPipe(pipeline, "meal", "userDishes")
}
