create table task (
	id UUID, 
	method varchar,
	url varchar,
	headers map<varchar, varchar>,
	http_status_code int,
	task_status varchar,
	response_length int,
	request_body text,
	response_body text,
PRIMARY KEY(id));
