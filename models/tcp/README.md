Задание: модуль Tcp
=============================
База данных: PostgreSQL

Атрибуты
--------

 1. id [uuid] - генерим в PG
 2. user_id [uuid]
 3. channel [string]
 4. active [timestamp] - генерирует PG
 5. opened [bool]
 6. ip [string]
 
 \* В качестве генератора uuid использовать uuid_generate_v4()

Методы
------

 1. Open (Tcp, *bool) - Добавление новой записи в таблице Tcp. Если запись с **user_id и channel** уже существует, изменить opened->true и обновить время в active. Передаются **user_id, channel, ip**.
 2. Close (Tcp, *bool) - изменение **opened -> false**
 3. GetAll (user_id string, *[]Tcp) - возврат всех Tcp пользователя
 4. Get (id string, *Tcp) - возврат конкретного Tcp
 
  

Требования
----------

 1. Все методы работающие с БД должны быть написаны в файле **sql.go**
 2. Все методы кроме описанных в пункте Методы, должны быть **приватными**
 3. Каждый метод из пункта Методы, должен быть написан по требованиям пакета **RPC** (https://golang.org/pkg/net/rpc)
 4. SQL скрипт создания таблицы должен лежать в файле **tcp.sql**
 5. Создавать индексы в БД не нужно
 6. Код должен быть задокументирован
 7. **ORM не используется**


Используемые пакеты
-------------------

 1. net/rpc (https://golang.org/pkg/net/rpc)
 2. database/sql (https://golang.org/pkg/database/sql/)
 3. github.com/lib/pq
