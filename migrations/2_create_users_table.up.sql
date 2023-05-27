CREATE TABLE if not exists users (
    id VARCHAR(36) NOT NULL,
    username VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    password_hash VARCHAR(200) NOT NULL,
    owner_id VARCHAR (36) NOT NULL,

    PRIMARY KEY (id),

    CONSTRAINT FK_owner_id FOREIGN KEY (owner_id)
    REFERENCES owners(id)
);
