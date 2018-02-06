FROM alpine
LABEL maintainer="drmfly.liw@gmail.com"
WORKDIR /srv/gopkg
ADD public ./public
ADD gopkg ./
ENTRYPOINT ["./gopkg"]