package domain

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input body")

	ErrUserNotFound            = errors.New("Не верный логин или пароль")
	ErrGenerateToken           = errors.New("Could not login")
	ErrNoAccessThisPage        = errors.New("Нет доступа к этой странице")
	ErrHeaderAuthorizedIsEmpty = errors.New("Пустой заголовок Authorized")
	ErrInvalidHeaderAuthorized = errors.New("Не валидный заголовок Authorized")
	ErrTokenIsEmpty            = errors.New("Токен пустой")
	ErrByDownloadAppTokenFile  = errors.New("Ошибка при скачивании токена приложения")
	ErrByDownloadUserTokenFile = errors.New("Ошибка при скачивании токена клиента")
	ErrInvalidApiKey           = errors.New("Не верный api key")
	ErrInvalidChannelId        = errors.New("Не верный channel id")
	ErrChannelIdNoUniqueness   = errors.New("Такой channel id уже используется")
	ErrUnableCreateUserToken   = errors.New("Неудалось создать токен пользователя")
	ErrUnableOpenUrl           = errors.New("Неудалось открыть ссылку в браузере")
	ErrMarkIsUses              = errors.New("Эта метка используется")
	ErrPhoneNoUniqueness       = errors.New("Такой номер уже используется")
	ErrByDownloadSessionFile   = errors.New("Ошибка при скачивании .session файла")

	ErrEmptyAuthHeader         = errors.New("empty authorization header")
	ErrInvalidAuthHeaderFormat = errors.New("invalid authorization header format")

	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")

	ErrUnableReadAppToken  = errors.New("Неудалось прочитать токен приложения")
	ErrUnableCreateConfig  = errors.New("Неудалось создать config на основании токена приложения")
	ErrUnableStartServer   = errors.New("Неудалось запустить веб сервер")
	ErrUnableOpenAuthUrl   = errors.New("Неудалось открыть ссылку для авторизации")
	ErrUnableRetrieveToken = errors.New("Неудалось получить токен из браузера")

	ErrConnectPostgreSQL = errors.New("all attempts are exceeded, unable to connect to PostgreSQL")

	ErrSessionNotFound      = errors.New("session not found")
	ErrSessionBlocked       = errors.New("session has been blocked")
	ErrIncorrectSessionUser = errors.New("incorrect session user")
	ErrMismatchedSession    = errors.New("mismatched session token")

	ErrSecretCodeInvalid = errors.New("code is incorrect")
	ErrSecretCodeExpired = errors.New("code is expired")
)
