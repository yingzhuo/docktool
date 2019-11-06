
```
NAME:
   docktool sleep - make current thread sleep

USAGE:
   docktool sleep [argument]

DESCRIPTION:
   make current thread sleep

OPTIONS:
   --help   print this usage

EXAMPLES:
   docktool sleep
   docktool sleep 5s
   docktool sleep 1m30s

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
