# 验证

## 注册

### 请求路径

```http
POST /api/user/register
application/json
```

### header

无

### 请求参数

| 名称     | 位置 | 类型   | 必选 | 说明   |
| -------- | ---- | ------ | ---- | ------ |
| username | body | string | 是   | 用户名 |
| password | body | string | 是   | 密码   |

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
    "info":"username already exist"
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
POST /api/user/token
application/json
```

### header

无

### 请求参数

| 名称      | 位置 | 类型   | 必选 | 说明               |
| --------- | ---- | ------ | ---- | ------------------ |
| username  | body | string | 是   | 用户名/手机号/邮箱 |
| paassword | body | string | 是   | 密码               |

### 返回参数

| 字段名 | 必选 | 类型          | 说明  |
| ------ | ---- | ------------- | ----- |
| token  | 是   | Bearer $token | token |

### 返回示例

**正确返回**

```json
//登录成功
{
    "status":200
    "info":"login successfully"
    "data":{
       "refresh_token":string
       "token":string
}
}
```

**错误返回**

```json
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

## 发送验证码

### 请求路径

```http
POST /api/user/code
application/x-www-form-urlencoded
```

### header

无

### 请求参数

| 名称 | 位置 | 类型   | 必选 | 说明 |
| ---- | ---- | ------ | ---- | ---- |
| mail | body | string | 是   | 邮箱 |

### 返回参数

无

### 返回示例

**正确返回**

```json
//发送验证码成功
{
   "status" :200
   "info":"send code successfullly"
}
```

**错误返回**

```json
//邮箱已被注册
{
"status":200
"info":"mail is already signed"
}

//邮箱不正确
{
  "status":200
  "info":"mail is invalid"
}
```

## 退出登录

### 请求路径

```http
GET /api/user/logout
```

### header

| 字段名        | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

无

### 返回参数

无

### 返回示例

**正确返回**

```json
{
 "status":200
 "msg":"logout successfully"
}
```



# 文章

## 创建文章草稿

### 请求路径

```http
POST /api/content/article_draft/create
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称          | 位置 | 必选 | 类型   | 说明     |
| ------------- | ---- | ---- | ------ | -------- |
| content       | body | 是   | string | 文章正文 |
| brief_content | body | 是   | string | 文章摘要 |
| title         | body | 是   | string | 文章标题 |
| tag_ids       | body | 是   | array  | 标签     |
| category_id   | body | 是   | string | 分类     |

### 返回参数

| 名称     | 必选 | 类型   | 说明       |
| -------- | ---- | ------ | ---------- |
| draft_id | 是   | string | 文章草稿id |
| user_id  | 是   | string | 用户id     |

### 返回示例

**正确返回**

```json
{
"status":200
"message":"create draft successfully"
}
```

## 删除草稿

### 请求路径

```http
DELETE /api/content/draft
```

### header

| 字段名        | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称     | 位置 | 必选 | 类型   | 说明     |
| -------- | ---- | ---- | ------ | -------- |
| draft_id | body | 是   | string | 文章正文 |

### 返回参数

无

### 返回示例





## 更新文章草稿内容

### 请求路径

```http
POST /api/content/article_draft/update
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称          | 位置 | 必选 | 类型   | 说明       |
| ------------- | ---- | ---- | ------ | ---------- |
| content       | body | 是   | string | 文章正文   |
| brief_content | body | 是   | string | 文章摘要   |
| title         | body | 是   | string | 文章标题   |
| tag_ids       | body | 是   | array  | 标签       |
| category_id   | body | 是   | string | 分类       |
| id            | body | 是   | string | 文章草稿id |

### 返回参数

无

### 返回示例

**正确示例**

```json
{
"status":200
"message":"update content successfully"
}
```

**错误示例**

```json
//找不到文章草稿
{
    "status":400
    "messgae":"draft id is invalid"
}
```

## 获取文章草稿内容

### 请求路径

