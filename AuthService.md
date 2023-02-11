## Auth Service:
### For Authorization:
>### ```POST: /auth/sign/in```
> ### Description:
> Before that you need make SignUp. 
> If success server return ```Access Code```
> for 1 hour and status ```Authorized``` \
> Else Server return empty ```Access Code``` and error in status like ```"wrong password"``` or ```"unregistered"```
> 
> ### Request:
> ```json lines
> {
> "email": "user@email.net",
> "password": "MyPassword"
> }
> ```
> ### Response:
> ```json lines
> {
> "access-code": "123jsdfk13jk1243jk",
> "status": "authorized"
> }
> ```

### For Registration:
>### ```POST: /auth/sign/up```
> ### Description:
> If success server return ```Access Code```
> for 1 hour and status ```Authorized``` \
> Else Server return empty ```Access Code``` and error in status like ```"broken data"``` or ```"email used before"```
> 
> ### Request:
> ```json lines
> {
> "username": "Сергей",        // Not bound be unique
> "date_birth": "12.12.2001",  // Must be in format: dd.mm.yyyy
> "email": "user@email.net",
> "location": {
>   "country": "Russia",
>   "region":" Krasnoyarskiy krai",
>   "city": "Krasnoyarsk"
>   },
> "password": "MyPassword"
> }
> ```
> ### Response:
> ```json lines
> {
> "access-code": "123jsk13jk1243jk123123tdg",
> "status": "registered"
> }
> ```
### Process of Recovering Password
> ### ```POST: /auth/recover```
> ### Description:
> After got request: server send code to Email from Request and returned you ActionUID
> ### Request:
> ```json lines
> {
> "email": "user@email.net",
> "location": "123.234.23.12"
> }
>```
> ### Response:
> ```json lines
> {
> "action-uid": "123kdf123",
> "status": "success, code sent",
> }
>```
> ### ```POST: /auth/recover/code```
> ### Description:
> After got request: server send code to Email from Request and returned you ActionUID
> ### Request:
> ```json lines
> {
> "action-uid": "123kdf123",
> "code": 124819
> }
>```
> ### Response:
> ```json lines
> {
> "status": "success",
> "access-code": "123l1afd2asd2ajl123532asd"
> }
>```
> ### ```POST: /auth/me```
> ### Description:
> After got request: server send code to Email from Request and returned you ActionUID
> ### Request:
> ```json lines
> {
> "access-code": "123ksd2fsf134123df123"
> }
>```
> ### Response:
> ```json lines
> {
> "status": "success",
> "username": "UserName",
> "email": "user@gmail.ru",
> "date-birth": "12.02.2000",
> "location": "Country, Region, City",
> "photos": [[],[],[],[]]
> }
>```