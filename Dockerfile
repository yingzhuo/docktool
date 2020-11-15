FROM busybox:1.32.0

LABEL maintainer="应卓 <yingzhor@gmail.com>"

COPY docktool /bin/docktool

ENTRYPOINT ["docktool", "-h"]