```http
GET /api/content/article_draft/detail
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称     | 位置  | 必选 | 类型   | 说明       |
| -------- | ----- | ---- | ------ | ---------- |
| draft_id | query | 是   | string | 文章草稿id |

### 返回参数

| 名称          | 必选 | 类型   | 说明       |
| ------------- | ---- | ------ | ---------- |
| content       | 是   | string | 文章正文   |
| brief_content | 是   | string | 文章摘要   |
| title         | 是   | string | 文章标题   |
| tag_ids       | 是   | array  | 标签       |
| category_id   | 是   | string | 分类       |
| id            | 是   | string | 文章草稿id |
| user_id       | 是   | string | 作者id     |

### 返回示例

**正确返回**

```json
{
 "status":200
 "message":"get draft detail successfully"
 "data":{
  "content":string
  "brief_content":string
  "title":string
  "tag_ids":array
  "catagory_id":number
  "id":string
  "user_id":string
 }
}
```

**错误返回**

```json
//找不到draft id对应的草稿
{
"status":400
"message":"draft id is inalid"
}
```

## 发布文章

### 请求路径

```http
POST /api/content/article/publish
application/json
```

### **header**

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称     | 位置 | 必选 | 类型   | 说明       |
| -------- | ---- | ---- | ------ | ---------- |
| draft_id | body | 是   | string | 文章草稿id |

### 返回参数

| 名称       | 必选 | 类型   | 说明   |
| ---------- | ---- | ------ | ------ |
| article_id | 是   | string | 文章id |
| user_id    | 是   | string | 用户id |

### 返回示例

**正确返回**

```json
{
"status":200
"message":"publish post successfully"
"data":{
     "article_id":string
     "user_id":string
}
}
```

**错误返回**

```json
//找不到文章草稿
{
"status":400
"message":"draft id invalid"
}
```

# 收藏

## 新建收藏夹

### 请求路径

```http
POST /api/collectionset
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称            | 位置 | 必选 | 类型   | 说明                   |
| --------------- | ---- | ---- | ------ | ---------------------- |
| collection_name | body | 是   | string | 收藏夹名称             |
| description     | body | 否   | string | 简介                   |
| permission      | body | 是   | number | 公开性：0-公开、1-私密 |

### 返回参数

| 名称               | 必选 | 类型       | 说明         |
| ------------------ | ---- | ---------- | ------------ |
| collectionset_info | 是   | 复杂结构体 | 收藏夹的信息 |

### 返回示例

**正确返回**

```json
{  
  "status":200
  "message":"create collection set succefully"
  "data":{
    "collection_id":string       //收藏夹id
    "collection_name":string      
    "user_id":string
    "permission":number           //访问权限 1-私密 0-公开
    "post_article_count":number   
    "create_time":timestamp     
    "update_time":timestamp
  }
}
```

## 收藏文章

### 请求路径

```http
POST /api/collection
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称                 | 位置 | 必选 | 类型   | 说明         |
| -------------------- | ---- | ---- | ------ | ------------ |
| article_id           | body | 是   | string | 文章id       |
| select_collection_id | body | 是   | string | 选中的收藏夹 |

### 返回参数

无

### 返回示例

```json
{
 "status":200
 "message":"add colection successfully"
}
```

## 获取收藏夹详情

### 请求路径

```http
GET /api/collection/detail
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称          | 位置  | 必选 | 类型   | 说明         |
| ------------- | ----- | ---- | ------ | ------------ |
| collection_id | query | 是   | string | 收藏夹id     |
| limit         | query | 是   | string | 显示最大数量 |

### 返回参数

| 名称            | 必选 | 类型       | 说明             |
| --------------- | ---- | ---------- | ---------------- |
| articles        | 是   | array      | 收藏的文章的信息 |
| collection_info | 是   | 复杂结构体 | 收藏夹的信息     |

### 返回示例

