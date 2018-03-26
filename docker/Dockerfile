FROM alpine:3.7
LABEL maintainer="dev_support@xgopkg.com"
WORKDIR /srv/xgopkg
COPY ["./xgopkg", "./config.yml","./"]
RUN apk --no-cache add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
ENTRYPOINT ["./xgopkg"]