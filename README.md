# apps_by_platform

Перед деплоем в файле config.json значение ключа сессий вынести в энвы:
```
"sessions": {
"key": "MY_SESSION_KEY", - должно быть в .env 
"cyclekey": true
}
```