```json
{
   "status":200
   "msg":"get collection detail successfully"
   "data": {
      "articles":[
      0:{
          "article_id":string
          "article_major":{
                "title":string
                "brief_content":string
                "cover_img":string
                ""
          }
          "article_counter":{
                 "view_count":number
                 "collect_count":number
                 "comment_count":number
                 "create_time":timestamp
          }
          "tag_ids":array
          "author_user_info":{user_info}
      }
      ]
    "collection_info":{
        "collection_id":string
        "collection_name":string
        "user_id":string
        "description":string
        "permission":number
        "post_article_count":number
        "update_time":timestamp
    }
   }
}
```

# 个人详情信息

## 获取用户个人信息

### 请求路径

```http
GET /api/user/info
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

无

### 返回参数

| 名称         | 必选 | 类型       | 说明         |
| ------------ | ---- | ---------- | ------------ |
| user_id      | 是   | string     | 用户id       |
| username     | 是   | string     | 用户名       |
| create_time  | 是   | timestamp  | 创建时间     |
| user_basic   | 是   | 复杂结构体 | 用户基本信息 |
| user_counter | 是   | 复杂结构体 | 用户数据信息 |

### 返回示例

**正确返回**

```json
{
"status":200
"info":"get user info successfully"
"data":
    "user_id":string
    "username":string
    "create_time":timestamp
    "user_basic":{
       "avatar":string
       "company":string
       "description":string
       "job_title":string
    }
    "user_counter":{
        "digg_article_count":number
        "digg_shortmsg_count":number
        "followee_count":number
        "follower_count":number
        "got_digg_count":number
        "got_view_count":number
        "post_article_count":number
        "post_shortmsg_count":number
        "select_online_course_count":number    //选中的课
    }
}
```

**错误返回**

无

获取用户个人发布的文章

## 编辑个人资料

### 请求路径

```http
POST /api/user/update
mutipart/form-data
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称         | 位置 | 必选 | 类型   | 说明        |
| ------------ | ---- | ---- | ------ | ----------- |
| username     | body | 否   | string | 用户名      |
| job_title    | body | 否   | string | 职位        |
| company      | body | 否   | string | 公司        |
| blog_address | body | 否   | string | 个人主页url |
| description  | body | 否   | string | 个人介绍    |
| avatar       | body | 否   | file   | 头像文件    |

### 返回参数

无

### 返回示例

**正确返回**

```json
{
"status":200
"message":"update user info successfully"
}
```

# 课程

## 获取掘金小册列表

### 请求路径

```http
POST /api/course/bokklet/list_by_category
aplication/json
```

### header

无

### 请求参数

| 名称      | 位置 | 必选 | 类型   | 说明     |
| --------- | ---- | ---- | ------ | -------- |
| category  | body | 是   | number | 分类     |
| page_size | body | 是   | number | 列表大小 |
| sort      | body | 是   | number | 排列规则 |

### 返回参数

| 名称       | 必选 | 类型       | 说明                                           |
| ---------- | ---- | ---------- | ---------------------------------------------- |
| booklet_id | 是   | string     | 掘金小册id                                     |
| base_info  | 是   | 复杂结构体 | 基本信息：标题、概括、作者……                   |
| discount   | 是   | 复杂结构体 | 折扣信息：开始时间、结束时间、折扣、折扣名字…… |
| user_info  | 是   | 复杂结构体 | 作者信息                                       |

### 返回案例

**正确返回**

```json
{
 "status":200
 "message":"get booklet list successfully"
 "data":[{"booklet_id":string
    "base_info":{
      "cover_img":string //封面的url
      "title":string
      "summary":string
      "is_finished":number
      "buy_count":number
      "section_count":number
      "category_id":number
    }    
    "discount":{
        "name":string
        "desc":"{desc}"
        "start_time":time
        "end_time":time
        "price":number
        "discount_rate":number
    }
   "user_info":{
       "user_id":"{user_id}"
       "avatar":string //头像的url
       "company":string
       "job_tile":string
    }
   }
 ]
}
```

**错误返回**

无

## 获取掘金小册详情

