package domain

import "errors"

var (
	ErrUserNotFound            = errors.New("Не верный логин или пароль")
	ErrGenerateToken           = errors.New("Could not login")
	ErrByDownloadTokenFile     = errors.New("Ошибка при скачивании токен файла")
	ErrInvalidApiKey           = errors.New("Не верный api key")
	ErrInvalidChannelId        = errors.New("Не верный channel id")
	ErrChannelIdNoUniqueness   = errors.New("Такое channel id уже используется")
	ErrNoAccessThisPage        = errors.New("Нет доступа к этой странице")
	ErrHeaderAuthorizedIsEmpty = errors.New("Пустой заголовок Authorized")
	ErrInvalidHeaderAuthorized = errors.New("Не валидный заголовок Authorized")
	ErrTokenIsEmpty            = errors.New("Токен пустой")
)
