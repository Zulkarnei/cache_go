# cache_go

Реализован пакет in-memory-cache со следуюшими методами:

- `Set(key string, value interface{})` - запись значения `value` в кеш по ключу `key`
- `Get(key string)`
- `Delete(key)`

Пакет может быть импортирован в виде библиотеки. Чтобы установить его в себе проект введите команду в terminal go get -u github.com/Zulkarnei/cache_go
