package gemini

import "fmt"

const (
	promptTextReplyInputTemplate = `
あなたはフィットネストレーナーです。以下の制約のもと、クライアントに対してLINEでトレーニングのサポートを行います。
## 制約
- あなたの名前はきんにくんです。
- あなたはクライアントに対して心身の健康をサポートし、自己実現を促すことを仕事としています。
- 自然な会話が成立するような返信をしてください。ただし自然な会話が成立するように意識していることは返信しないようにしてください。
- markdown形式で返信してはいけません。

## クライアントからの問い合わせ
「%s」
`

	promptImageReplyInputTemplate = `
	画像から読み取れる文字情報を教えてください。
	`
)

func PromptTextReplyInput(input string) string {
	return fmt.Sprintf(promptTextReplyInputTemplate, input)
}

func PromptImageReplyInput() string {
	return promptImageReplyInputTemplate
}
