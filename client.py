import argparse
import json
import socket
from datetime import datetime
from time import sleep


class TimeGetter:
    def __init__(self, host: str, port: int, token: str) -> None:
        self.host = host
        self.port = port
        self.request = {
            'id': 0,
            'method': 'TimeService.GetTime',
            'params': [{
                'AuthToken': token,
            }],
        }

    def get(self) -> datetime:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((self.host, self.port))
            s.sendall(json.dumps(self.request).encode('ascii'))
            response = s.recv(1024)

            data = json.loads(response.decode('ascii'))
            if data['error']:
                print(data['error']['message'])
                exit(1)
            elif data['result']['Status'] != 'ok':
                print(data['result']['Status'])
                exit(1)
            else:
                return datetime.fromtimestamp(data['result']['Time'])


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--host', default='localhost')
    parser.add_argument('--port', default=8088, type=int)
    parser.add_argument('--token', default='123456')
    args = parser.parse_args()

    time_getter = TimeGetter(args.host, args.port, args.token)
    while 1:
        remote = time_getter.get().strftime('%H:%M:%S')
        local = datetime.now().strftime('%H:%M:%S')
        print(f'\r{remote}(remote) | {local}(local)', end='')
        sleep(0.2)


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        pass
