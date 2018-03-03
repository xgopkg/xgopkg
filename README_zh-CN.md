# XGoPkg
## 开发指南
1. 获取代码
```shell
$ git clone https://github.com/xgopkg/xgopkg.git
```
2. dep安装 如已安装忽略此步骤

```shell
$ brew install dep
$ brew upgrade dep
```
&emsp;&emsp;[dep使用](https://golang.github.io/dep/)

3. 下载go依赖包
```shell
$ cd xgopkg
$ dep ensure
```

3. 前端

```shell
$ cd frontend
$ yarn install
$ yarn start
```
## 代码结构
```
.
├── Gopkg.lock
├── Gopkg.toml
├── LICENSE
├── README.md
├── README_zh-CN.md
├── docker
│   └── Dockerfile
├── frontend
│   ├── README.md
│   ├── config-overrides.js
│   ├── node_modules
│   ├── package.json
│   ├── public
│   ├── src
│   └── yarn.lock
├── pkg
│   ├── api
│   ├── client
│   ├── cmd
│   ├── db
│   └── resource
├── scripts
│   └── build.sh
└── vendor
    ├── github.com
    ├── golang.org
    └── gopkg.in
```
