1. You need to have rabbitmq and postgresql installed on your computer
2. create schema with commands:

```
sudo su - postgres
psql
```

After this command you will be in console of postgres, where you should perform next commands:

```
1) create database test_axxonsoft;
2) create user test_axxonsoft with password '123456';
3) grant all privileges on database test_axxonsoft to test_axxonsoft;
```

3. After that run the project the migration scripts will be performed automatically on schema

4. Example of request: 

```
POST /task HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 232

{
    "method":"GET",
    "url":"https://tengrinews.kz/world_news/premer-ministr-italii-djordja-meloni-otmenila-vizit-v-astanu-499621/",
    "requestHeaders": {
        "Authorization":"123123",
        "X-Token":"123123123"
    }
}
```

```
GET /task/50cdbf4e-8177-4bdb-9605-2c8e992eac1a HTTP/1.1
Host: localhost:8080
```