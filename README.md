![gin-taro-api](./imgs/20220609102431.jpg)

# GIN-TARO-API

[gin framework](https://gin-gonic.com/zh-cn/docs/quickstart/)

[go-gin-api](https://github.com/xinliangnote/go-gin-api)

[moose-go](https://gitee.com/shizidada/moose-go)

## start

go run github.com/singcl/gin-taro-api -env pro

## mysql

```sql
-- root 用户登录
sudo mysql
-- 创建数据库
CREATE DATABASE gin_taro_api DEFAULT CHARACTER SET = 'utf8mb4';
/* 为远程用户授权 */
GRANT ALL PRIVILEGES ON gin_taro_api.* TO taro'@'127.0.0.1' WITH GRANT OPTION;
```

## mail
告警邮件发送调试通过 SUCCESS😇

## 界面
![授权列表](./imgs/authorized_list.png)
![新增授权](./imgs/authorized_add.png)
![授权详情](./imgs/authorized_detail.png)