### 请求路径

```http
GET /api/course/booklet_id/detail{booklet_id}
```

### header

无

### 请求参数

| 名称       | 位置  | 必选 | 类型   | 说明       |
| ---------- | ----- | ---- | ------ | ---------- |
| booklet_id | query | 是   | string | 掘金小册id |

### 返回参数

| 名称         | 必选 | 类型       | 说明         |
| ------------ | ---- | ---------- | ------------ |
| booklet_id   | 是   | string     | 掘金小册id   |
| base_info    | 是   | 复杂结构体 | 基本信息     |
| discount     | 是   | 复杂结构体 | 折扣信息     |
| user_info    | 是   | 复杂结构体 | 用户信息     |
| introduction | 是   | 复杂结构体 | 掘金小册介绍 |

### 返回示例

**正确返回**

```json
{
"status":200
"message":"get bytecourse detail successfully"
"data":{
  "booklet_id":string
  "base_info":{
      "cover_img":string //封面的url
      "title":string
      "summary":string
      "is_finished":number
      "buy_count":number
      "section_count":number
      "category_id":number
  }    
  "discount":{
        "name":string 
        "desc":"{desc}"
        "start_time":time
        "end_time":time
        "price":number
        "discount_rate":number
  }
  "user_info":{
       "user_id":string
       "avatar":string //头像的url
       "company":string
       "job_tile":string //职位
  }
  "introduction":{
        "content":"{content}(<h2>作者介绍</h2>\……)"
        ""
  }
 }
}
```

**错误返回**

无

## 获取字节内部课列表

### 请求路径

```http
GET /api/course/bytecourse//list_by_category{category_id}{page_size}
```

### header

无

### 请求参数

| 名称      | 位置  | 必选 | 类型   | 说明     |
| --------- | ----- | ---- | ------ | -------- |
| category  | query | 是   | number | 分类     |
| page_size | query | 是   | number | 列表大小 |

### 返回参数

| 名称      | 必选 | 类型       | 说明                   |
| --------- | ---- | ---------- | ---------------------- |
| course_id | 是   | string     | 字节内部课id           |
| base_info | 是   | 复杂结构体 | 基本信息：标题、概括…… |

### 返回示例

**正确返回**

```json
{
"status":200
"message":"get bytecourse list successfully"
"data":[ "course_id":"{course_id}"
  "base_info"：{
    "name":string
    "summary":string
    "cover_image":string  //封面的url
    "videos_count":number
    "course_time":number
  }
 ]
}
```

**错误返回**

无

## 获取字节内部课详情

### 请求路径

```http
GET /api/course/bytecourse/detail{course_id}
```

### header

无

### 请求参数

| 名称      | 位置  | 必选 | 类型   | 说明         |
| --------- | ----- | ---- | ------ | ------------ |
| course_id | query | 是   | string | 字节内部课id |

### 返回参数

| 名称    | 必选 | 类型       | 说明     |
| ------- | ---- | ---------- | -------- |
| content | 是   | 复杂结构体 | 课程详情 |

### 返回示例

**正确返回**

```json
{
 "status":200
 "message":"get bytecourse detail successfully"
 "data":{
 "content":{
   "name:
   "course_id":number
   "cover_image":url
   "summary":string
   "content":{content(## 课程介绍\n在云原生时代浪潮下，Se……)}
  }
 }
}
```

**错误返回**

无

# 支付（待定）

### 请求路径

### header

### 请求参数

### 返回参数

### 返回示例

# 消息

## 获取点赞/关注/评论/系统消息

### 请求路径

