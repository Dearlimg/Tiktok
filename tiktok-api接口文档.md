---
title: 个人项目
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# 个人项目

Base URLs:

# Authentication

# Default

## GET 获取视频流

GET /tiktok/feed

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|latest_time|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "刷新视频流成功!",
  "video_list": [
    {
      "id": 2,
      "AuthorId": 1,
      "title": "第一个测试视频",
      "play_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok559062a5-a654-4e59-9b50-51f4ecc7db37.mp4",
      "cover_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok559062a5-a654-4e59-9b50-51f4ecc7db37.mp4?x-oss-process=video/snapshot,t_2000,m_fast",
      "CreatedAt": "2025-07-15T19:38:51+08:00",
      "UpdatedAt": "2025-07-15T19:38:51+08:00",
      "author": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 2,
        "favorite_count": 1
      },
      "favorite_count": 1,
      "comment_count": 0,
      "is_favorite": false
    },
    {
      "id": 1,
      "AuthorId": 1,
      "title": "第一个测试视频",
      "play_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok4fe82021-1ce1-46b3-b724-41732be15de6.mp4",
      "cover_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok4fe82021-1ce1-46b3-b724-41732be15de6.mp4?x-oss-process=video/snapshot,t_2000,m_fast",
      "CreatedAt": "2025-07-15T19:35:34+08:00",
      "UpdatedAt": "2025-07-15T19:35:34+08:00",
      "author": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 2,
        "favorite_count": 1
      },
      "favorite_count": 0,
      "comment_count": 7,
      "is_favorite": false
    }
  ],
  "next_time": 1752579334
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» video_list|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» AuthorId|integer|true|none||none|
|»» title|string|true|none||none|
|»» play_url|string|true|none||none|
|»» cover_url|string|true|none||none|
|»» CreatedAt|string|true|none||none|
|»» UpdatedAt|string|true|none||none|
|»» author|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» follow_count|integer|true|none||none|
|»»» follower_count|integer|true|none||none|
|»»» is_follow|boolean|true|none||none|
|»»» avatar|string|true|none||none|
|»»» background_image|string|true|none||none|
|»»» signature|string|true|none||none|
|»»» total_favorited|integer|true|none||none|
|»»» work_count|integer|true|none||none|
|»»» favorite_count|integer|true|none||none|
|»» favorite_count|integer|true|none||none|
|»» comment_count|integer|true|none||none|
|»» is_favorite|boolean|true|none||none|
|» next_time|integer|true|none||none|

## POST 用户注册

POST /tiktok/user/register

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|username|query|string| 是 |none|
|password|query|string| 是 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "Register Success",
  "user_id": 3,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MywiVXNlcm5hbWUiOiJkdXJsaW0xIiwiZXhwIjoxNzUyNzM2MzQxLCJpc3MiOiJ3d3cuZHVybGltLnh5eiJ9.OyfR7F0lCUqTyibE283DdQA8qre5aJe3L81ajMEpRCA"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» user_id|integer|true|none||none|
|» token|string|true|none||none|

## POST 用户登录

POST /tiktok/user/login

> Body 请求参数

