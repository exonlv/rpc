Задание: модуль Status
=============================
База данных: Tarantool

Атрибуты
--------

 1. command_id [string] - обязателен unique
 2. user_id [string] - обязателен
 3. created [datetime] - обязателен now()
 4. status_id [integer] - обязателен : DEFAULT 0
 5. initiator [string] - обязательно


Методы
------

 1. Update (Status, *bool) - добавление нового статута или обновление 
 2. Get (command_id, *Status) - поиск по command_id
 3. FindByUser(user_id, *[]Status) - возвращение всех статусов пользователя


Требования
----------

 1. Все методы работающие с БД должны быть написаны в файле tarantool.go
 2. Все методы кроме описанных в пункте Методы, должны быть приватными
 3. Каждый метод из пункта Методы, должен быть написан по требованиям пакета RPC (https://golang.org/pkg/net/rpc)
 4. LUA скрипт создания таблицы должен лежать в файле tarantool.lua
 5. Необходимо создать индексы по полям требующиеся в SELECT методах 
 6. Код должен быть задокументирован
 7. ORM не используется
 8. Должна быть валидация переданных параметров перед каждым запросом (см. модель User)
 9. Типы сделать через указатели (см. модель User)


Используемые пакеты
-------------------

 1. net/rpc (https://golang.org/pkg/net/rpc)
 2. https://github.com/tarantool/go-tarantool