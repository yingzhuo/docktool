
```
NAME:
   docktool sleep - make current thread sleep

USAGE:
   docktool sleep [argument]

DESCRIPTION:
   make current thread sleep

OPTIONS:
   -h, --help   print this usage

EXAMPLES:
   docktool sleep
   docktool sleep 5s
   docktool sleep 1m30s
   docktool sleep '2019-11-08 12:00:00 +08:00'

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/sleep.md

```

### Examples:

In your `docker-entrypoint.sh`

```sh
#!/bin/sh

set -e

docktool sleep
```

this shell can make your container won't stop.
