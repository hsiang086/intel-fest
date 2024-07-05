import requests
import json
import os

path = os.path.dirname(os.path.abspath(__file__))

url = "http://127.0.0.1:1588/api/signup"

payload = json.load(open(f"{path}/structure.json"))

p = json.dumps(payload["User"])

print(p)

res = requests.post(url, data=p)

print(res.json())
