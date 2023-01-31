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
> "location": "Country, Region, City", // Must be only in this format, else skip all that after unknown
> "password": "MyPassword"
> }
> ```
> ### Response:
> ```json lines
> {
> "access-code": "123jsdfk13jk1243jk123123tsgsdg",
> "status": "registered"
> }
```
