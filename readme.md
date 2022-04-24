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
# using the return value:
tim -s && echo yes || echo no
# q returns 0.
# ctrl-c returns 127.
```

### `-c` countdown mode

```sh
tim -c 10s
tim -c 10h5m
tim -c 1m1s # order does not matter.

# using the return value:
tim -c 10s && echo timer ran out || echo timer stopped early
```

## development

```sh
# use the pre-commit hook:
git config --local core.hooksPath .githooks
```
