
```
NAME:
   docktool ping - test if tcp is reachable

USAGE:
   docktool ping [options ...] arg1 [arg2] [argN]

DESCRIPTION:
   test if tcp is reachable

OPTIONS:
   --logic <logic>   logic of testing (ALL | ANY | NONE) (default: ALL)
   --help            print this usage

EXAMPLES:
   docktool ping localhost:8080
   docktool ping localhost:3306 localhost:6379
   docktool ping --logic=ALL localhost:3306 localhost:6379
   docktool ping --logic=ANY localhost:3306 localhost:6379

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/ping.md

```

### Examples:

In your `docker-entrypoint.sh`

```sh
#!/bin/sh

docktool ping mysql:3306 redis:6379 && \
    java -jar /opt/my-application.jar "$@"

echo "Can NOT startup service"; false
```

Test if `mysql` or `redis` not ready first.
