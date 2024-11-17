Технологии индустриального программирования
Курнаков Кирилл Александрович ЭФМО-02-24

# Практическая работа 9

## Создание таблицы category
```
CREATE TABLE category (
    category_id VARCHAR(2) PRIMARY KEY,
    name VARCHAR(40) NOT NULL
);
```

## Создание таблицы item
```
CREATE TABLE item (
    item_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(40) NOT NULL,
    category_id VARCHAR(2) REFERENCES category(category_id),
    description VARCHAR(100) NOT NULL,
    price FLOAT NOT NULL,
    quantityStock INTEGER NOT NULL
);
```

## Создание таблицы package
```
CREATE TABLE package (
    package_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    price FLOAT NOT NULL
);
```

## Создание таблицы userType
```
CREATE TABLE userType (
    userType_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
```

## Создание таблицы user
```
CREATE TABLE "user" (
    user_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(18) NOT NULL,
    userType_id VARCHAR(5) REFERENCES userType(userType_id)
);
```

## Создание таблицы review
```
CREATE TABLE review (
    review_id VARCHAR(5) PRIMARY KEY,
    rating FLOAT NOT NULL,
    comment VARCHAR(100) NOT NULL,
    reviewDate TIMESTAMP NOT NULL,
    anonymous BOOLEAN NOT NULL,
    item_id VARCHAR(5) REFERENCES item(item_id)
);
```

## Создание таблицы orderStatus
```
CREATE TABLE orderStatus (
    orderStatus_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
```

## Создание таблицы delivery
```
CREATE TABLE delivery (
    deliveryType_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    numberKilometer FLOAT NOT NULL,
    price FLOAT NOT NULL,
    timeDelivery TIME NOT NULL
);
```

## Создание таблицы order
```
CREATE TABLE "order" (
    order_id VARCHAR(5) PRIMARY KEY,
    user_id VARCHAR(5) REFERENCES "user"(user_id),
    orderDate TIMESTAMP NOT NULL,
    deliveryType_id VARCHAR(5) REFERENCES delivery(deliveryType_id),
    addressDelivery VARCHAR(50) NOT NULL,
    numberKilometer FLOAT NOT NULL,
    deliveryPrice FLOAT NOT NULL,
    finalPrice FLOAT NOT NULL,
    orderStatus_id VARCHAR(5) REFERENCES orderStatus(orderStatus_id)
);
```

## Создание таблицы paymentType
```
CREATE TABLE paymentType (
    paymentType_id VARCHAR(5) PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
```

## Создание таблицы payment
```
CREATE TABLE payment (
    payment_id VARCHAR(5) PRIMARY KEY,
    order_id VARCHAR(5) REFERENCES "order"(order_id),
    paymentDate TIMESTAMP NOT NULL,
    paymentType_id VARCHAR(5) REFERENCES paymentType(paymentType_id),
    finalPrice FLOAT NOT NULL,
    confirmation BOOLEAN NOT NULL
);
```

## Создание таблицы orderItem
```
CREATE TABLE orderItem (
    orderItem_id VARCHAR(5) PRIMARY KEY,
    order_id VARCHAR(5) REFERENCES "order"(order_id),
    item_id VARCHAR(5) REFERENCES item(item_id),
    quantity INTEGER NOT NULL,
    sumPriceItem FLOAT NOT NULL,
    package_id VARCHAR(5) REFERENCES package(package_id),
    quantityPackage INTEGER NOT NULL,
    sumPricePackage FLOAT NOT NULL
); 
```

## Вставка в таблицу category
```
/*INSERT INTO category VALUES 
('C1', 'Мягкая игрушка'),
('C2', 'Кружка'),
('C3', 'Кофта');
```

## Вставка в таблицу item
```
INSERT INTO item VALUES 
('I001', 'Мишка', 'C1', 'Мягкая игрушка в виде мишки', 850.00, 15),
('I002', 'Собачка', 'C1', 'Мягкая игрушка в виде собачки', 850.00, 15),
('I003', 'Кошка', 'C1', 'Мягкая игрушка в виде кошки', 850.00, 15),
('I004', 'Птичка', 'C1', 'Мягкая игрушка в виде птички', 850.00, 15),
('I005', 'Капибара', 'C1', 'Мягкая игрушка в виде капибары', 850.00, 15),

('I006', 'Кружка с мишкой', 'C2', 'Прикольная кружка с мишкой', 450.00, 10),
('I007', 'Кружка с собачкой', 'C2', 'Прикольная кружка с собачкой', 450.00, 10),
('I008', 'Кружка с кошкой', 'C2', 'Прикольная кружка с кошкой', 450.00, 10),
('I009', 'Кружка с птичкой', 'C2', 'Прикольная кружка с птичкой', 450.00, 10),
('I0010', 'Кружка с капибарой', 'C2', 'Прикольная кружка с капибарой', 450.00, 10),

('I0011', 'Кофта с мишкой', 'C3', 'Крутая кофта с мишкой', 850.00, 30),
('I0012', 'Кофта с собачкой', 'C3', 'Крутая кофта с собачкой', 850.00, 30),
('I0013', 'Кофта с кошкой', 'C3', 'Крутая кофта с кошкой', 850.00, 30),
('I0014', 'Кофта с птичкой', 'C3', 'Крутая кофта с птичкой', 850.00, 30),
('I0015', 'Кофта с капибарой', 'C3', 'Крутая кофта с капибарой', 850.00, 30);*/
```

