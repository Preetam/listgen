listgen [![Build Status](https://drone.io/github.com/PreetamJinka/listgen/status.png)](https://drone.io/github.com/PreetamJinka/listgen/latest) [![BSD License](https://img.shields.io/pypi/l/Django.svg)](https://github.com/PreetamJinka/listgen/blob/master/LICENSE)
===
listgen is a lock-free ordered linked list generator for Go.

Usage
---
First get listgen using `go get github.com/PreetamJinka/listgen`.

```
$ listgen -h
Usage of ./listgen:
  -cmp="": Comparison function body. The argument names are `a' and `b'.
  -list-type="List": List type.
  -out="": Output file. Leave blank for stdout.
  -package="": Package name to use for the list.
  -value-type="": Value type.
```

For example, the stringlist implementation is generated using:

```bash
$ listgen \
  -package "stringlist" \
  -list-type "StringList" \
  -value-type "string" \
  -cmp "if a < b {return -1}; if a > b {return 1}; return 0;" \
  -out "stringlist/stringlist.go"

```

License
---
BSD (see LICENSE)
