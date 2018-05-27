
DEV ENV:
- go 1.9+
- clone code:
    ```
    mkdir  $GOPATH/src/xgopkg.com
    cd $GOPATH/src/xgopkg.com
    git clone https://github.com/xgopkg/xgopkg.git
    ```
- server port:

    We need two port for running http server:
    1. 8080 is used for 'go get pkg'
    2. 8001 is used for api server & UI

- MySQL

    Install mysql in docker.
    ```bash
    docker run --name mysql -p 3306:3306 --restart=always \
    -v $HOME/docker/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD= -e TZ='Asia/Shanghai'  -d mysql:5.7
    ```
- Openapi docs

    http://host:8001/apidocs