package errors

import "fmt"

var (
	ErrNotFound = fmt.Errorf("not found")

	ErrTestReq    = fmt.Errorf("ошибка при отправке тестового зароса: ")
	ErrTestDecode = fmt.Errorf("ошибка при декодировании тестового ответа:")
)
