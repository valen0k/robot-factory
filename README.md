# robot-factory

## Task

You own a robot factory producing vacuum cleaner and window cleaner robots. Produced robots are stored at your warehouse (which for simplicity has unlimited storage space) and then sold.

New robots are shipped to the warehouse at the end of the day they were produced and are removed from warehouse immediately after purchase.

Each robot (type) has its own manufacturing cost (coins per unit), storage cost (coins per unit per day), selling price (coins per unit) and manufacturing rate (units per day).
To be done

Design and develop a solution which allows to:

    Get the current remains of robots at warehouse
    Get the profit for arbitrary past period
    Sell robots to customers
    Get the projected profit for arbitrary future period based on selling statistics (advanced)

## Installation

Initialize database:

```shell
$ psql databasename < init/database.sql
```

Compile sources:

```shell
$ make
```

Add regular tasks into crontab:

```crontab
0 0 * * * /project-dir/updateStorageCost
0 0 * * * /project-dir/updateNumberRobots
```

Start service:

```shell
$ make server
```

API:

1. Get all robots `GET /robots`
2. Create robot `POST /robots`
3. Update robot `PUT /robots/{robotId}`
4. Delete robot `DELETE /robots/{robotId}`
5. Get robot `GET /robots/{robotId}`
6. Sell robots `PUT /robots/{robotId}/sell_robots`
7. Get profit `GET /profit`