## Вставка в таблицу userType
```
INSERT INTO userType VALUES 
('UT01', 'Клиент'),
('UT02', 'Администратор'),
('UT03', 'Менеджер');
```

## Вставка в таблицу user
```
INSERT INTO "user" VALUES 
('U001', 'Иван', 'Иванов', 'Сергеевич', 'ivan.ivanov@example.com', '89001234567', 'UT01'),
('U002', 'Ольга', 'Смирнова', 'Викторовна', 'olga.smirnova@example.com', '89007654321', 'UT01'),
('U003', 'Анна', 'Кузнецова', 'Алексеевна', 'anna.kuznetsova@example.com', '89006543210', 'UT02'),
('U004', 'Дмитрий', 'Петров', 'Иванович', 'dmitriy.petrov@example.com', '89005432109', 'UT03'),
('U005', 'Мария', 'Васильева', 'Олеговна', 'maria.vasileva@example.com', '89004321098', 'UT03');
```

## Вставка в таблицу orderStatus
```
INSERT INTO orderStatus VALUES 
('OS01', 'Ожидает оплаты'),
('OS02', 'Подтвержден'),
('OS03', 'Доставлен'),
('OS04', 'Отменен'),
('OS05', 'Возврат средств');
```

## Вставка в таблицу delivery
```
INSERT INTO delivery VALUES 
('D001', 'Стандартная доставка', 5.0, 300.00, '02:00:00'),
('D002', 'Экспресс-доставка', 10.0, 600.00, '01:00:00'),
('D003', 'Самовывоз', 0.0, 0.00, '00:30:00');
```

## Вставка в таблицу review
```
INSERT INTO review VALUES 
('R001', 5.0, 'Крутая игрушка!', '2024-11-01 10:00:00', FALSE, 'I001'),
('R002', 4.5, 'Крутая крушка в подарок!', '2024-11-02 12:00:00', TRUE, 'I007'),
('R003', 5.0, 'Мягкая и удобная кофта', '2024-11-03 15:30:00', FALSE, 'I0014'),
('R004', 4.0, 'Ребенок рад!', '2024-11-04 18:45:00', TRUE, 'I004'),
('R005', 5.0, 'Кружка топ!', '2024-11-05 20:00:00', FALSE, 'I007');
```

## Вставка в таблицу paymentType
```
INSERT INTO paymentType VALUES 
('PT01', 'Банковская карта'),
('PT02', 'Наличные'),
('PT03', 'СБП');
```

## Вставка в таблицу package
```
INSERT INTO package VALUES 
('PK001', 'Коробка стандартная', 'Картонная коробка для транспортировки', 100.00),
('PK002', 'Упаковка премиум', 'Премиум-упаковка с защитой', 500.00),
('PK003', 'Полиэтиленовый пакет', 'Простая упаковка', 20.00),
('PK004', 'Подарочная упаковка', 'Красочная подарочная упаковка', 700.00),
('PK005', 'Без упаковки', 'Доставка без упаковки', 0.00);
```

## Вставка в таблицу order
```
INSERT INTO "order" VALUES 
('O001', 'U001', '2024-11-05 10:00:00', 'D001', 'ул. Ленина, д. 1', 5.0, 300.00, 55300.00, 'OS02'),
('O002', 'U002', '2024-11-05 12:00:00', 'D002', 'ул. Советская, д. 15', 10.0, 600.00, 45000.00, 'OS02'),
('O003', 'U003', '2024-11-05 14:00:00', 'D003', 'ул. Мира, д. 8', 0.0, 0.00, 14000.00, 'OS03'),
('O004', 'U004', '2024-11-05 16:00:00', 'D002', 'ул. Победы, д. 23', 7.0, 400.00, 300.00, 'OS03'),
('O005', 'U005', '2024-11-05 18:00:00', 'D001', 'ул. Гагарина, д. 42', 20.0, 1500.00, 9500.00, 'OS01');
```

## Вставка в таблицу orderItem
```
INSERT INTO orderItem VALUES 
('OI001', 'O001', 'I001', 1, 850.00, 'PK001', 1, 950.00),
('OI002', 'O002', 'I006', 1, 450.00, 'PK005', 1, 450.00),
('OI003', 'O003', 'I008', 1, 850.00, 'PK001', 1, 950.00),
('OI004', 'O004', 'I0012', 2, 1700.00, 'PK003', 1, 1720.00),
('OI005', 'O005', 'I0013', 1, 850.00, 'PK005', 1, 850.00);
```



