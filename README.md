# Rectangles exercise

## Profiling the code

These commands can help to exercise on the measurement of the performance
Based on [golang's blog article](https://blog.golang.org/pprof).

```bash
go test --bench . --benchmem -memprofile memprofile.prof
go tool pprof -text ./memprofile.prof
go test --bench . --benchmem -cpuprofile cpuprofile.prof
go tool pprof -text ./cpuprofile.prof
```
