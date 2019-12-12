package mongo

import (
	"context"
	"fmt"
	"github.com/xd-meal-back-end/pkg/logging"
	"github.com/xd-meal-back-end/pkg/setting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"strings"
	"time"
)

type MongoUtils struct {
	Con *mongo.Client
	Db  *mongo.Database
}

func (o *MongoUtils) OpenConn() (con *mongo.Client) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", setting.MongoSetting.User, setting.MongoSetting.Password, setting.MongoSetting.Host,
		setting.MongoSetting.Port)
	// Initialize and Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = con.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("mongo.Setup err: %v", err)
		return
	}
	//fmt.Println("connected to nosql database:", uri)
	o.Con = con
	return con
}

func Setup() {
	m := MongoUtils{}
	(&m).OpenConn()
}

func (o *MongoUtils) SetDb(db string) {
	if o.Con == nil {
		panic("连接为空...")
	}
	o.Db = o.Con.Database(db)
}

func (o *MongoUtils) UpdateAll(col string, filter bson.M, update bson.M) (interface{}, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("没有初始化连接和数据库信息！")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := table.UpdateMany(ctx, filter, update)
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	return result.ModifiedCount, nil
}

func (o *MongoUtils) FindOne(col string, filter bson.M) (bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("没有初始化连接和数据库信息！")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := table.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *MongoUtils) FindMore(col string, filter bson.M) ([]bson.M, error) {
	if o.Db == nil || o.Con == nil {
		return nil, fmt.Errorf("没有初始化连接和数据库信息！")
	}
	table := o.Db.Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	cur, err2 := table.Find(ctx, filter)
	if err2 != nil {
		fmt.Print(err2)
		return nil, err2
	}
	defer cur.Close(ctx)
	var resultArr []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err3 := cur.Decode(&result)
		if err3 != nil {
			return nil, err3
		}
		resultArr = append(resultArr, result)
	}
	return resultArr, nil
}

func createRow(insert interface{}, dbName string, tableName string) interface{} {
	utils := MongoUtils{}
	utils.OpenConn()
	defer utils.Con.Disconnect(context.Background())
	utils.SetDb(dbName)
	col := utils.Db.Collection(tableName)
	suc := true
	_ = utils.Con.UseSession(context.Background(), func(ses mongo.SessionContext) error {
		_ = ses.StartTransaction()
		_, err := col.InsertOne(context.Background(), insert)
		if err != nil {
			suc = false
			_ = ses.AbortTransaction(ses)
			fmt.Println(err)
			return err
		}
		_ = ses.CommitTransaction(ses)
		return nil
	})

	if !suc {
		return "err"
	}
	//b, _ := json.Marshal(insert)
	//res := string(b)
	return insert
}

func FindAllSelected(filter bson.M, dbName string, tableName string) []bson.M {
	utils := MongoUtils{}
	utils.OpenConn()
	utils.SetDb(dbName)
	result, err := utils.FindMore(tableName, bson.M(filter))
	if err != nil {
		logging.Info(err)
	}
	return result
}

func FindOneSelected(filter bson.M, dbName string, tableName string) bson.M {
	utils := MongoUtils{}
	utils.OpenConn()
	utils.SetDb(dbName)
	result, err := utils.FindOne(tableName, bson.M(filter))
	if err != nil {
		logging.Info(err)
	}
	return result
}

func UpdateAll(filter bson.M, update bson.M, dbName string, tableName string) interface{} {
	utils := MongoUtils{}
	utils.OpenConn()
	utils.SetDb(dbName)
	result, err := utils.UpdateAll(tableName, filter, update)
	if err != nil {
		logging.Info(err)
	}
	return result
}

func (o *MongoUtils) createUniqueIndex(col string, keys ...string) {
	table := o.Db.Collection(col)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	indexView := table.Indexes()
	keysDoc := bsonx.Doc{}

	// 复合索引
	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}

	// 创建索引
	result, err := indexView.CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keysDoc,
			Options: options.Index().SetUnique(true),
		},
		opts,
	)
	if result == "" || err != nil {
		log.Fatalf("EnsureIndex error:%v", err)
	}
}
