## sap2gRPCdirx
## On Russian

Webserver-parser модуль rest-запросов.
Формирует и передает данные MSSQL-модулю для копирования таблиц между БД.
Передает в URL-строке имя БД-донора и имя БД-реципиента.
Представлены тестовые модули для обмена данными.

Схема обмена данными:

> REST(imitation from web browser) <---> web-server_URL-parser/rest2mssql/mssqldsn <---> MSSQL/mssqlinsert

Для проверки запустить сервер rest2mssql из строки браузера и создать запрос типа:

	http://localhost:8080/DirectumRX.dbo.Integration_Data:DirectumRXTest.dbo.Integration_Data

> где
> 
> DirectumRX.dbo.Integration_Data -- имя БД донора
> 
> DirectumRXTest.dbo.Integration_Data -- имя БД реципиента

## On English

Webserver-parser module of rest-requests.
You can shape and send data to MSSQL-module for copy midle tables.
Send to URL name table-donor and name table-recipient.
We can see tests modules for exchange data.

Communication scheme:

> REST(imitation from web browser) <---> web-server_URL-parser/rest2mssql/mssqldsn <---> MSSQL/mssqlinsert

To check, start the rest2mssql server from the browser line and create a query like:

	http://localhost:8080/DirectumRX.dbo.Integration_Data:DirectumRXTest.dbo.Integration_Data
	
where

> DirectumRX.dbo.Integration_Data -- donor database name
> 
> DirectumRXTest.dbo.Integration_Data -- recipient database name
