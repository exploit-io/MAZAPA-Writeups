#!/bin/bash

username=""
password="" 

server_ip=$(ip r | grep default | cut -f 3 -d ' ')
server_ip=""
case "$1" in
logout)
	curl --insecure -X GET https://$server_ip:4081/internal/logout
	;;
*)

	if [ -z $username ]; then
		read -p "Enter the Username:" username
	fi

	# if password is empty get it!
	if [ -z $password ]; then
		read -sp "Password:" password
	fi

	curl --insecure -X POST https://$server_ip:4081/internal/dologin.php?NTLM=0 --data-urlencode "kerio_username=$username" --data-urlencode "kerio_password=$password"

	res=$(curl -I https://www.google.com -m 5 2>/dev/null | head -n 1 | cut -f 2 -d " ")
	if [[ $res == "200" ]]; then
		printf "\033[1m\033[92mYou Are Connected Now!\033[0m"
	else
		printf "\033[1m\033[91mNot Connected!\033[0m"
	fi
	printf "\n\n"
esac
