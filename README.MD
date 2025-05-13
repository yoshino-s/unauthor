# Unauthor

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/yoshino-s/unauthor)

## Usage

```
Usage:
  soar-helper [flags]
  soar-helper [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     

Flags:
      --generate-config.enable          generate config enable
      --generate-config.format string   generate config format, one of json, yaml, env (default "yaml")
      --generate-config.path string     generate config path
  -h, --help                            help for soar-helper
      --log.file string                 log file path
      --log.format string               log format, one of json, console, empty for default (console for dev, json for prod)
      --log.level string                log level (default "info")
      --log.levels.console string       log level for console, empty for same as log.level
      --log.levels.file string          log level for file, empty for same as log.level
      --log.otel                        enable sending logs to otel collector
      --log.rotate.enable               enable log rotation
      --log.rotate.max_age int          max age of log file in days (default 28)
      --log.rotate.max_backups int      max number of log file backups (default 3)
      --log.rotate.max_size int         max size of log file in MB (default 500)
  -c, --scanner.concurrent int          concurrent (default 100)
  -p, --scanner.protocol string         protocol, one of zookeeper,redis,memcached,ftp,dubbo,jdwp
  -t, --scanner.targets strings         targets
  -f, --scanner.targets-file strings    targets file
  -T, --scanner.timeout duration        timeout (default 5s)

Use "soar-helper [command] --help" for more information about a command.
```