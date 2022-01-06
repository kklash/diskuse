# `diskuse`

Human-readable pretty-printing of file sizes.

```go
diskuse.ReadableFileSize(234) // "234 b"
diskuse.ReadableFileSize(38812) // "37.9 KiB"
diskuse.ReadableFileSize(99299992) // "94.7 MiB"
diskuse.ReadableFileSize(711222222222222) // "647 TiB"
```