```http
GET /api/message
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称         | 位置  | 必选 | 类型   | 说明                                       |
| ------------ | ----- | ---- | ------ | ------------------------------------------ |
| limit        | query | 是   | number | 获取私信最大条数                           |
| message_type | query | 是   | number | 信息类型： 1-点赞 2-关注 3-评论 4-系统消息 |

### 返回参数

| 名称        | 必选 | 类型       | 说明             |
| ----------- | ---- | ---------- | ---------------- |
| count       | 是   | number     | 查询到的消息条数 |
| message     | 是   | 复杂结构体 | 消息的信息       |
| dst_info    | 是   | 复杂结构体 | 目标消息的信息   |
| parent_info | 否   | 复杂结构体 | 上一级信息       |
| src_info    | 是   | 复杂结构体 | 信息源信息       |

### 返回示例

**正确示例**

```json
{
"status":200
"message":"get message successfully"
 "data":[
    {"dst_info":{      //获取到的目标信息
      "detail":string    //内容
      "id_type":number    //id的类型 1-用户 2-文章 5-评论
      "item_id":number
      "is_digg":boolean   //是否被赞
      "name":string
}
    "message":{        
      "message_id":string
      "message_type":number
      "create_time":timestamp   
      "dst_id":string
      "dst_type":number
      "owner_id":string   //信息拥有者id
      "src_id":number     //信息发起者id
      "src_type":number
}  
   "parent_info":{     //上一级信息
      "detail":string    
      "id_type":number
      "item_id":string
       "is_digg":bool
      "name":string    
   }
   "src_info":{
       "detail":string
        "id_type":number
        "item_id":string
        "name":string       
   }
    
}]
}
```

**错误示例**

无

# 评论

## 发表评论

### 请求路径

```http
POST /api/comment
aplication/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称            | 位置 | 必选 | 类型   | 说明     |
| --------------- | ---- | ---- | ------ | -------- |
| comment_content | body | 是   | string | 评论内容 |
| item_id         | body | 是   | string | 项目id   |
| id_type         | body | 是   | number | id类型:2 |

### 返回参数

| 名称         | 必选 | 类型       | 说明       |
| ------------ | ---- | ---------- | ---------- |
| comment_id   | 是   | string     | 评论的id   |
| comment_info | 是   | 复杂结构体 | 评论的信息 |

### 返回示例

**正确返回**

```json
{
 "status":200
 "message":"comment successfully"
 "comment_id":string
 "comment_info":{
      "comment_id":string
      "comment_content":string
      "comment_replys":array
      "ctime":timestamp
      "is_digg":boolean
      "digg_count":number
      "item_id":string
      "item_type":number
      "user_id":string
      "replys":array
 }
}
```

## 删除评论

### 请求路径

```http
DELETE /api/comment
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称       | 位置 | 必选 | 类型   | 说明   |
| ---------- | ---- | ---- | ------ | ------ |
| comment_id | body | 是   | string | 评论id |

### 返回参数

无

### 返回示例

```json
 {
 "status":200 
 "message":"delete comment successfully"
 }
```

## 回复评论/回复

### 请求路径

```http
POST  /api/reply
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称                   | 位置 | 必选 | 类型   | 说明                 |
| ---------------------- | ---- | ---- | ------ | -------------------- |
| reply_content          | body | 是   | string | 回复内容             |
| reply_to_comment_id    | body | 是   | string | 评论id               |
| reply_to_reply_id      | body | 是   | string | 回复id               |
| reply_to_reply_user_id | body | 是   | string | 回复用户id           |
| item_id                | body | 是   | string | 项目id(这里是文章id) |
| id_type                | body | 是   | number | id类型:2             |

### 返回参数

| 名称         | 必选 | 类型                            | 说明       |
| ------------ | ---- | ------------------------------- | ---------- |
| reply_id     | 是   | string                          | 回复的id   |
| reply_info   | 是   | 复杂结构体                      | 回复的信息 |
| parent_reply | 是   | 复杂结构体 结构与reply_info相同 | 上一级回复 |

### 返回示例

**正确返回**

