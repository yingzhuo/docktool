# docktool

```
NAME:
   docktool - tools for docker container

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
   test      test env/file/tcp

GLOBALS OPTIONS:
   -q, --quiet     quiet mode (default: false)
       --help      print this usage
       --version   print version information

SEE ALSO:
   https://github.com/yingzhuo/docktool

Run 'docktool COMMAND --help' for more information on a command.

```

### Installing

#### On your computer

if you're using linux or macOS, you can install this tool by `curl` or `wget`.

```bash
VERSION=v1.1.x

# linux
sudo curl -fsSL "https://github.com/yingzhuo/docktool/releases/download/$VERSION/docktool-linux-amd64-$VERSION" -o /usr/local/bin/docktool
sudo chmod +x /usr/local/bin/docktool

# macOS
sudo curl -fsSL "https://github.com/yingzhuo/docktool/releases/download/$VERSION/docktool-darwin-amd64-$VERSION" -o /usr/local/bin/docktool
sudo chmod +x /usr/local/bin/docktool
```

if you're using windows, you can download it from [release page](https://github.com/yingzhuo/docktool/releases).

#### On docker image

also, you can install it on your docker image. for example:

```dockerfile
FROM busybox

ARG VERSION=v1.1.x

RUN wget https://github.com/yingzhuo/docktool/releases/download/$DOCKTOOL_VERSION/docktool-linux-amd64-$DOCKTOOL_VERSION -O /bin/docktool && \
    chmod +x /bin/docktool
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
