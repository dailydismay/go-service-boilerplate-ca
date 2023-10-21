package login

import "authsvc/internal/core/domainerr"

/*
Для чего определять ошибки на уровне usecase?
Возможно это и не понадобилось бы, но поведенческие ошибки,
которые должны конкретезировать негативный результат исполнения кода usecase никакого отношения к доменному слою не имеют

Можно лишь не дублировать, а проверять еще на уровне errors.Is(user_domain.NotFound, например), но мне проще продублировать это здесь, чем импортить все домены в rpc/http/... хендлере
То есть - не импортим все подряд в Interface слое, ошибки поведения - только тут
*/
var (
	ErrAccountDoesNotExist = domainerr.New("account does not exist")
	ErrIncorrectPassword   = domainerr.New("err incorrect password")
)
