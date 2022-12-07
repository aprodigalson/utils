from server.server import start_server, Router
from application.app import login


def register_route():
    Router.register_route('/login', login)


def start():
    register_route()
    start_server()


if __name__ == "__main__":
    start()
