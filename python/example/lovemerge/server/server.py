import http.server


class RequestHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.wfile.write("hello\n".encode('utf-8'))

def start_server():
    server_address = ('127.0.0.1', 12345)
    server = http.server.HTTPServer(server_address, RequestHandler)
    server.serve_forever()