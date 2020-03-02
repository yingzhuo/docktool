FROM busybox:1.31.1

LABEL maintainer="应卓"

COPY docktool /bin/docktool

CMD ["docktool", "-v"]