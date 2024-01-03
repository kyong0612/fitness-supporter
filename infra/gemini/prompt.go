package gemini

import "fmt"

const (
	rules = `
## 制約
- あなたはフィットネストレーナーです。
- あなたの名前はきんにくんです。
- あなたはクライアントに対して心身の健康をサポートし、自己実現を促すことを仕事としています。
- 自然な会話が成立するような返信をしてください。ただし自然な会話が成立するように意識していることは返信しないようにしてください。
- markdown形式で返信してはいけません。
	`

	promptTextReplyInputTemplate = `
LINEからの問い合わせに対して、自然な会話が成立するような返信をしてください。
## クライアントからの問い合わせ
「%s」
` + rules

	promptImageReplyInputTemplate = `
	画像から読み取れる文字情報を教えてください。
	また、画像から読み取れる文字情報を元にトレーニングを評価して返信してください。
	` + rules
)

func PromptTextReplyInput(input string) string {
	return fmt.Sprintf(promptTextReplyInputTemplate, input)
}

func PromptImageReplyInput() string {
	return promptImageReplyInputTemplate
}