```json
{
 "status":200
 "message":"reply successfully"
  "reply_id":string
  "parent_reply":{
      "reply_id":string
      "reply_content":string
      "ctime":timestamp
      "is_digg":boolean
      "digg_count":number
      "item_id":string      //这里是文章id
      "item_type":number
      "reply_comment_id":string   //所属评论的id
      "reply_to_user_id":string   //这里是2
      "user_id":string 
}
  "reply_info":{
      "reply_id":string
      "reply_content":string
      "ctime":timestamp
      "is_digg":boolean
      "digg_count":number
      "item_id":string      //这里是文章id
      "item_type":number     //这里是2
      "reply_comment_id":string
      "reply_to_user_id":string
      "user_id":string     
 }
}
```

## 删除回复

### 请求路径

```http
DELETE /api/reply
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称       | 位置 | 必选 | 类型   | 说明   |
| ---------- | ---- | ---- | ------ | ------ |
| reply_id   | body | 是   | string | 回复id |
| comment_id | body | 是   | string | 评论id |

### 返回参数

无

### 返回示例

**正确返回**

```json
{
 "status":200
 "message":"delete reply successfully"
}
```



# 点赞

## 点赞

### 请求路径

```http
POST /api/digg
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称    | 位置 | 必选 | 类型   | 说明                        |
| ------- | ---- | ---- | ------ | --------------------------- |
| item_id | body | 是   | string | 项目id                      |
| id_type | body | 是   | number | id类型:2-文章 5-评论 6-回复 |

### 返回参数

无

### 返回示例

**正确返回**

```json
{
   "status":200
   "message":"digg successfully"
}
```

## 取消点赞

### 请求路径

```
DELETE /api/digg
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称    | 位置 | 必选 | 类型   | 说明                   |
| ------- | ---- | ---- | ------ | ---------------------- |
| item_id | body | 是   | string | 项目id                 |
| id_type | body | 是   | number | 项目类型:2-文章 5-评论 |

### 返回参数

无

### 返回示例

**正确返回**

```json
{
   "status":200
   "message":"undo digg successfully"
}
```

# 私信(待定)

## 发私信

### 请求路径

```http
POST /api/chat/message
applcation/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称            | 位置 | 必选 | 类型   | 说明       |
| --------------- | ---- | ---- | ------ | ---------- |
| to_user_id      | body | 是   | string | 对方的id   |
| message_content | body | 是   | string | 私信内容   |
| item_id         | body | 是   | string | 项目id     |
| item_type       | body | 是   | number | 项目类型:3 |

### 返回参数

| 名称         | 必选 | 类型       | 说明       |
| ------------ | ---- | ---------- | ---------- |
| message_info | 是   | 复杂结构体 | 消息的信息 |

### 返回示例

**正确返回**

```json
{
"status":200
"message":"sent message successfully"
 "message_info":{        
     "message_id":string
     "create_time":timestamp   
     "src_id":number     //信息发起者id
     "src_type":number   //这里是1
     "item_id":string
     "id_type":number     //这里是3
    }  
}
```

## 获取聊天信息

### 请求路径

```http
GET /api/chat
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

无

### 返回参数

### 返回示例

# 关注

## 添加关注

### 请求路径

```http
POST /api/follow
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称 | 位置 | 必选 | 类型   | 说明         |
| ---- | ---- | ---- | ------ | ------------ |
| id   | body | 是   | string | 关注用户的id |

### 返回参数

无

### 返回示例

**正确返回**

```json
{
"status":200
 "message":"do follow successfully"
}
```

## 取消关注

### 请求路径

```http
DELETE /api/follow
application/json
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称 | 位置 | 必选 | 类型   | 说明               |
| ---- | ---- | ---- | ------ | ------------------ |
| id   | body | 是   | string | 要取消关注用户的id |

### 返回参数

无

### 返回示例

**正确返回**

```json
{
"status":200
 "message":"undo follow successfully"
}
```

# 内容管理

## 获取文章/草稿/专栏/沸点数目

### 请求路径

```http
GET /api/author_center/count
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称     | 位置  | 必选 | 类型  | 说明                       |
| -------- | ----- | ---- | ----- | -------------------------- |
| id_types | query | 是   | array | 要获取的内容的id类型的数组 |

