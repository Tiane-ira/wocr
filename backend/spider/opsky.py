import json
import requests

headers = {
    'accept': 'application/json, text/javascript, */*; q=0.01',
    'accept-language': 'zh,en;q=0.9,zh-CN;q=0.8',
    # 'content-type': 'multipart/form-data; boundary=----WebKitFormBoundaryENyRbkcyMIRU0fc9',
    'cookie': 'JSESSIONID=4097707AF74FFE17E7C4A5F66555E2CF',
    'origin': 'https://www.opsky.com.cn',
    'priority': 'u=1, i',
    'referer': 'https://www.opsky.com.cn/productDetail/0/126?sdclkid=A5epALopb5FpxLA',
    'sec-ch-ua': '"Chromium";v="128", "Not;A=Brand";v="24", "Google Chrome";v="128"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-platform': '"macOS"',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'same-origin',
    'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36',
    'x-requested-with': 'XMLHttpRequest',
}


def getSysTime():
    response = requests.post('https://www.opsky.com.cn/getSysTime', headers=headers)
    return response.text


files = {
    ("file", ('1.pdf', open('1.pdf', 'rb'), 'Content-Type: application/pdf')),
    ("typeId", (None, '20090')),
    ('phone', (None, 'null')),
    ('email', (None, 'null')),
    ('lang', (None, 'zh')),
    ('reconType', (None, '慧票通票据识别系统 识别')),
    ('sysTime', (None, getSysTime())),
}

response = requests.post('https://www.opsky.com.cn/onlineExperience', headers=headers, files=files)
items = json.loads(response.json()["data"])["cardsinfo"][0]["items"]
for item in items:
    print(f"{item['desc']}: {item['content']}")
