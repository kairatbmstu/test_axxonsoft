create table headers (
    id uuid PRIMARY KEY,
    request_headers_task_id uuid  REFERENCES task(id),
	response_headers_task_id uuid  REFERENCES task(id),
	header_name varchar(8192),
	header_value text
)