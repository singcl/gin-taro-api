![gin-taro-api](./imgs/20220609102431.jpg)

[![Build workflow](https://github.com/singcl/gin-taro-api/actions/workflows/build.yml/badge.svg)](https://docs.github.com/cn/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)

# GIN-TARO-API

Inspired by [go-gin-api](https://github.com/xinliangnote/go-gin-api). about some other resource [gin framework](https://gin-gonic.com/zh-cn/docs/quickstart/)
,[moose-go](https://gitee.com/shizidada/moose-go)

## Start

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

## Mail

告警邮件发送调试通过 SUCCESS😇

## Debug

在 vscode 中调试应用程序

相关参考文档

```sh
  # https://github.com/golang/vscode-go/blob/master/docs/debugging.md
  # https://code.visualstudio.com/docs/languages/go#_debugging
  # https://juejin.cn/post/6844904087763304462
```

1. 手动安装`delve` Go debugger

```sh
# Install the latest release:
$ go install github.com/go-delve/delve/cmd/dlv@latest

```

2. 配置 launch.json

```js
// Use IntelliSense to learn about possible attributes.
// Hover to view descriptions of existing attributes.
// For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
// https://github.com/golang/vscode-go/blob/master/docs/debugging.md
// https://code.visualstudio.com/docs/languages/go#_debugging
// https://juejin.cn/post/6844904087763304462
```

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}"
    }
  ]
}
```

## 微信小程序登陆核心代码片段

[小程序登陆](./Wx.md)

## CHANGELOG

project [changelog](./CHANGELOG.md) is hear.use [git-chaglog](https://github.com/git-chglog/git-chglog) auto generate changelog.

```sh

# 安装
go install -v github.com/git-chglog/git-chglog/cmd/git-chglog@latest

# git-chglog requires configuration files and templates to generate a CHANGELOG
# However, it is a waste of time to create configuration files and templates from scratch.
# Therefore we recommend using the --init option which will create them interactively 👍
git-chglog --init

# You are now ready for configuration files and templates!
# Let's immediately generate a CHANGELOG of your project. By doing the following simple command, Markdown for your CHANGELOG is displayed on stdout.
git-chglog

# Use -o (--output) option if you want to output to a file instead of stdout.
git-chglog -o CHANGELOG.md

```

## 管理界面

|                 预览                  |                预览                |
| :---------------------------------------: | :------------------------------------: |
|  ![授权列表](./imgs/authorized_list.png)  | ![新增授权](./imgs/authorized_add.png) |
| ![授权详情](./imgs/authorized_detail.png) |       -                                 |
