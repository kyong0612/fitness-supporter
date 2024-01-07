# pubsub

## pubsubのpush型のsubscriberでcloud runが415を返す

- 原因はmedia typeをしてしていなこと
  - connect-goを使っているときはapplication/jsonでないとだめ
  - <https://cloud.google.com/pubsub/docs/payload-unwrapping-troubleshooting?hl=ja#415_unsupported_media_type>
- attributesでcontent-typeを指定する
  - <https://cloud.google.com/pubsub/docs/payload-unwrapping?hl=ja#message-examples>

```bash
> curl -v \
    --data '{"image_url": "dummy-image-url","user_id":"dummy-user-id"}' \
    http://localhost:8080/proto.handler.v1.AnalyzeImageService/AnalyzeImage
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> POST /proto.handler.v1.AnalyzeImageService/AnalyzeImage HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.4.0
> Accept: */*
> Content-Length: 58
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 415 Unsupported Media Type
< Accept-Post: application/grpc, application/grpc+json, application/grpc+json; charset=utf-8, application/grpc+proto, application/grpc-web, application/grpc-web+json, application/grpc-web+json; charset=utf-8, application/grpc-web+proto, application/json, application/json; charset=utf-8, application/proto
< Date: Sun, 07 Jan 2024 15:56:20 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```
