
```
NAME:
   docktool wait - wait until tcp reachable or timeout

USAGE:
   docktool wait [options ...]

DESCRIPTION:
   wait until tcp reachable or timeout

OPTIONS:
   -e, --environment-prefix [<prefix>]   environment name prefix of waiting list (default: WAIT_)
   -w, --wait <addr>                     addr to wait, can be passed multiple times
   -t, --timeout <time>                  timeout
       --logic <logic>                   logic of testing (ALL | ANY) (default: ALL)
       --help                            print this usage

EXAMPLES:
   docktool wait -w localhost:3306 -w localhost:6379
   docktool wait -e WAIT_
   docktool wait -w localhost:8080 -t "2019-11-02 14:34:40 +0800"
   docktool wait -w localhost:3306 -w localhost:6379 -w localhost:8080 --logic=ANY
   docktool wait -w "localhost:3306,localhost:6379"
   WAIT_MYSQL=localhost:3306 WAIT_REDIS=localhost:6379 docktool wait

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/wait.md


```

### Examples:

In your `docker-entrypoint.sh`

```sh
#!/bin/sh

set -e

docktool wait -w mysql:3306 -w redis:6379
java -jar /opt/my-application.jar
```

Before you startup your service, you can wait for `mysql` and `redis` get ready.