### 返回参数

| 名称          | 必选 | 类型       | 说明     |
| ------------- | ---- | ---------- | -------- |
| article_cnt   | 否   | 复杂结构体 | 文章数目 |
| draft_cnt     | 否   | 复杂结构体 | 草稿数目 |
| column_cnt    | 否   | 复杂结构体 | 专栏数目 |
| short_msg_cnt | 否   | 复杂结构体 | 沸点数目 |

### 返回示例

```json
{
"status":200
"message":"get count successfully"
"data":{
   "article_cnt":{
     "all_cnt":0         //全部
     "auditing_cnt":0    //审核中
     "audit_pass_cnt":0   //已发布
     "audit_fail_pass":0  //未通过
   }
   "column_cnt":{
     "all_cnt":0         //全部
     "auditing_cnt":0    //审核中
     "audit_pass_cnt":0   //已发布
     "audit_fail_pass":0  //未通过
   }
}
}
```

## 获取文章/草稿/专栏/沸点列表

### 请求路径

```http
GET /api/author_centor/{type}/list{page_no}{page_size}
```

### header

| 名称          | 必选 | 数值          | 说明              |
| ------------- | ---- | ------------- | ----------------- |
| Authorization | 是   | Bearer $token | 用户身份验证token |

### 请求参数

| 名称      | 位置  | 必选 | 类型   | 说明                                                         |
| --------- | ----- | ---- | ------ | ------------------------------------------------------------ |
| type      | url   | 是   | string | 要获取列表的类型：article-文章 column-专栏 draft-草稿 short-msg-沸点 |
| page_no   | query | 是   | number | 页码                                                         |
| page_size | query | 是   | number | 页面显示最大条数                                             |

### 返回参数

| 名称           | 必选 | 类型       | 说明     |
| -------------- | ---- | ---------- | -------- |
| count          | 是   | number     | 数目     |
| article_info   | 否   | 复杂结构体 | 文章信息 |
| column_info    | 否   | 复杂结构体 | 专栏信息 |
| draft_info     | 否   | 复杂结构体 | 草稿信息 |
| short_meg-info | 否   | 复杂结构体 | 沸点信息 |

### 返回示例

**正确返回**

```json
{
 "status":200
 "message":"get list succesfully"
 "data":[
   0:"draft_info":{
   "article_id":string
   "summary":string
   "catagory_id":string
   "cover_img":string
   "creat_time":timestamp
   "update_time":timestamp
   "title":string
   "tag_ids":array        //标签的id
   "user_id"       //用户id   
   }
 ]
}
```

```json
{
 "status":200
 "message":"get list succesfully"
 "data":[
   0:"article_info":{
   "article_id":string
   "audit_status":number    //状态: 1-审核中 2-已发布 3-未通过
   "summary":string
   "catagory_id":string
   "cover_img":string
   "display_cnt":number
   "digg_cnt":number
   "comment_cnt":number
   "collect_cnt":number 
   "creat_time":timestamp
   "update_time":timestamp
   "title":string
   "tag_ids":array        //标签的id
   "user_id"       //用户id   
   }
 ]
}
```

# 数据中心(待定)

## 获取卡片上的数据

### 请求路径

```

```

### header

### 请求参数

### 返回参数

### 返回示例

## 获取趋势图的数据

### 请求路径

```

```

### header

### 请求参数

### 返回参数

### 返回示例

# 文本id类型对照表

| 编号 |   含义   |
| :--: | :------: |
|  1   |   用户   |
|  2   |   文章   |
|  3   |   专栏   |
|  4   |   草稿   |
|  5   |   评论   |
|  6   |   回复   |
|  24  |   沸点   |
| 1001 | 系统消息 |

# 消息类型对照表

| 编号 |   含义   |
| :--: | :------: |
|  1   |   点赞   |
|  2   |   关注   |
|  3   |   评论   |
|  4   | 系统消息 |

















