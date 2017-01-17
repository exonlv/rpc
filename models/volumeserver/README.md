Задание: модуль VolumeServer
=============================
База данных: PostgreSQL

Атрибуты
--------

 1. volumeserver_id [uuid] - генерим в PG
 2. ip [string]
 3. path [string]
 4. memory [integer]
 5. created [timestamp without time zone] - генерим в PG
 6. active [boolean] - DEFAULT: true
 7. group [string] - DEFAULT: "default"
 8. disk_type [string] - DEFAULT: "HDD"


Методы
------

 1. Add(VolumeServer, *bool) - добавление
 2. Activate(volumeserver_id, *bool) - active -> true
 3. Deactivate(volumeserver_id, *bool) - active -> false
 4. ChangeGroup(VolumeServer, *bool) - изменение group
 5. ChangeDiskType(VolumeServer, *bool) - изменение disk_type
 6. ChangePath(VolumeServer, *bool) - изменение path
 7. GetByGroup(group, *[]VolumeServer) - все VolumeServer определенной группы
 8. GetByActivity(bool active, *[]VolumeServer) - все VolumeServer с определнным active
 9. GetByDiskType(string disk_type, *[]VolumeServer) - все VolumeServer с определенным диском

Требования
----------

 1. Все методы работающие с БД должны быть написаны в файле sql.go
 2. Все методы кроме описанных в пункте Методы, должны быть приватными
 3. Каждый метод из пункта Методы, должен быть написан по требованиям пакета RPC (https://golang.org/pkg/net/rpc)
 4. SQL скрипт создания таблицы должен лежать в файле volumeserver.sql
 5. Создавать индексы в БД не нужно
 6. Код должен быть задокументирован
 7. ORM не используется
 8. Должна быть валидация перед каждым запросом


Используемые пакеты
-------------------

 1. net/rpc (https://golang.org/pkg/net/rpc)
 2. database/sql (https://golang.org/pkg/database/sql/)
 3. github.com/lib/pq
