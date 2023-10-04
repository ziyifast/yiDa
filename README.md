# 易搭

## 0 环境搭建
> 系统：MacOS
IDE：VSCode
vue:2.9.6
node:10.1.0
```bash
brew install npm
npm install vue-cli -g
vue init webpack yiDa
npm install element-ui --save
# 解决 ERR! ENOENT: no such file or directory, open,如果运行npm run dev报错，查看自己路径是否正确，npm run dev 默认会在当前寻找package.json
npm install axios

# 安装vuex，用于保存用户数据到浏览器
npm install vuex@^3.6.0
npm install --save vuex-persistedstate@^2.0.0
```


## 1 介绍
> V1版本

1. 用户访问易搭首页，会要求输入账号，如果首次登录系统则需创建用户，输入基本信息，即可创建完成
2. 登录易搭系统后，进入发现页面，可以看到其他人发布的搭活动
3. 用户如果对该活动感兴趣，则可以点击“开搭”，表明自己对该活动感兴趣，会按照要求如约参加该活动


## 2 技术选型
xorm+iris

## 3 项目结构
> 前后端分离项目

## 4 数据库设计
```sql
-- 用户表
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    password VARCHAR(20) NOT NULL,
    wechat VARCHAR(20) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    city VARCHAR(20) NOT NULL,
    description VARCHAR(100),
    birth_day DATE,
    hobbies jsonb default '[]',
    create_time timestamp NOT NULL,
    update_time timestamp NOT NULL
);


-- 活动详情表 
create table activity_detail (
    id serial primary key,
    theme varchar(20) not null,
    owner int not null, -- 活动发起人
    description varchar(100) , -- 活动描述
    user_list varchar(2048) not null, -- 参与人列表
    status int not null default 0, -- 0 未开始 1 进行中 2 已结束
    create_time timestamp not null,
    update_time timestamp not null
);

-- 将表activity_detail 字段update_time 从date 修改成timestamp
-- alter table activity_detail  alter column update_time  type timestamp using update_time::timestamp;


-- 用户与活动关联表
create table user_activity (
    uid int not null,
    aid int not null,
    primary key (uid, aid)
)
```

## 5 接口列表
### 5.1 用户相关
#### 5.1.1 用户注册
> 请求路径：/user/add
> 请求方式：POST
> 请求参数：
```json
{
    "id": 1,
    "username": "jack",
    "password": "14321421",
    "wechat": "fasdfafas",
    "phone": "18271629821",
    "city": "xiamen",
    "description": "my name is jack, i like hiking",
    "create_time": "2023-09-28 11:49:14",
    "update_time": "2023-09-28 11:49:14"
}
```

#### 5.1.2 用户登录
> 请求路径：/api/user/login
> 请求方式：POST
> 请求参数：

#### 5.1.3 修改用户
> 请求路径：/api/user/update
> 请求方式：POST

V2版本：
> 用户之间的聊天， 智能推荐