FROM alpine
LABEL maintainer="drmfly.liw@gmail.com"
WORKDIR /srv/xgopkg
ADD public ./public
ADD xgopkg ./
ENTRYPOINT ["./xgopkg"]