CREATE TABLE if not exists owners (
    id VARCHAR(36) NOT NULL,
    access_key VARCHAR(200) NOT NULL,
    app_name VARCHAR(200) NOT NULL,

    PRIMARY KEY (id)
);
