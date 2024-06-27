# Unauthor

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yoshino-s/unauthor)

## Usage

```
Usage:
  unauthor [flags]
  unauthor [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     

Flags:
  -c, --concurrent int         concurrent number (default 20)
  -h, --help                   help for unauthor
  -t, --targets strings        target files or directories
  -f, --targets-file strings   target files or directories
  -T, --timeout duration       timeout seconds (default 10s)
      --type string            scan type, one of redis, zookeeper, memcached, dubbo, jdwp

Use "unauthor [command] --help" for more information about a command.
```