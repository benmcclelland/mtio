# mtio
golang library for mt ioctls

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/benmcclelland/mtio)

example using sample program:
```
# ./mtcmd /dev/nst0 status
2017/05/03 20:37:55 mtcmd.go:36: getting status
Generic SCSI-2 tape (114)
Residual count: 0
Device registers: 58000000
Status registers: 1010000
ONLINE IM_REP_EN
Error register: 0
Possibly inaccurate:
  Current file: 1
  Current block number: -1
```
