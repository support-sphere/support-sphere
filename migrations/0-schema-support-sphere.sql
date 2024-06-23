 CREATE SEQUENCE tickets_id_seq
     START 1000
     INCREMENT BY 1000;
CREATE SEQUENCE
 CREATE TABLE tickets (
     id BIGINT PRIMARY KEY DEFAULT nextval('tickets_id_seq'),
     title VARCHAR(255) NOT NULL,
     description TEXT,
     status VARCHAR(50) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     created_by VARCHAR(255)
 );