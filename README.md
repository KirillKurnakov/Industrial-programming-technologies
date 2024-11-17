Технологии индустриального программирования (магистратура)
Курнаков Кирилл Александрович ЭФМО-02-24

## Практическая работа 9
## Аутентификация и авторизация в REST API.  
Ссылка на код:  

### 9.1. Получение токена  
POST http://localhost:8080/login  
![image](https://github.com/user-attachments/assets/1e334fb6-1b85-4587-8056-35609ffce901)


### 9.2. Получение продуктов из корзины  авторизованным пользователем
Корзина пустая, так как пользователь не добавил товары в корзину
GET http://localhost:8080/cart  
![image](https://github.com/user-attachments/assets/e8fbdad0-9270-4e1f-9abb-a19dc78c321d)

### 9.3. Получение продуктов из корзины неавторизованным пользователем  
GET http://localhost:8080/cart  
![image](https://github.com/user-attachments/assets/1702580f-0892-47c0-b7f7-643f8f394ef3)

### 9.4. Получение нового токена  
POST http://localhost:8080/login  
![image](https://github.com/user-attachments/assets/279b5bf6-5a61-4c18-aa8e-4013a60f1f64)

### 9.5. Проверка срока действия токена  
GET http://localhost:8080/cart  
![image](https://github.com/user-attachments/assets/f9f06820-e02d-4bd5-bb9f-19b88b1749fe)

### 9.6. Запрос рефреша токена  
POST http://localhost:8080/refresh  
![image](https://github.com/user-attachments/assets/484e7bf8-4584-4b00-a799-41b8a665bd79)

### 9.7. Проверка нового токена
GET http://localhost:8080/cart  
![image](https://github.com/user-attachments/assets/48a825a6-296c-40f0-8e3f-427c4ba6c64e)
  
