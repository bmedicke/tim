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
* `-x` timer/countdown mode exits with 0 when pressing *q*

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
tim -c 10h5s # not all values need to be specified.
tim -c 1m1s # order does not matter.

# using the return value:
tim -c 10s && echo timer ran out || echo timer stopped early
# q returns 1 because the countdown was stopped early.
# ctrl-c returns 127.

# making q return 0:
tim -c 10s -x ; echo $?
# ctrl-c still returns 127.
```

## development

```sh
# use the pre-commit hook:
git config --local core.hooksPath .githooks
```
