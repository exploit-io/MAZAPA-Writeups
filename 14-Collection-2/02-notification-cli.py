#! /usr/bin/python3

import json
import sys

import requests
import websocket


# Function to login and get the session ID
def login(username, password):
    login_url = "http://notification.petromaz.ir/accounts/login/"  # Update with your login URL
    response = requests.post(login_url, data={'username': username, 'password': password, "isAPI": True})
    if response.status_code == 200:
        cookies = response.cookies.get_dict()
        session_id = cookies.get('sessionid')
        return session_id
    else:
        return None


# Function to send data to WebSocket URL
def send_data_to_websocket(session_id, message, chat_user, username):

    websocket_url = f"ws://notification.petromaz.ir/ws/chat/{username}_{chat_user}/"  # Update with your WebSocket URL
    # headers = {
    #     "Cookie": f"sessionid={session_id}"
    # }
    ws = websocket.create_connection(websocket_url)

    # Sending data to WebSocket view
    data = {
        "command": "new_message",
        "message": message,
    }
    ws.send(json.dumps(data))
    print("message sent...")
    # Close WebSocket connection
    ws.close()


def receive_data_from_websocket(session_id, num, chat_user, username):
    url = "http://notification.petromaz.ir/chat/nlast_messages/"  # Update with your WebSocket URL
    data = {
        "num": num,
        "session_id": session_id,
        "chat_user": chat_user,
    }
    response = requests.post(url, data)
    serialized_messages = response.json()
    serialized_messages = list(reversed(serialized_messages))
    for message in serialized_messages:
        print(f"{message['author']}:\t{message['content']}")


if __name__ == "__main__":
    # Replace 'your_username' and 'your_password' with actual credentials
    try:
        username = sys.argv[1]
        password = sys.argv[2]
        mode = sys.argv[3]
        if mode == "s":
            message = sys.argv[4]
            chat_user = sys.argv[5]
        else:  # 'r'
            num = sys.argv[4]
            chat_user = sys.argv[5]
    except Exception:
        raise Exception(
            '''\n 
                Use it to communicate with other people.
                use: notification-cli [username] [password] [mode:s | r] ([message] [chat_user] | [num] [chat_user])\n
                Example: \n
                sending message: notification-cli hossein 1234 s "hello" ali \n
                receiving message: notification-cli hossein 1234 r 20 ali
            '''
        )
    # Login and get session ID
    session_id = login(username, password)

    if session_id:
        # Send data to WebSocket URL using the obtained session ID
        if mode == "s":
            send_data_to_websocket(session_id, message, chat_user, username)
        else:  # 'r'
            receive_data_from_websocket(session_id, num, chat_user, username)
    else:
        print("Login failed.")