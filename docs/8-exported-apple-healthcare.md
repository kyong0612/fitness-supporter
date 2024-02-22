# Exported Apple Healthcare Data

## Overview

- apple health care dataを永続化するAPIを提供する

## Requirements

- POST /sync/healthcare/apple
  - Request
    - `Authorization: Basic <base64 encoded username:password>`
  - Response
    - `200 OK`
    - `400 Bad Request`
    - `401 Unauthorized`
    - `500 Internal Server Error`
  - Logic
    - exportされたmetricsをlogに出力する
    - exportされたmetricsをbigqueryに永続化する

## References

- <https://zenn.dev/miketako3/articles/0705876f451f8b>
- <https://www.healthexportapp.com/>
