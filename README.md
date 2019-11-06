# docktool

```
NAME:
   docktool - tool set for docker container

USAGE:
   docktool [global options] COMMAND [command options] [arguments ...]

VERSION:
   1.x.x

AUTHORS:
   应卓 <yingzhor@gmail.com>

COMMANDS:
   filegen   generate file using template
   wait      wait until tcp reachable or timeout
   ping      test if tcp is reachable
   sleep     make current thread sleep

GLOBALS OPTIONS:
   -q, --quiet     quiet mode
       --help      print this usage
       --version   print version information

SEE ALSO:
   https://github.com/yingzhuo/docktool

Run 'docktool COMMAND --help' for more information on a command.

```

### Installation

you can install this tool by `curl` or `wget`.

#### On your computer

```bash
VERSION=1.x.x

# linux
sudo curl -fsSL "https://github.com/yingzhuo/docktool/releases/download/$VERSION/docktool-linux-amd64-$VERSION" -o /usr/local/bin/docktool
sudo chmod +x /usr/local/bin/docktool

# macOS
sudo curl -fsSL "https://github.com/yingzhuo/docktool/releases/download/$VERSION/docktool-darwin-amd64-$VERSION" -o /usr/local/bin/docktool
sudo chmod +x /usr/local/bin/docktool
```

#### On docker image

also, you can install it on your docker image. for example:

```dockerfile
FROM busybox

ENV DOCKTOOL_VERSION=1.x.x

RUN wget --no-check-certificate "https://github.com/yingzhuo/docktool/releases/download/$DOCKTOOL_VERSION/docktool-linux-amd64-$DOCKTOOL_VERSION" -O /bin/docktool && \
    chmod +x /bin/docktool

CMD ["docktool --help"]
```

#### Build it on your computer

```bash
git clone git@github.com:yingzhuo/docktool.git
cd ./docktool/
make install
```

### Sub Command

* [filegen](./.github/filegen.md)
* [wait](./.github/wait.md)
* [ping](./.github/ping.md)
* [sleep](./.github/sleep.md)

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
