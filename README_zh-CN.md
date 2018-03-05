# go-rest-sample
## 开发指南
1. 获取代码
```shell
$ git clone https://github.com/chclouid/go-rest-sample.git
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
4

```
$ cd pkg/cmd/server
$ go run main.go
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
