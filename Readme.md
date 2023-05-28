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