# Discovery 1

1. Resume From [🗂 Recon 1 & 2](../02-03-Recon%201%20&%202/)
2. Connect to SSH
```
ssh parhizkar@remote.petromaz.ir -p 2200
password : 3wymn92r3ENm
```
3. Find Application
```
# Proccess List
top

# Proccess List + Full Path
top -c -n 1
```
4. First Item: `/bin/bash /app/start.sh`
5. Read it: `cd /app; cat start.sh`
```
#! /bin/bash

# Write Flag File (Recon 2 Flag)
echo "$flag" > /flag.txt

# Run SSH Service
/usr/sbin/sshd -D -e $@ &

# Run Unknown File
./main &

tail -f /dev/null
```

6. Download `/app/main`
```
scp -P 2200 parhizkar@remote.petromaz.ir:/app/main ./06-main
```

7. Strings
```
# Search for MAZAPA
strings 06-main | grep MAZAPA

# Search for petromaz
strings 06-main | grep petro

# Search for mazehacker
strings 06-main | grep maze
```
8. Reverse Engineering `06-main` into `10-malware\main.go`
9. Get Encrypted Flag `nc mazehacker.onion 8080 > 09-mazehacker.json`
10. Run Decryption Code `go run 10-malware\main.go`


## Reverse Engineering Go Binary

- https://pkg.go.dev/
- https://www.youtube.com/watch?v=_cL-OwU9pFQ
- https://www.youtube.com/watch?v=YRqTrq11ebg