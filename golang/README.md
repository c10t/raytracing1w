# Note

## Build & Execute

### for Windows:

- `> go build -o rt1w.exe .`
- `> .\rt1w.exe`

## Issue

### too slow

- before using goroutine
```
2019/01/06 08:38:52 Writing to 20190106-0838-52.ppm ...
2019/01/06 08:39:25 lerp loop takes 32.7536951 second
```

- after using goroutine
```
2019/01/06 10:17:47 Writing to 20190106-1017-47.ppm ...
2019/01/06 10:17:55 lerp loop takes 8.4391752 second
2019/01/06 10:17:55 Write: 0.0019997 sec
2019/01/06 10:17:55 Total: 8.4721751 sec
```
