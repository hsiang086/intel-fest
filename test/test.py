import requests
import json
import os
import argparse

parser = argparse.ArgumentParser()

parser.add_argument(
    "-s",
    "--signup",
    action="store_true",
    help="Sign up",
)
parser.add_argument(
    "-l",
    "--login",
    action="store_true",
    help="Login",
)
parser.add_argument(
    "-c",
    "--cookie",
    type=str,
    help="Get cookie",
)
parser.add_argument(
    "-p",
    "--path",
    type=str,
    help="Path",
)

args = parser.parse_args()

path = os.path.dirname(os.path.abspath(__file__))

url = "http://127.0.0.1:1588"

payload = json.load(open(f"{path}/structure.json" if not args.path else args.path))

cookie = None
if args.cookie:
    cookie = {
        "__yumm__": args.cookie
    }

if args.signup:
    res = requests.post(f"{url}/api/signup", json=payload["User"], cookies=cookie)
    print(f"res:\t{res.json()}\ncookie:\t{res.cookies.get_dict()}")
elif args.login:
    res = requests.post(f"{url}/api/login", json=payload["UserLogin"], cookies=cookie)
    print(f"res:\t{res.json()}\ncookie:\t{res.cookies.get_dict()}")

