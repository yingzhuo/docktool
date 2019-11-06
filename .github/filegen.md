
```
NAME:
   docktool filegen - generate file using template

USAGE:
   docktool filegen [options]

DESCRIPTION:
   generate file using template from environment, arguments, json/yaml/toml config files

OPTIONS:
   -t, --template <file>           template file
       --json <file>               json file for input
       --yaml <file>               yaml file for input
       --toml <file>               toml file for input
   -v, --var <key>=<value>         vars for input, can be passed multiple times
       --delims [<left>:<right>]   template tag delimiters (default: {{:}})
       --stdin                     get template content from console
   -o, --output <file>             file for output
   -a, --append                    append output to file
       --perm <value>              perm of output file (default: 0660)
       --help                      print this usage

EXAMPLES:
   docktool filegen -t ./filegen.tmpl -v name=bill
   docktool filegen -t ./filegen.tmpl --json in.json --yaml in.yaml --toml in.toml
   docktool filegen -t ./filegen.tmpl -o ./gen.txt
   echo "{{ .foo }}" | docktool filegen --stdin -v foo=foo
   echo "{{ .ENV.foo }}" | foo=foo docktool filegen --stdin
   echo "{{ .ENV.HOME }}" | docktool filegen --stdin

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/filegen.md

```

### Examples

#### (1) template files

`/path/to/template.tmpl`

```txt
{{ .YAML.author.name }}
{{ .TOML.author.name }}
{{ .JSON.author.name }}
{{ .ENV.PWD }}
{{ .foo }}
{{ uuid32 }}
```

#### (2) input json

`/path/to/in.json`

```json
{
  "author": {
    "name": "应卓",
    "email": "yingzhor@gmail.com"
  }
}
```

#### (3) input yaml

`/path/to/in.yaml`

```yaml
author:
  name: "应卓"
  email: "yingzhor@gmail.com"
```

#### (4) toml yaml

`/path/to/in.toml`

```toml
[author]
    name = "应卓"
    email = "yingzhor@gmail.com"
```

#### (5) run command

```bash
docktool filegen \
    --template=/path/to/template.tmpl \
    --json=/path/to/in.json \
    --yaml=/path/to/in.yaml \
    --yaml=/path/to/in.toml \
    --var="foo=foo" \
    --output=/tmp/fg.txt
```

you will create a new file `/tmp/fg.txt`. it's content will be:

```txt
应卓
应卓
应卓
/Users/yingzhuo/开发/go-workspace/docktool
foo
0833bcd570e3925cbeb6fc9dd84a0513
```

**Note:** if template file NOT specified, `docktool` will using `$PWD/filegen.tmpl` (if file exists).
In the same way, will find default files.

- `--template`
    - `./filegen.tmpl`
    - `./docktool/filegen.tmpl`
- `--json`
    - `./filegen.in.json`
    - `./docktool/filegen.in.json`
- `--yaml`
    - `./filegen.in.yaml`
    - `./filegen.in.yml`
    - `./docktool/filegen.in.yaml`
    - `./docktool/filegen.in.yml`
- `--toml`
    - `./filegen.in.toml`
    - `./docktool/filegen.in.toml`

### Golang template funcs

See [fn package](https://github.com/yingzhuo/docktool/tree/master/fn).
