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

## 检查用户是否登录
> 请求方式
```
GET
```
> 路由
```
/api/v1/CheckUserLogin
```
> 入参

> 出参
```
{
    "code": 200,
    "data": "5e09b5f57d469f1dd46dc5e8",
    "msg": "已登陆"
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

## 用户更新密码
> 请求方式
```
GET
```
> 路由
```
/api/v1/ResetPasswordByUser
```

> 入参
```
{
    "oldPassword":"shenzhuo1234"
	"password":"shenzhuo123"
}

```

> 出参
```
{
    "code": 200,
    "msg": "密码修改成功，请重新登录"
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

## 差评
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/EvalDish
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|enable|int|订餐开关0-关闭，1-开启|Y
```
{
	"id":"5df8ac95822b24bfb38d8c26"
}
{
	"id":"5df9af62fc8017cce5772a15",//用户菜品_id
	"eval":false//bool类型
}
```
> 出参
```
{
    "code": 400,
    "msg": "菜品不存在或已提交评价"
}
```

## 生成取餐二维码
> 请求方式
```
GET
```
> 路由
```
/api/v1/GetDishCode
```
> 入参

无
```
{
	"token":"C5df9eddf11d06036fb5a76a9"
}
```
> 出参
```
{
    "code": 200,
    "data": {
        "_id": "5df9eddf11d06036fb5a76a9",
        "badEval": false,
        "dishId": "5df88cb0717a45dd5764d7c1",
        "mealDay": "2019-12-18",
        "mealNum": 0,
        "name": "东北手卷春饼",
        "status": 1,
        "supplier": "大宁东北水饺",
        "typeA": 2,
        "uid": "5df75734dcd9fa4184580f55",
        "updateTime": "2019-12-18T17:14:07.433+08:00"
    },
    "msg": "不能重复取餐"
}
```

## 二维码取餐
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/ScanDishCode
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|token|string|二维码|Y
```
{
	"token":"C5df9eddf11d06036fb5a76a9"
}
```
> 出参
```
{
    "code": 200,
    "data": {
        "_id": "5df9eddf11d06036fb5a76a9",
        "badEval": false,
        "dishId": "5df88cb0717a45dd5764d7c1",
        "mealDay": "2019-12-18",
        "mealNum": 0,
        "name": "东北手卷春饼",
        "status": 1,
        "supplier": "大宁东北水饺",
        "typeA": 2,
        "uid": "5df75734dcd9fa4184580f55",
        "updateTime": "2019-12-18T17:14:07.433+08:00"
    },
    "msg": "不能重复取餐"
}
```


## 根据邮箱搜索用户
> 请求方式
```
GET
```
> 路由
```
/api/v1/GetUserByEmail
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|email|string|邮箱|Y

> 出参
```
{
    "code": 200,
    "data": {
        "email": "937728009@qq.com",
        "id": "5df75734dcd9fa4184580f55",
        "name": "李小磊"
    },
    "msg": "success"
}
```


## 单独加菜
> 请求方式
```
POST-application/json
```
> 路由
```
/api/v1/AddMenuSingle
```
> 入参

|参数|类型|含义|是否必须 
|:----- |:----- |:----- |:----- | 
|uid|string|用户id|Y
```
{
	"uid":"5df75734dcd9fa4184580f56"
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

## 本次选餐统计
> 请求方式
```
GET
```
> 路由
```
/api/v1/GetMealTotal
```
> 入参

无

> 出参
```
{
    "code": 200,
    "data": [
        {
            "_id": "健康水果轻食",
            "supplier": "多点沙拉",
            "total": 1
        },
        {
            "_id": "爆炒子姜鸭",
            "supplier": "红采餐饮",
            "total": 1
        },
        {
            "_id": "自助餐",
            "supplier": "园沁餐厅",
            "total": 14
        }
    ],
    "msg": "ok"
}
```