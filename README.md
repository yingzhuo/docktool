# docktool

```
NAME:
   docktool

USAGE:
   docktool [global options] COMMAND [command options] [arguments ...]

VERSION:
   v1.1.x

AUTHORS:
   应卓 <yingzhor@gmail.com>

COMMANDS:
   filegen   generate file using template
   filedel   delete files/dirs using wildcard
   wait      wait until tcp reachable or timeout
   sleep     make current thread sleep
   test      test env/file/dir/tcp

GLOBALS OPTIONS:
   -q, --quiet     quiet mode (default: false)
   -h, --help      print this usage
   -v, --version   print version information

SEE ALSO:
   https://github.com/yingzhuo/docktool

Run 'docktool COMMAND --help' for more information on a command.

```

### Installing

#### On your computer

You can download it from [release page](https://github.com/yingzhuo/docktool/releases) and you're able to verify the archives with GPG.

Author 's GPG fingerprint is: `3825 E69D 2277 CFA0 95F9  AA45 6B11 FB7F E9EC A55D`

#### On docker image

also, you can install it on your docker image. for example:

```dockerfile
FROM busybox

ARG DOCKTOOL_VERSION=1.1.x

RUN wget https://github.com/yingzhuo/docktool/releases/download/v${DOCKTOOL_VERSION}/docktool-linux-amd64-${DOCKTOOL_VERSION}.tar.gz -O ./docktool.tar.gz && \
    mkdir -p ./docktool/ && \
    tar -zxf ./docktool.tar.gz -C ./docktool/ && \
    cp ./docktool/docktool /bin/docktool && \
    chmod +x /bin/docktool && \
    rm -rf ./docktool ./docktool.tar.gz
```

#### Build it on your computer

```bash
git clone git@github.com:yingzhuo/docktool.git
cd ./docktool/
make install
```

### Sub Command

* [filegen](./.github/filegen.md)
* [filedel](./.github/filedel.md)
* [wait](./.github/wait.md)
* [sleep](./.github/sleep.md)
* [test](./.github/test.md)

### Contributing

* Fork it
* Create your feature branch (git checkout -b my-new-feature)
* Commit your changes (git commit -am 'add some feature')
* Push to the branch (git push origin my-new-feature)
* Create new Pull Request

### License

Apache 2.0 license. See [LICENSE](./LICENSE)

### Chang Log

See [chang log](./CHANGELOG.md)

### Authors

* 应卓 - [github](https://github.com/yingzhuo)

See also the list of [contributors](https://github.com/yingzhuo/docktool/graphs/contributors) who participated in this project.

### Acknowledgments

* [subchen](https://github.com/subchen)
	* [go-cli](https://github.com/subchen/go-cli)
