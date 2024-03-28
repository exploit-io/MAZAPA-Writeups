# Recon 1 & 2

1. Search `petromaz` in LinkedIn
2. Follow
3. Posts: `Mahmoud Parhizkar`
4. Twitter: `MParhizkar`
5. First Tweet: `@ParhizGreenRoadBot`
6. Telegram Bot Flag: MAZAPA_c086438561ac9ea5559fd10e26d537e9
7. BlackMail `@MParhizkar54`
![Blackmail](./07-blackmail.png)

8. Get User, Password
```
remote.petromaz.ir:2200
user: parhizkar
password : 3wymn92r3ENm
```
9. Connect to SSH
```
ssh parhizkar@remote.petromaz.ir -p 2200
password : 3wymn92r3ENm

# inside ssh connection
cat /flag.txt
MAZAPA_f64106a3c687d38c0a399e26f1e722d7
```
![Get Flag](./09-ssh.png)