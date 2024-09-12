import json
import requests

headers = {
    'accept': 'application/json, text/javascript, */*; q=0.01',
    'accept-language': 'zh,en;q=0.9,zh-CN;q=0.8',
    # 'content-type': 'multipart/form-data; boundary=----WebKitFormBoundaryENyRbkcyMIRU0fc9',
    'cookie': 'JSESSIONID=4097707AF74FFE17E7C4A5F66555E2CF',
    'origin': 'https://www.wintone.com.cn',
    'priority': 'u=1, i',
    'referer': 'https://www.wintone.com.cn/productDetail/23/0?bd_vid=rjDzPW6YrjT1PWDYnHDsP1fsndtkrj0kg1FxnH0sg1wxrjm3rH6Ln1D1P10&sdclkid=A5epALopb5FpxLA&bd_vid=8186436089874723069',
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
    response = requests.post('https://www.wintone.com.cn/getSysTime', headers=headers)
    return response.text


files = {
    ("file", ('1.pdf', open('1.pdf', 'rb'), 'Content-Type: application/pdf')),
    ("typeId", (None, '20090')),
    ('phone', (None, 'null')),
    ('email', (None, 'null')),
    ('lang', (None, 'zh')),
    ('reconType', (None, '发票识别 识别')),
    ('sysTime', (None, getSysTime())),
}

response = requests.post('https://www.wintone.com.cn/onlineExperience', headers=headers, files=files)
items = json.loads(response.json()["data"])["cardsinfo"][0]["items"]
for item in items:
    print(f"{item['desc']}: {item['content']}")
