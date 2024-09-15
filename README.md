# Env

Provides simple functions for working with ENV variables

## Usage

```go
Get(path string, value ...string) string
GetBool(path string, value ...bool) bool
GetInt(path string, value ...int) int
GetInt32(path string, value ...int32) int32
GetInt64(path string, value ...int64) int64
GetFloat32(path string, value ...float32) float32
GetFloat64(path string, value ...float64) float64
GetTime(path string, value ...time.Time) time.Time
GetDuration(path string, value ...time.Duration) time.Duration
GetMapStringBool(path string, defaults ...string) (ret map[string]bool)
GetMapStringInt(path string, defaults ...string) (ret map[string]int)
Set(path string, v string) error
Delete(path string) error
```
