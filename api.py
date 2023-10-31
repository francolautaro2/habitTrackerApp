from flask import Flask

app = Flask(__name__)

@route.get("/")
def hello_world():
    return "hello world!"

print("hello")
