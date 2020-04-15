FROM busybox:1.31.1

LABEL maintainer="应卓 <yingzhor@gmail.com>"

COPY docktool /bin/docktool

ENTRYPOINT ["docktool", "-h"]