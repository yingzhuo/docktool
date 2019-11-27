
```
NAME:
   docktool filedel - delete files/dirs using wildcard

USAGE:
   docktool filedel [options]

DESCRIPTION:
   delete files/dirs using wildcard

OPTIONS:
       --dir <dir>            start dir
   -p, --pattern <wildcard>   wildcard pattern, can be passed multiple times
       --delete-empty-dir     delete empty dir at last (default: false)
       --help                 print this usage

EXAMPLES:
   docktool filedel --dir=/my/dir -p=*.yml -p=*yaml
   docktool filedel --dir=/my/dir -p=data[0-9].csv
   docktool filedel --dir=/my/dir -p=*.txt -p=*.md -p=LICENSE
   docktool filedel --dir=/my/dir -p=sub-dir-*/
   docktool filedel --dir=/my/dir --delete-empty-dir

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/filedel.md

```

### Examples:

In your `Dockerfile`

```dockerfile
# ...

RUN tar -zxf /opt/zookeeper-3.5.6.tar.gz -C /opt && \
    mv /opt/apache-zookeeper-3.5.6-bin/ /opt/zookeeper && \
    docktool filedel --dir=/opt/zookeeper -p=*.txt -p=*.html
```

Remember! if you want to delete a directory, pattern must end with "/".