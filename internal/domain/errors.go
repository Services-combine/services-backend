package domain

import "errors"

var (
	ErrUserNotFound            = errors.New("Не верный логин или пароль")
	ErrGenerateToken           = errors.New("Could not login")
	ErrNoAccessThisPage        = errors.New("Нет доступа к этой странице")
	ErrHeaderAuthorizedIsEmpty = errors.New("Пустой заголовок Authorized")
	ErrInvalidHeaderAuthorized = errors.New("Не валидный заголовок Authorized")
	ErrTokenIsEmpty            = errors.New("Токен пустой")
	ErrByDownloadTokenFile     = errors.New("Ошибка при скачивании токена приложения")
	ErrInvalidApiKey           = errors.New("Не верный api key")
	ErrInvalidChannelId        = errors.New("Не верный channel id")
	ErrChannelIdNoUniqueness   = errors.New("Такое channel id уже используется")
	ErrUnableCreateUserToken   = errors.New("Неудалось создать токен пользователя")
	ErrUnableOpenUrl           = errors.New("Неудалось открыть ссылку в браузере")

	ErrUnableReadAppToken  = errors.New("Неудалось прочитать токен приложения")
	ErrUnableCreateConfig  = errors.New("Неудалось создать config на основании токена приложения")
	ErrUnableStartServer   = errors.New("Неудалось запустить веб сервер")
	ErrUnableOpenAuthUrl   = errors.New("Неудалось открыть ссылку для авторизации")
	ErrUnableRetrieveToken = errors.New("Неудалось получить токен из браузера")
)
