# Gin+html+css仿稀土掘金


## 第三方包

**雪花算法**：`github.com/bwmarrin/snowflake`

**jwt**:`github.com/golang-jwt/jwt/v4`

**uuid**:`github.com/google/uuid`

**配置管理**：`github.com/spf13/viper`

**日志管理**:   `go.uber.org/zap`

**邮件**:  `gopkg.in/gomail.v2`

**七牛云对象存储**: `github.com/qiniu/go-sdk/v7`

**令牌桶**: `github.com/juju/ratelimit`

**cron**: `github.com/robfig/cron/v3`

## 中间件

### 全局

1. cors
2. 日志输出
3. 令牌桶
4. 捕获异常记录日志

### 局部

1. jwt验证

## 实现功能 :tada:

1. 草稿（创建、更新、删除）:scroll:
2. 文章（发布文章、获取文章、根据分类和排序方式获取文章列表）:bookmark_tabs:
3. 用户（注册、登录、验证邮箱）:bust_in_silhouette:
4. 个人信息（获取、修改、上传个人头像）:information_source:
5. 点赞 :+1:
6. 关注 :eyes:
7. 收藏（创建收藏夹、获取收藏夹信息、收藏文章、取消收藏） :heart:
8. 评论（发表评论，删除评论）:right_anger_bubble:
9. 回复（回复评论、删除回复）
10. 作者榜  :triangular_flag_on_post:

## 表结构

![](https://github.com/GDshenqingNo1/juejin/blob/master/db.jpg)

## **缓存设计**

### jwt黑名单存放主动过期的token

过期时间设置为原来token的过期时间

```
key blacklist:$token val ""
```

![cache1](https://github.com/GDshenqingNo1/juejin/blob/master/cache/cache1.png)

### 点赞状态

过期时间设置为两天，后台任务定时将缓存同步到mysql中，不删除缓存

status 为 1 时 已点赞

status 为 0 时 未点赞 

```
key $userid:$item_id:$item_type val {$status}
```

![cache2](https://github.com/GDshenqingNo1/juejin/blob/master/cache/cache2.png)

### 点赞数、浏览数计数

hash保存点赞数和浏览数，定时任务将缓存的数据同步到mysql对应的表后删除缓存

**用户**

```
key user_counter field {$user_id:digg_article_count/got_digg_count/got_view_count} value {$count}
```

![cache4](https://github.com/GDshenqingNo1/juejin/blob/master/cache/cache4.png)

**文章**

```
key article_counter field {$article_id:digg_count/view_count} value {$count}
```

![cache5](https://github.com/GDshenqingNo1/juejin/blob/master/cache/cache5.png)

**评论**

```
key comment_counter field {$comment_id:digg_count} value {$count}
```

![cache6](https://github.com/GDshenqingNo1/juejin/blob/master/cache/cache6.png)

**回复**

```
key reply_counter field {$comment_id:digg_count} value {$count}
```

![cache7](https://github.com/GDshenqingNo1/juejin/blob/master/cache/cache7.png)








