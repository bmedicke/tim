# tim

Tim is a (not yet) multifunction command line timer with the following features:

* stopwatch

## installation

```sh
go install github.com/bmedicke/tim@latest
```

## usage

* `-s` stopwatch mode
* `-c` countdown mode
* `-t` timer mode (counts up)
* `-q` quite flag, surpresses output

### `-s` stopwatch mode

```sh
# use return value:
tim -s && echo yes || echo no
# q returns 0.
# ctrl-c returns 127.
```

## development

```sh
# use the pre-commit hook:
git config --local core.hooksPath .githooks
```
