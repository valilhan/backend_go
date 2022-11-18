

Mail: ilyasovvalihan@gmail.com

**Download**:

```jsx
git clone https://github.com/valilhan/backend_go.git
cd backend_go
```

**Build project:**

```jsx
docker-compose up --build
```

**Starting project:**

```jsx
docker-compose up
```

**Migration**:

После запуска проекта необходимо выполнить миграцию базы данных с помощью

```jsx
make migrateup
```

## Метод начисления средств на баланс

- Метод позволяет начислить деньги на счёт пользователя.
- Создает в таблицы main_balance новый row с UserId, balance
- Создает нового пользователя с базовым балансом

**HTTP POST** [http://localhost:5001/fund](http://localhost:5001/fund)

Пример запроса JSON format:

```json
{
"UserId": "test",
"Balance": 100
}
```

Пример ответа без ошибки JSON format:

```json
{
"UserId": "test",
"Balance": 100
}
```

Пример ответа с ошибкой JSON format:

```json
{
"status_text": "Internal server error",
"message": "pq: duplicate key value violates unique constraint \"main_balance_pkey\""
}
```

```json
{
"status_text": "Bad request",
"message": "UserId is a required field"
}
```

## **Метод получения баланса пользователя**

- Метод позволяет узнать баланс пользователя.
- Если указанный пользователь существует, вернётся информация о его балансе, иначе - ошибка.

**HTTP GET** [http://localhost:5001/fund](http://localhost:5001/fund)

Пример запроса JSON format:

```json
{
"UserId": "test"
}
```

Пример ответа без ошибки JSON format:

```json
{
"UserId": "test",
"Balance": 100
}
```

Пример ответа с ошибкой JSON format:

```json
{
"status_text": "Bad request",
"message": "UserId is a required field"
}
```

```json
{
"status_text": "Bad request",
"message": "no matching record"
}
```

## **Метод резервирования средств с основного баланса на отдельном счете**

- Метод позволяет зарезервировать деньги пользователя за заказ на отдельном счёте.
- Создается row  с userId, ServiceId, OrderId, price в таблице reserve_balance
- (`orderId) - unique`повторяться не может.
- У одного пользователя может быть неограниченное количество заказов.
- Чтобы повторно заказать услугу, необходимо создать заказ с уникальным `orderId`.

**HTTP POST** [http://localhost:5001/reserve](http://localhost:5001/reserve)

Пример запроса JSON format:

```json
{
"UserId": "test",
"ServiceId" : "test",
"OrderId":"test",
"Price": 20
}
```

Пример ответа без ошибки JSON format:

```json
{
"UserId": "test",
"ServiceId" : "test",
"OrderId":"test",
"Price": 20
}
```

Пример ответа с ошибкой JSON format:

```json
{
"status_text": "Bad request",
"message": "ServiceId is a required filed"
}
```

```json
{
"status_text": "Internal server error",
"message": "pq: insert or update on table \"reserve_balance\" violates foreign key constraint \"reserve_balance_user_id_fkey\""
}
```

## **Метод признания выручки**

- Метод позволяет признать зарезервированную сумму.
- Если указанная сумма больше, чем зарезервированная, то происходит резервирование денег:
    - Приходит код с ошибкой
- Если указанная сумма меньше или равна зарезервированной, то происходит списание денег:
    - Создается в таблице revenue новый row c userId, serviceId, orderId, price
    - C таблицы main_balance у пользователя списывается деньги с основного баланса

**HTTP POST** [http://localhost:8000/accept](http://localhost:8000/accept)

Пример запроса JSON format:

```json
{
"UserId": "test",
"ServiceId" : "test",
"OrderId":"test",
"Price": 20
}
```

Пример ответа без ошибки JSON format:

```json
{
"UserId": "test",
"ServiceId" : "test",
"OrderId":"test",
"Price": 20
}
```

Пример ответа с ошибкой JSON format:

```json
{
"status_text": "Bad request",
"message": "ServiceId is a required filed"
}
```

```json
{
"status_text": "Internal server error",
"message": "pq: insert or update on table \"reserve_balance\" violates foreign key constraint \"reserve_balance_user_id_fkey\""
}
```