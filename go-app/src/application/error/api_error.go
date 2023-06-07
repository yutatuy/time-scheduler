package application_error

/**
TODO: このエラーオブジェクトは下記の問題から、後々には回収した方がよい
- apiという単語名前に入っている。
- エラーのカテゴリーに合わせた名前にすべき。
*/

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}
