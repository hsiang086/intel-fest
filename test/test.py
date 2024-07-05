import requests
import json
import os

path = os.path.dirname(os.path.abspath(__file__))

url = "http://127.0.0.1:1588/api/signup"

payload = json.load(open(f"{path}/structure.json"))

p = json.dumps(payload["User"])

print(f"payload:\t{p}")

res = requests.post(url, data=p)

print(f"res:\t{res.json()}\ncookie:\t{res.cookies.get_dict()}")
