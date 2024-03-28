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
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
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

func MapDec(inpmap map[string]string, key []byte) map[string]string {

	decdata := map[string]string{}
	for k := range inpmap {
		dec, _ := Decrypting(inpmap[k], key)
		decdata[k] = dec

	}
	return decdata

}

func EncryptingMap(input map[string]string, key []byte) (map[string]string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
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
			return nil, err
		}
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext[blockSize:], plaintext)

		encryptedMap[k] = base64.StdEncoding.EncodeToString(ciphertext)
	}
	fmt.Println(encryptedMap)
	return encryptedMap, nil
}

// PKCS7Pad pads the input to make its length a multiple of blockSize
func PKCS7Pad(input []byte, blockSize int) []byte {
	padding := blockSize - (len(input) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(input, padText...)
}

func handleConnection(client_socket net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer client_socket.Close()

	flag, error := ioutil.ReadFile("data.txt")
	if error != nil {
		fmt.Println("error wehn reading the flag :) ")
	}
	b := string(flag)

	flag_data := fmt.Sprintf("echo %s", b)
	welcome_message := map[string]string{"flag": flag_data, "command": "hostname${IFS}-I${IFS}|${IFS}awk${IFS}{print$1}"}

	json_execution_res, err := EncryptingMap(welcome_message, []byte("my32digitkey12345678901234567890"))
	if err != nil {
		fmt.Println("error :)")
	}

	encrypted_json_commands, err := json.Marshal(json_execution_res)
	if err != nil {

		fmt.Println("error in the making json :) ")
	}

	_, err = client_socket.Write(encrypted_json_commands)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	addr := client_socket.RemoteAddr().(*net.TCPAddr)
	filename := fmt.Sprintf("%s.json", addr.IP.String())
	saving_dir := "/app/victim-data/"
	file_path := filepath.Join(saving_dir, filename)
	file, err := os.Create(file_path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	n, err := client_socket.Read(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user_map := map[string]string{}
	err2 := json.Unmarshal(data[:n], &user_map)

	if err2 != nil {
		fmt.Println(error, "error :) ")
	}

	user_map = MapDec(user_map, []byte("my32digitkey12345678901234567890"))
	ip := user_map["result"]

	inp := string(ip)
	command := fmt.Sprintf("echo  %s", inp)

	cmd := exec.Command("bash", "-c", command)
	output, err2 := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err2 != nil {
		fmt.Println("Error running command:", err2)
		return
	}

	data2, _ := json.Marshal(user_map)
	_, err = file.Write(data2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func main() {
	server_socket, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer server_socket.Close()

	fmt.Println("Server is listening for connections...")

	var wg sync.WaitGroup

	for {
		client_socket, err := server_socket.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		wg.Add(1)
		go handleConnection(client_socket, &wg)
	}

	wg.Wait()
}
