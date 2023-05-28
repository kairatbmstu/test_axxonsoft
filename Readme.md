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

4. Example of GET request: 

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

5. Example of POST request:

```
POST /task HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 460

{
    "method":"POST",
    "url":"https://64738023d784bccb4a3caaf7.mockapi.io/api/something",
    "requestHeaders": {
        "Authorization":"123123",
        "X-Token":"123123123"
    },
    "requestBody": "{\r\n    \"createdAt\": \"2023-05-28T03:50:56.827Z\",\r\n    \"name\": \"Georgia Schulist Sr.\",\r\n    \"avatar\": \"https:\/\/cloudflare-ipfs.com\/ipfs\/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye\/avatar\/419.jpg\",\r\n    \"id\": \"20\"\r\n}"
}

Response:
{
    "id": "6586b269-a344-4716-82f6-95daaf078aad"
}
```

```
GET /task/6586b269-a344-4716-82f6-95daaf078aad HTTP/1.1
Host: localhost:8080

Response: 
{
    "id": "6586b269-a344-4716-82f6-95daaf078aad",
    "method": "POST",
    "url": "https://64738023d784bccb4a3caaf7.mockapi.io/api/something",
    "httpStatusCode": 201,
    "taskStatus": "done",
    "length": 181,
    "headers": {
        "Access-Control-Allow-Headers": "X-Requested-With,Content-Type,Cache-Control,access_token",
        "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,OPTIONS",
        "Access-Control-Allow-Origin": "*",
        "Connection": "keep-alive",
        "Content-Length": "181",
        "Content-Type": "application/json",
        "Date": "Sun, 28 May 2023 16:53:52 GMT",
        "Server": "Cowboy",
        "Vary": "Accept-Encoding",
        "Via": "1.1 vegur",
        "X-Powered-By": "Express"
    }
}

```