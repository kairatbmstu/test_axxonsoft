CREATE TABLE task (
	id uuid not null PRIMARY KEY,
	method varchar(10) not null,
	url varchar(2048) NULL,
	http_status_code int not null,
	task_status varchar(50),
	response_length int not null,
	request_body text,
	response_body text
);
