
CREATE DATABASE log;


CREATE TABLE IF NOT EXISTS test_log (
    msec                     	UInt32,
    time_iso8601            	Datetime,
    status                   	UInt32,
    request_method           	String,
    request_uri              	String,
    remote_addr              	String,
    remote_port              	UInt32,
    server_addr              	String,
    server_port              	UInt32,
    http_host                 	String,
    http_name                 	String,
    args                     	String,
    request_length             	UInt32,
    body_bytes_sent         	UInt32,
    bytes_sent                  UInt32,
    request_time             	Float32,
    connection_requests     	UInt32,
    connection                 	UInt32,
    upstream                    String,
    upstream_status             UInt32,
    upstream_response_time      UInt32,
    upstream_response_length	UInt32,
    http_referer                String,
    http_user_agent             String,
    pid                         UInt32,
    ssl_protocol                String,
    scheme                      String,
    server_protocol             String
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(time_iso8601)
PRIMARY KEY time_iso8601
ORDER BY(time_iso8601)
TTL time_iso8601 + INTERVAL 72 HOUR;

