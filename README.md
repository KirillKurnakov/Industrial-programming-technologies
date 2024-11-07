Технологии индустриального программирования
Курнаков Кирилл Александрович ЭФМО-02-24

# Практическая работа 8

## CRUD для товаров в Интернет-магазине подарков и сувениров

### Получение всех товаров

GET http: //localhost:8080/items

![image](https://github.com/user-attachments/assets/3289731c-6eb7-4045-ad4c-f3baf67cd8af)

### Получение товара по ID

GET http: //localhost:8080/items/2

![image](https://github.com/user-attachments/assets/b9032c0e-10d9-4882-91aa-fee3db4c7c3a)

### Обновление существующего товара

PUT http: //localhost:8080/items/2

Отправка PUT запроса
![image](https://github.com/user-attachments/assets/b26786c0-1390-44c6-8a4f-3cdd8365d545)

Отправка GET запроса для проверки
![image](https://github.com/user-attachments/assets/07a6061d-225f-43d2-8730-45dfbf26fa8d)

### Удаление товара

DELETE http: //localhost:8080/items/2

Отправка DELETE запроса
![image](https://github.com/user-attachments/assets/a1aad60f-ad4f-4b0f-9bce-1223b9fc1668)

Отправка GET запроса для проверки
![image](https://github.com/user-attachments/assets/03ca413b-8432-42cc-bf9a-c8c0703f5665)

### Создание нового товара

POST http://localhost:8080/items

Отправка POST запроса
![image](https://github.com/user-attachments/assets/fad901c9-de8c-4e7b-a219-67a562b92cf5)

Отправка GET запроса для проверки
![image](https://github.com/user-attachments/assets/c1103245-f907-4f55-ab32-f9687055e0d1)

## Корзина
### Получение всех товаров из корзины

GET http: //localhost:8080/id/1/basket

![image](https://github.com/user-attachments/assets/d56df080-dcc0-4392-9150-8d3673408600)

### Добавление товара в корзину

POST http: //localhost:8080/id/1/basket

Отправка POST запроса
![image](https://github.com/user-attachments/assets/720676d3-1b8e-468d-a780-ed28f96985cc)

Отправка GET запроса для проверки
![image](https://github.com/user-attachments/assets/72f1e6d6-10fd-4093-9e9b-143c279dbe34)

### Удаление товара из корзины

Отправка GET запроса для начального состояния
![image](https://github.com/user-attachments/assets/420e3607-046a-45bc-9ef1-703c9c3fb100)


Отправка DELETE запроса
![image](https://github.com/user-attachments/assets/e85b587f-e154-433f-a75e-2472d57a6937)


Отправка GET запроса для проверки
![image](https://github.com/user-attachments/assets/5b0490bd-5a00-40b5-8da1-9eba4f99b513)
