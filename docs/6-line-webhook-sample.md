# line webhook sample

- dump request

```bash
POST /line/webhook HTTP/1.1
Host: fitness-supporter-prod-x4zvxzl3ua-an.a.run.app
Content-Length: 564
Content-Type: application/json; charset=utf-8
Forwarded: for="147.92.149.166";proto=https
Traceparent: 00-581f9dfc02c2a06c332a2e6c7e628e6f-ab46feec28c3abad-01
User-Agent: LineBotWebhook/2.0
X-Cloud-Trace-Context: 581f9dfc02c2a06c332a2e6c7e628e6f/12341832119105072045;o=1
X-Forwarded-For: 147.92.149.166
X-Forwarded-Proto: https
X-Line-Signature: SPF+aswY1QKSO41ERcOs3IDju6ZkLHYDhKPbe6RUsZw=

{"destination":"Ued4095ce4cf48cdb76b561d079297db7","events":[{"type":"message","message":{"type":"text","id":"488830589727146065","quoteToken":"xxxxxxxx","text":"asdfadfa"},"webhookEventId":"01HK50YAT7682957NXRYYNARM6","deliveryContext":{"isRedelivery":false},"timestamp":1704197040565,"source":{"type":"user","userId":"Ub3460cc5512a5791102446b8d075b066"},"replyToken":"xxxxxx","mode":"active"}]}
```

- payload

```bash
curl --location 'http://localhost:8080/line/webhook' \
--header 'Content-Type: application/json' \
--data '{
  "destination": "Ued4095ce4cf48cdb76b561d079297db7",
  "events": [
    {
      "type": "message",
      "message": {
        "type": "text",
        "id": "488830589727146065",
        "quoteToken": "xxxxxxxx",
        "text": "asdfadfa"
      },
      "webhookEventId": "01HK50YAT7682957NXRYYNARM6",
      "deliveryContext": {
        "isRedelivery": false
      },
      "timestamp": 1704197040565,
      "source": {
        "type": "user",
        "userId": "Ub3460cc5512a5791102446b8d075b066"
      },
      "replyToken": "xxxxxx",
      "mode": "active"
    }
  ]
}'
```