```yaml
application/octet-stram: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcm5hbWUiOiJnYW9qaWF5dWFuIiwiZXhwIjoxNzUyNjUzODYxLCJpc3MiOiJ3d3cuZHVybGltLnh5eiJ9.NCxJGcnIJ0KepusX9gLx5ceaUkG5OXemkDBP_z_gsvU

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|username|query|string| 是 |none|
|password|query|string| 是 |none|
|body|body|object| 否 |none|
|» application|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "Login Success",
  "user_id": 2,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwiVXNlcm5hbWUiOiJkdXJsaW0iLCJleHAiOjE3NTI3MzYwMDcsImlzcyI6Ind3dy5kdXJsaW0ueHl6In0.8chstC95K4ln7ouDTibPxHb6pmVckuCv2hn-ZXKDilY"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» user_id|integer|true|none||none|
|» token|string|true|none||none|

## GET 用户信息

GET /tiktok/user

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|token|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "Query Success",
  "user": {
    "id": 1,
    "name": "gaojiayuan",
    "follow_count": 1,
    "follower_count": 1,
    "is_follow": false,
    "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
    "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
    "signature": "我想休息",
    "total_favorited": 1,
    "work_count": 2,
    "favorite_count": 1
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» user|object|true|none||none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|»» follow_count|integer|true|none||none|
|»» follower_count|integer|true|none||none|
|»» is_follow|boolean|true|none||none|
|»» avatar|string|true|none||none|
|»» background_image|string|true|none||none|
|»» signature|string|true|none||none|
|»» total_favorited|integer|true|none||none|
|»» work_count|integer|true|none||none|
|»» favorite_count|integer|true|none||none|

## POST 发布视频

POST /tiktok/publish/action

> Body 请求参数

```yaml
data: string
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcm5hbWUiOiJnYW9qaWF5dWFuIiwiZXhwIjoxNzUyNjUzOTc4LCJpc3MiOiJ3d3cuZHVybGltLnh5eiJ9.zzO5hGPt5Xt_ekvQC_LSW8rm_F-NO1Cg9w3mh_aWKF0
title: 第一个测试视频

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|title|query|string| 是 |none|
|body|body|object| 否 |none|
|» data|body|string(binary)| 是 |none|
|» token|body|string| 是 |none|
|» title|body|string| 是 |none|

> 返回示例

> 200 Response

