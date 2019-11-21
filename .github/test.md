
```
NAME:
   docktool test - test env/file/tcp

USAGE:
   docktool test [options]

DESCRIPTION:
   test env/file/tcp

OPTIONS:
   -e, --env <env-name>     env name for test, can be passed multiple times
   -f, --file <file-name>   file name for test, can be passed multiple times
   -d, --dir <dir-name>     dir name for test, can be passed multiple times
       --tcp <tcp-addr>     tcp addr for test, can be passed multiple times
       --help               print this usage

EXAMPLES:
   docktool test --env="JAVA_HOME"
   docktool test --file="/Users/yingzhuo/my-profile.sh"
   docktool test --dir="/tmp"
   docktool test --tcp="localhost:8080"

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/test.md

```

### Examples:

In your `docker-entrypoint.sh`

```sh
#!/bin/sh

set -e

docktool test --tcp="localhost:8080"
```
