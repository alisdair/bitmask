# bitmap stringification

Using UNIX permissions as an example, but ought to work for anything. Is there a simpler way?

```
go get golang.org/x/tools/cmd/stringer
go generate
go run mode.go mode_string.go
```
