
```
NAME:
   docktool secret - encode/decode a string

USAGE:
   docktool secret [global options] COMMAND [command options] [arguments ...]

DESCRIPTION:
   encode/decode a string

COMMANDS:
   base64
   md4
   md5
   sha1
   sha256
   sha512
   sha384

GLOBALS OPTIONS:
   -h, --help   print this usage

EXAMPLES:
   docktool secret base64 "secret"
   docktool secret base64 "secret" -n
   docktool secret base64 --decode "c2VjcmV0"
   docktool secret md4 "secret"
   docktool secret md5 "secret"
   docktool secret sha1 "secret"
   docktool secret sha256 "secret"
   docktool secret sha512 "secret"
   docktool secret sha384 "secret"
   docktool secret bcrypt 'hello'
   docktool secret bcrypt --check 'hello' '$2a$10$tTD.tJFkdsJR67V6YQYinOAbcZEROfbIjz2wInYftX.DRhOV0OBe2'

SEE ALSO:
   https://github.com/yingzhuo/docktool/tree/master/.github/secret.md

Run 'docktool secret COMMAND --help' for more information on a command.

```
