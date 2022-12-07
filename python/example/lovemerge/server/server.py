import json

import http.server

RouterPath = {
}


class Router(object):
    @staticmethod
    def register_route(path, func):
        RouterPath[path] = func

    @staticmethod
    def get_func(path):
        return RouterPath[path]


class RequestHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        func = Router.get_func(self.path)
        data = self.rfile.read(int(self.headers['content-length']))
        print(data)
        code, info = func(json.loads(data))
        self.send_response(code)
        self.wfile.write(json.dumps(info, indent=4).encode('utf-8'))

    def do_POST(self):
        path = self.path
        self.send_response(200)
        self.wfile.write("hello\n".encode('utf-8'))


def start_server():
    server_address = ('127.0.0.1', 12345)
    server = http.server.HTTPServer(server_address, RequestHandler)
    server.serve_forever()
