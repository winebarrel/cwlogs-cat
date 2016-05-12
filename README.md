# cwlogs-cat

Put message to  to CloudWatch Logs.

## Usage

```
Usage of ./cwlogs-cat:
  -g string
      log-group-name
  -s string
      log-stream-name
```

```sh
cat hello | cwlogs-cat -g my-group -s my-stream
```
