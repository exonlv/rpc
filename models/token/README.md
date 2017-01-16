Задание: модуль Token
=============================
База данных: PostgreSQL

Атрибуты
--------

 1. token [uuid] - (PK)
 2. user_id [uuid] - (FK)
 3. created [timestamp without time zone] - генерим в PG
 3. expired [timestamp without time zone] - генерим в PG : created + 90 days
 4. active [boolean] - DEFAULT: true


Методы
------

 1. Add (Token, *bool) - добавление нового токена
 2. GetByUser(user_id, *Token) - поиск токена пользователя с условием active == true && expired >= now()
 3. Activate(token, *bool) - обновление active -> true у конкретного токена
 4. Deactivate(token, *bool) - обновление active -> false у конкретного токена
 5. GetAll(user_id, *[]Token) - поиск всех токенов пользователя
 6. Extend(token, *bool) - продлить created на 90 дней
 7. Check(token, *user_id) - найти токен пользователя c условием active == true && expired >= now()
 8. Get(token, *Token) - возврат конкретного токена


Требования
----------

 1. Все методы работающие с БД должны быть написаны в файле sql.go
 2. Все методы кроме описанных в пункте Методы, должны быть приватными
 3. Каждый метод из пункта Методы, должен быть написан по требованиям пакета RPC (https://golang.org/pkg/net/rpc)
 4. SQL скрипт создания таблицы должен лежать в файле token.sql
 5. Создавать индексы в БД не нужно
 6. Код должен быть задокументирован
 7. ORM не используется


Используемые пакеты
-------------------

 1. net/rpc (https://golang.org/pkg/net/rpc)
 2. database/sql (https://golang.org/pkg/database/sql/)
 3. github.com/lib/pq
