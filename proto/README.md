# Proto Compilation

Install protoc compiler. Make sure it's in your PATH if you're on Windows.

```
# From top level tangle/. 
protoc proto\*.proto --go_out=proto\
```