```
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 发布视频列表

GET /tiktok/publish/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|user_id|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "获取用户发布的视频列表成功！",
  "video_list": [
    {
      "id": 1,
      "AuthorId": 1,
      "title": "第一个测试视频",
      "play_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok4fe82021-1ce1-46b3-b724-41732be15de6.mp4",
      "cover_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok4fe82021-1ce1-46b3-b724-41732be15de6.mp4?x-oss-process=video/snapshot,t_2000,m_fast",
      "CreatedAt": "2025-07-15T19:35:34+08:00",
      "UpdatedAt": "2025-07-15T19:35:34+08:00",
      "author": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "favorite_count": 0,
      "comment_count": 7,
      "is_favorite": false
    },
    {
      "id": 2,
      "AuthorId": 1,
      "title": "第一个测试视频",
      "play_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok559062a5-a654-4e59-9b50-51f4ecc7db37.mp4",
      "cover_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok559062a5-a654-4e59-9b50-51f4ecc7db37.mp4?x-oss-process=video/snapshot,t_2000,m_fast",
      "CreatedAt": "2025-07-15T19:38:51+08:00",
      "UpdatedAt": "2025-07-15T19:38:51+08:00",
      "author": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "favorite_count": 1,
      "comment_count": 0,
      "is_favorite": true
    },
    {
      "id": 3,
      "AuthorId": 1,
      "title": "第一个测试视频",
      "play_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok77c0fc57-3786-46f2-aab7-ca50a9d2d723.mp4",
      "cover_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok77c0fc57-3786-46f2-aab7-ca50a9d2d723.mp4?x-oss-process=video/snapshot,t_2000,m_fast",
      "CreatedAt": "2025-07-16T15:14:17+08:00",
      "UpdatedAt": "2025-07-16T15:14:17+08:00",
      "author": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "favorite_count": 0,
      "comment_count": 0,
      "is_favorite": false
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» video_list|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» AuthorId|integer|true|none||none|
|»» title|string|true|none||none|
|»» play_url|string|true|none||none|
|»» cover_url|string|true|none||none|
|»» CreatedAt|string|true|none||none|
|»» UpdatedAt|string|true|none||none|
|»» author|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» follow_count|integer|true|none||none|
|»»» follower_count|integer|true|none||none|
|»»» is_follow|boolean|true|none||none|
|»»» avatar|string|true|none||none|
|»»» background_image|string|true|none||none|
|»»» signature|string|true|none||none|
|»»» total_favorited|integer|true|none||none|
|»»» work_count|integer|true|none||none|
|»»» favorite_count|integer|true|none||none|
|»» favorite_count|integer|true|none||none|
|»» comment_count|integer|true|none||none|
|»» is_favorite|boolean|true|none||none|

## POST 取消点赞

POST /tiktok/favorite/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|video_id|query|string| 是 |none|
|action_type|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "favourite action success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|

## GET 获取点赞列表

GET /tiktok/favorite/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|user_id|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "get favouriteList success",
  "video_list": [
    {
      "id": 2,
      "AuthorId": 1,
      "title": "第一个测试视频",
      "play_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok559062a5-a654-4e59-9b50-51f4ecc7db37.mp4",
      "cover_url": "durlim-tiktok.oss-cn-beijing.aliyuncs.comtiktok559062a5-a654-4e59-9b50-51f4ecc7db37.mp4?x-oss-process=video/snapshot,t_2000,m_fast",
      "CreatedAt": "2025-07-15T19:38:51+08:00",
      "UpdatedAt": "2025-07-15T19:38:51+08:00",
      "author": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "favorite_count": 1,
      "comment_count": 0,
      "is_favorite": true
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» video_list|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» AuthorId|integer|false|none||none|
|»» title|string|false|none||none|
|»» play_url|string|false|none||none|
|»» cover_url|string|false|none||none|
|»» CreatedAt|string|false|none||none|
|»» UpdatedAt|string|false|none||none|
|»» author|object|false|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» follow_count|integer|true|none||none|
|»»» follower_count|integer|true|none||none|
|»»» is_follow|boolean|true|none||none|
|»»» avatar|string|true|none||none|
|»»» background_image|string|true|none||none|
|»»» signature|string|true|none||none|
|»»» total_favorited|integer|true|none||none|
|»»» work_count|integer|true|none||none|
|»»» favorite_count|integer|true|none||none|
|»» favorite_count|integer|false|none||none|
|»» comment_count|integer|false|none||none|
|»» is_favorite|boolean|false|none||none|

## POST 删除评论

POST /tiktok/comment/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|video_id|query|string| 是 |none|
|action_type|query|string| 是 |none|
|comment_id|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "delete commentId success",
  "comment": {
    "id": 0,
    "user": {
      "id": 0,
      "name": "",
      "follow_count": 0,
      "follower_count": 0,
      "is_follow": false,
      "avatar": "",
      "background_image": "",
      "signature": "",
      "total_favorited": 0,
      "work_count": 0,
      "favorite_count": 0
    },
    "content": "",
    "create_date": "",
    "like_count": 0,
    "tease_count": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» comment|object|true|none||none|
|»» id|integer|true|none||none|
|»» user|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» follow_count|integer|true|none||none|
|»»» follower_count|integer|true|none||none|
|»»» is_follow|boolean|true|none||none|
|»»» avatar|string|true|none||none|
|»»» background_image|string|true|none||none|
|»»» signature|string|true|none||none|
|»»» total_favorited|integer|true|none||none|
|»»» work_count|integer|true|none||none|
|»»» favorite_count|integer|true|none||none|
|»» content|string|true|none||none|
|»» create_date|string|true|none||none|
|»» like_count|integer|true|none||none|
|»» tease_count|integer|true|none||none|

## GET 获取评论列表

GET /tiktok/comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|video_id|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "comment_list": [
    {
      "id": 8,
      "user": {
        "id": 2,
        "name": "durlim",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 0,
        "work_count": 0,
        "favorite_count": 0
      },
      "content": "啊啊啊啊啊啊啊宝宝你是个宝宝,啊啊啊啊啊啊啊啊啊啊啊啊啊啊`!!!!!!啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
      "create_date": "2025-07-16 09:38:24",
      "like_count": 71,
      "tease_count": 19
    },
    {
      "id": 7,
      "user": {
        "id": 2,
        "name": "durlim",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 0,
        "work_count": 0,
        "favorite_count": 0
      },
      "content": "啊啊啊啊啊啊啊宝宝你是个宝宝,啊啊啊啊啊啊啊啊啊啊啊啊啊啊`!!!!!!",
      "create_date": "2025-07-16 09:31:52",
      "like_count": 63,
      "tease_count": 49
    },
    {
      "id": 6,
      "user": {
        "id": 2,
        "name": "durlim",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 0,
        "work_count": 0,
        "favorite_count": 0
      },
      "content": "啊啊啊啊啊啊啊宝宝你是个宝宝,啊啊啊啊啊啊啊啊啊啊啊啊啊啊`!!!!!!",
      "create_date": "2025-07-16 09:24:07",
      "like_count": 86,
      "tease_count": 36
    },
    {
      "id": 5,
      "user": {
        "id": 2,
        "name": "durlim",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 0,
        "work_count": 0,
        "favorite_count": 0
      },
      "content": "啊啊啊啊啊啊啊宝宝你是个宝宝,啊啊啊啊啊啊啊啊啊啊啊啊啊啊`",
      "create_date": "2025-07-16 09:06:45",
      "like_count": 53,
      "tease_count": 14
    },
    {
      "id": 4,
      "user": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "content": "啊啊啊啊啊啊啊宝宝你是个宝宝",
      "create_date": "2025-07-15 21:04:49",
      "like_count": 18,
      "tease_count": 5
    },
    {
      "id": 3,
      "user": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "content": "今天碰到了一个下头男",
      "create_date": "2025-07-15 21:04:29",
      "like_count": 59,
      "tease_count": 7
    },
    {
      "id": 2,
      "user": {
        "id": 1,
        "name": "gaojiayuan",
        "follow_count": 1,
        "follower_count": 1,
        "is_follow": false,
        "avatar": "durlim-tiktok.oss-cn-beijing.aliyuncs.comhttps://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "background_image": "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3",
        "signature": "我想休息",
        "total_favorited": 1,
        "work_count": 3,
        "favorite_count": 1
      },
      "content": "家人们谁懂啊",
      "create_date": "2025-07-15 21:04:11",
      "like_count": 87,
      "tease_count": 90
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» comment_list|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» user|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» follow_count|integer|true|none||none|
|»»» follower_count|integer|true|none||none|
|»»» is_follow|boolean|true|none||none|
|»»» avatar|string|true|none||none|
|»»» background_image|string|true|none||none|
|»»» signature|string|true|none||none|
|»»» total_favorited|integer|true|none||none|
|»»» work_count|integer|true|none||none|
|»»» favorite_count|integer|true|none||none|
|»» content|string|true|none||none|
|»» create_date|string|true|none||none|
|»» like_count|integer|true|none||none|
|»» tease_count|integer|true|none||none|

## POST 取消关注

POST /tiktok/relation/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|to_user_id|query|string| 是 |none|
|action_type|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "操作成功"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|

## GET 用户关注列表

GET /tiktok/relation/follow/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|token|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "获取关注列表成功",
  "user_list": [
    {
      "id": 2,
      "name": "durlim",
      "follow_count": 0,
      "follower_count": 1,
      "is_follow": true,
      "avatar": "",
      "background_image": "",
      "signature": "",
      "total_favorited": 0,
      "work_count": 0,
      "favorite_count": 0
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» user_list|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» name|string|false|none||none|
|»» follow_count|integer|false|none||none|
|»» follower_count|integer|false|none||none|
|»» is_follow|boolean|false|none||none|
|»» avatar|string|false|none||none|
|»» background_image|string|false|none||none|
|»» signature|string|false|none||none|
|»» total_favorited|integer|false|none||none|
|»» work_count|integer|false|none||none|
|»» favorite_count|integer|false|none||none|

## GET 用户粉丝列表

GET /tiktok/relation/follower/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|token|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "获取粉丝列表成功",
  "user_list": []
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» user_list|[any]|true|none||none|

## GET 用户好友列表

GET /tiktok/relation/friend/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|token|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "获取好友列表成功",
  "user_list": []
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» user_list|[any]|true|none||none|

## POST 发送消息

POST /tiktok/message/action

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|to_user_id|query|string| 是 |none|
|action_type|query|string| 是 |none|
|content|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "Send Message success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|

## GET 消息列表

GET /tiktok/message/chat

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 是 |none|
|to_user_id|query|string| 是 |none|
|pre_msg_time|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status_code": 0,
  "status_msg": "获取消息成功",
  "message_list": [
    {
      "id": 1,
      "from_user_id": 1,
      "to_user_id": 2,
      "content": "hello啊",
      "create_time": 1752626684
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status_code|integer|true|none||none|
|» status_msg|string|true|none||none|
|» message_list|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» from_user_id|integer|false|none||none|
|»» to_user_id|integer|false|none||none|
|»» content|string|false|none||none|
|»» create_time|integer|false|none||none|

## GET ping

GET /tiktok/ping

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

