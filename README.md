# robot-factory

### Task

You own a robot factory producing vacuum cleaner and window cleaner robots. Produced robots are stored at your warehouse (which for simplicity has unlimited storage space) and then sold.

New robots are shipped to the warehouse at the end of the day they were produced and are removed from warehouse immediately after purchase.

Each robot (type) has its own manufacturing cost (coins per unit), storage cost (coins per unit per day), selling price (coins per unit) and manufacturing rate (units per day).
To be done

Design and develop a solution which allows to:

    Get the current remains of robots at warehouse
    Get the profit for arbitrary past period
    Sell robots to customers
    Get the projected profit for arbitrary future period based on selling statistics (advanced)

### Использование

Перед запуском сервиса, необходимо, добавить тип в базу данных с помощью команды `create type transaction as enum ('STORAGE', 'SALE');`

Для запуска сервиса и компиляции бинарных файлов необходимо прописать в консоли `make`. Сервис запущен, останется добавить бинарные файлы в планировщик Cron. Запись может выглядеть следующим образом `0 0 * * * path-to-bin/updateStorageCost` и `0 0 * * * path-to-bin/updateNumberRobots`.
    