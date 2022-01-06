# `diskuse`

Human-readable pretty-printing of file sizes.

```go
fmt.Println(diskuse.ReadableFileSize(234)) // "234 b"
fmt.Println(diskuse.ReadableFileSize(38812)) // "37.9 KiB"
fmt.Println(diskuse.ReadableFileSize(99299992)) // "94.7 MiB"
fmt.Println(diskuse.ReadableFileSize(711222222222222)) // "647 TiB"
```
