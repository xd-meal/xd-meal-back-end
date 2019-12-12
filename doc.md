#App接口
## 登录
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/Login
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|email|string|邮箱|Y
|password|string|密码|Y
```
{
	"email":"937728009@qq.com",
	"password":"xd123456"
}
```
> 出参
```
{
    "code": 200,
    "msg": "登录成功"
}
```

## 登出
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/LoginOut
```
> 入参

> 出参
```
{
    "code": 200,
    "data": "",
    "msg": "退出成功"
}
```


## 获取每周菜品
> 请求方式
```
GET
```
> 路由
```
/api/v1/GetDishes
```
> 入参

> 出参
```
{
    "code": 200,
    "data": [
        {
            "_id": "5dea015e26a606122cf74d4c",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d4d",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d4e",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d4f",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d50",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d51",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d52",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d53",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d54",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d55",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d56",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d57",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d58",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d59",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5a",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5b",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5c",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5d",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5e",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5f",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d60",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d61",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d62",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d63",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d64",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d65",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d66",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-08",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d67",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d68",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d69",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d6a",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d6b",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d6c",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d6d",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d6e",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d6f",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d70",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d71",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d72",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d73",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d74",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d75",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d76",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d77",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d78",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-10",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d79",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d7a",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d7b",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d7c",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d7d",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d7e",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d7f",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d80",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d81",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-11",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d82",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d83",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d84",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 2,
            "name": "爆炒子姜鸭",
            "status": 0,
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d85",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 3,
            "name": "水煮肉片",
            "status": 0,
            "supplier": "颂饭",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d86",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 4,
            "name": "特色干拌套餐",
            "status": 0,
            "supplier": "觅哥麻辣烫",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d87",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d88",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 1,
            "name": "牛肉串串",
            "status": 0,
            "supplier": "卤人甲",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d89",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 2,
            "name": "酸菜爱心鸡排堡",
            "status": 0,
            "supplier": "德克士",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015f26a606122cf74d8a",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-12",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        }
    ],
    "msg": "ok"
}
```

## 订餐
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/OrderDishes
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|dishIds|array|菜品id|Y

```
{
"dishIds":[
"5dea015e26a606122cf74d4c",
"5dea015e26a606122cf74d54",
"5dea015e26a606122cf74d5a",
"5dea015e26a606122cf74d68"
]
}
```
> 出参
```
{
    "code": 200,
    "data": "",
    "msg": "ok"
}
```

## 我的点餐
> 请求方式
```
GET
```
> 路由
```
/api/v1/GetOrderDishes
```
> 入参


> 出参
```
{
    "code": 200,
    "data": [
        {
            "_id": "5dea015e26a606122cf74d4c",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d54",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-06",
            "mealNum": 3,
            "name": "东北手卷春饼",
            "status": 0,
            "supplier": "大宁东北水饺",
            "typeA": 2,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d5a",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-07",
            "mealNum": 0,
            "name": "自助餐",
            "status": 0,
            "supplier": "园沁餐厅",
            "typeA": 2,
            "typeB": 1,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        },
        {
            "_id": "5dea015e26a606122cf74d68",
            "createTime": "2019-12-06T15:20:32.751+08:00",
            "mealDay": "2019-12-09",
            "mealNum": 1,
            "name": "健康水果轻食",
            "status": 0,
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "updateTime": "2019-12-06T15:20:32.751+08:00"
        }
    ],
    "msg": "ok"
}
```

#后台接口

## 导入外部用户
> 请求方式
```
POST-application/form-data
```
> 路由
```
/api/v1/ImportUser
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|file|file|要导入的excel|Y

> 出参
```
{
    "data": [
        {
            "ID": "000000000000000000000000",
            "name": "李小磊",
            "email": "937728009@qq.com",
            "password": "xd123456",
            "type": 2,
            "depart": "心动-测试支撑部外派",
            "createTime": "2019-12-12T15:00:46.315669+08:00"
        },
        {
            "ID": "000000000000000000000000",
            "name": "李泉",
            "email": "erichere@qq.com",
            "password": "xd123456",
            "type": 2,
            "depart": "心动-测试支撑部外派",
            "createTime": "2019-12-12T15:00:46.315669+08:00"
        }
    ],
    "msg": null
}
```

## 预览导入菜单
> 请求方式
```
POST-application/form-data
```
> 路由
```
/api/v1/ReadMenu
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|file|file|要导入的excel|Y

> 出参
```
{
    "data": [
        {
            "ID": "000000000000000000000000",
            "name": "自助餐",
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "mealDay": "2019-12-06",
            "mealNum": 0,
            "createTime": "2019-12-12T15:06:20.015575+08:00",
            "updateTime": "2019-12-12T15:06:20.015575+08:00",
            "status": 0
        },
        {
            "ID": "000000000000000000000000",
            "name": "健康水果轻食",
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "mealDay": "2019-12-06",
            "mealNum": 1,
            "createTime": "2019-12-12T15:06:20.015575+08:00",
            "updateTime": "2019-12-12T15:06:20.015575+08:00",
            "status": 0
        },
        {
            "ID": "000000000000000000000000",
            "name": "爆炒子姜鸭",
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "mealDay": "2019-12-06",
            "mealNum": 2,
            "createTime": "2019-12-12T15:06:20.015575+08:00",
            "updateTime": "2019-12-12T15:06:20.015575+08:00",
            "status": 0
        }
    ],
    "msg": null
}
```

## 确定导入菜单
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/ImportMenu
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|data|array|excel中的map|Y

```
{
    "data": [
        {
            "ID": "000000000000000000000000",
            "name": "自助餐",
            "supplier": "园沁餐厅",
            "typeA": 1,
            "typeB": 1,
            "mealDay": "2019-12-06",
            "mealNum": 0,
            "createTime": "2019-12-12T15:06:20.015575+08:00",
            "updateTime": "2019-12-12T15:06:20.015575+08:00",
            "status": 0
        },
        {
            "ID": "000000000000000000000000",
            "name": "健康水果轻食",
            "supplier": "多点沙拉",
            "typeA": 1,
            "typeB": 2,
            "mealDay": "2019-12-06",
            "mealNum": 1,
            "createTime": "2019-12-12T15:06:20.015575+08:00",
            "updateTime": "2019-12-12T15:06:20.015575+08:00",
            "status": 0
        },
        {
            "ID": "000000000000000000000000",
            "name": "爆炒子姜鸭",
            "supplier": "红采餐饮",
            "typeA": 1,
            "typeB": 2,
            "mealDay": "2019-12-06",
            "mealNum": 2,
            "createTime": "2019-12-12T15:06:20.015575+08:00",
            "updateTime": "2019-12-12T15:06:20.015575+08:00",
            "status": 0
        }
    ]
}
```

> 出参
```
{
    "code": 200,
    "data": "",
    "msg": "success"
}
```


## 开关订餐
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/EnableOrderSwitch
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|enable|int|订餐开关0-关闭，1-开启|Y
```
{
	"enable": 1,
}
```
> 出参
```
{
    "code": 200,
    "data": 1,//0-关闭 1-开启
    "msg": "开启订餐"
}
```

## 获取订餐开关状态
> 请求方式
```
GET
```
> 路由
```
/api/v1/GetOrderSwitch
```
> 入参

无

> 出参
```
{
    "code": 200,
    "data": 1, //0-关闭 1-开启
    "msg": "success"
}
```