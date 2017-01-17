Задание: модуль Volume
=============================
База данных: PostgreSQL

Атрибуты
--------

 1. volume_id [uuid] - генерим в коде
 2. label [string]
 3. replicas [integer]  min: 2 max: 5
 4. volumeservers [volumeserver_id uuid array]
 5. limit [integer]
 6. user_id [uuid]
 7. created [timestamp without time zone] - генерим в PG
 8. active [boolean] - DEFAULT true
 9. exists [boolean] - DEFAULT false


Методы
------

 1. Add(Volume, *bool) - добавление нового Volume
 2. GetByUser(user_id, *[]Volume) - поиск всех volumes пользователя
 3. Get(volume_id, *Volume) - получение Volume
 4. Scale(Volume, *bool) - обновление replicas. Действует на replicas, volumeservers
 5. Rename(Volume, *bool) - обновление replicas. Действует на label
 6. Resize(Volume, *bool) - обновление. Действует на limit. Новый лимит всегда больше старого значения, иначе ошибка
 7. Activate(volume_id, *bool) - active -> true
 8. Deactivate(volume_id, *bool) - active -> false
 9. FindByVolumeserver(volumeserver_id, *[]Volume) - возврат всех Volume упомянутых в массиве volumeservers
 10. UsageByVolumeserver(volumeserver_id, integer) - сумма limit всех упомянутых в массиве volumeservers

Требования
----------

 1. Все методы работающие с БД должны быть написаны в файле sql.go
 2. Все методы кроме описанных в пункте Методы, должны быть приватными
 3. Каждый метод из пункта Методы, должен быть написан по требованиям пакета RPC (https://golang.org/pkg/net/rpc)
 4. SQL скрипт создания таблицы должен лежать в файле volume.sql
 5. Создавать индексы в БД не нужно
 6. Код должен быть задокументирован
 7. ORM не используется
 8. Перед запросом, должна быть валидация параметров


Используемые пакеты
-------------------

 1. net/rpc (https://golang.org/pkg/net/rpc)
 2. database/sql (https://golang.org/pkg/database/sql/)
 3. github.com/lib/pq
