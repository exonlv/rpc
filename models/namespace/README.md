Задание: модуль Namespace
=============================
База данных: PostgreSQL

Атрибуты
--------

 1. id [uuid] - генерим в PG
 2. label [string]
 3. user_id [uuid]
 4. created [timestamp without time zone] - генерим в PG
 5. active [boolean] - DEFAULT: false
 6. removed [boolean] - DEFAULT: false
 7. kube_exist [boolean] - DEFAULT: false

Методы
------

 1. Add (Namespace, *bool) - Добавление новой записи в таблице Namespace. Передаются только: label, user_id
 2. Delete (Namespace, *bool) - изменение removed -> true
 3. GetAll (user_id string, *[]Namespace) - возврат всех Namespace пользователя
 4. Get (id string, *Namespace) - возврат конкретного Namespace пользователя
 5. Activate(id string, *bool) - изменение active -> true
 6. Deactivate (id string, *bool) - изменение active -> false
 7. CreatedInKube (id string, *bool) - изменение kube_exist -> true
 8. DeletedInKube (id string, *bool) - изменение kube_exist -> false
 8. Rename (Namespace, *bool) - изменение label
 
  

Требования
----------

 1. Все методы работающие с БД должны быть написаны в файле sql.go
 2. Все методы кроме описанных в пункте Методы, должны быть приватными
 3. Каждый метод из пункта Методы, должен быть написан по требованиям пакета RPC (https://golang.org/pkg/net/rpc)
 4. SQL скрипт создания таблицы должен лежать в файле namespace.sql
 5. Создавать индексы в БД не нужно
 6. Код должен быть задокументирован
 7. ORM не используется


Используемые пакеты
-------------------

 1. net/rpc (https://golang.org/pkg/net/rpc)
 2. database/sql (https://golang.org/pkg/database/sql/)
 3. github.com/lib/pq
