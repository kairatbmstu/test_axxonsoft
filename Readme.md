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
Content-Length: 225

{
    "method":"GET",
    "url":"https://tengrinews.kz/world_news/premer-ministr-italii-djordja-meloni-otmenila-vizit-v-astanu-499621/",
    "headers": {
        "Authorization":"123123",
        "X-Token":"123123123"
    }
}

Response:
{
    "id": "4df5c24f-9906-437f-b242-90afd7fa52b5"
}
```

```
GET /task/4df5c24f-9906-437f-b242-90afd7fa52b5 HTTP/1.1
Host: localhost:8080

Response:

{
    "id": "4df5c24f-9906-437f-b242-90afd7fa52b5",
    "method": "GET",
    "url": "https://tengrinews.kz/world_news/premer-ministr-italii-djordja-meloni-otmenila-vizit-v-astanu-499621/",
    "httpStatusCode": 200,
    "taskStatus": "done",
    "length": 121887,
    "headers": {
        "Cache-Control": "private, must-revalidate",
        "Content-Type": "text/html; charset=UTF-8",
        "Date": "Sun, 28 May 2023 17:04:16 GMT",
        "Expires": "-1",
        "Pragma": "no-cache",
        "Server": "nginx",
        "Set-Cookie": "XSRF-TOKEN=eyJpdiI6InlVSnRsWGx3aUMzakFBUnk5bytUa0E9PSIsInZhbHVlIjoiU0pUM0JDZ29CNWc1Yk9VcEVHNmtna1VzMExEYVprZThMQ0hWM3J1TTJLNkM3UzhOaDNlTVJZQTY1V0RITHJPOHB4R1RlSFpIaU1pdERTd1hyZmJvWGk4cVJNNktzZ3JXUlRZSVRtbkpXODc1MTRLS1kwRlZFc1VMbVpsWFYzckYiLCJtYWMiOiIxMjM2ODgzYzM2MTViMjUyMGQ5NDdiZTQyM2Y3NDQ4NTI0NmUyOWFhMWJmODk2ZTYxZjQzZjFmMWU3MDUwNzM4In0%3D; expires=Mon, 29-May-2023 13:04:16 GMT; Max-Age=72000; path=/",
        "Strict-Transport-Security": "max-age=15768000"
    }
}
```

5. Example of POST request:

```
POST /task HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 446

{
    "method":"POST",
    "url":"https://64738023d784bccb4a3caaf7.mockapi.io/api/something",
    "headers": {
        "Authorization":"123123",
        "X-Token":"123123123"
    },
    "body": "{\r\n    \"createdAt\": \"2023-05-28T03:50:56.827Z\",\r\n    \"name\": \"Georgia Schulist Sr.\",\r\n    \"avatar\": \"https:\/\/cloudflare-ipfs.com\/ipfs\/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye\/avatar\/419.jpg\",\r\n    \"id\": \"20\"\r\n}"
}

Response:
{
    "id": "7028d975-d531-411a-93ad-e473758dd06c"
}
```

```
GET /task/7028d975-d531-411a-93ad-e473758dd06c HTTP/1.1
Host: localhost:8080

Response: 
{
    "id": "7028d975-d531-411a-93ad-e473758dd06c",
    "method": "POST",
    "url": "https://64738023d784bccb4a3caaf7.mockapi.io/api/something",
    "httpStatusCode": 201,
    "taskStatus": "done",
    "length": 187,
    "headers": {
        "Access-Control-Allow-Headers": "X-Requested-With,Content-Type,Cache-Control,access_token",
        "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,OPTIONS",
        "Access-Control-Allow-Origin": "*",
        "Connection": "keep-alive",
        "Content-Length": "187",
        "Content-Type": "application/json",
        "Date": "Sun, 28 May 2023 17:02:46 GMT",
        "Server": "Cowboy",
        "Vary": "Accept-Encoding",
        "Via": "1.1 vegur",
        "X-Powered-By": "Express"
    }
}

```