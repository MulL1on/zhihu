# 用户

## 注册

### 请求路径

```http
POST /api/user/register
```

### header

无

### 请求参数

| 名称     | 位置 | 类型   | 必选 | 说明   |
| -------- | ---- | ------ | ---- | ------ |
| username | body | string | 是   | 用户名 |
| username | body | string | 是   | 密码   |

### 返回参数

无

### 返回示例

```JSON
//注册成功
{
    "status":200
    "info":"register succeessful"
}
//用户名已被注册
{
    "status":400
    "info":"username alreaddy exist"
}
//用户名不符合规范/用户名包含敏感词汇
{
    "status":400
    "info":"username is illegal"
}
```

## 登录(获取token)

### 请求路径

```http
GET /api/user/token
```

### header

无

### 请求参数

| 名称      | 位置  | 类型   | 必选 | 说明               |
| --------- | ----- | ------ | ---- | ------------------ |
| username  | query | string | 是   | 用户名/手机号/邮箱 |
| paassword | query | string | 是   | 密码               |

### 返回参数

| 字段名        | 必选 | 类型          | 说明          |
| ------------- | ---- | ------------- | ------------- |
| refresh_token | 是   | Bearer $token | refresh_token |
| token         | 是   | Bearer $token | token         |

### 返回示例

```json
//登录成功
{
    "status":200
    "info":"login successfully"
    "data":{
       "refresh_token":"{refresh_token}"
       "token":"{token}"
}
}
//密码错误
{
    "status":400
    "info":"password incorrect"
}
//用户不存在
{
    "status":"400"
    "info":"user is not exist"
}
```

## 刷新token

### 请求路径

```http
GET /api/user/token/refresh
```

### header

# 博客

