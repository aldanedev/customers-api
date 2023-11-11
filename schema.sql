CREATE TABLE cities (
  id UUID PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

INSERT INTO cities (id, name) VALUES ('ab90880f-694b-405d-bbf8-4578d99347c9', 'Chiclayo');
INSERT INTO cities (id, name) VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Lima');
INSERT INTO cities (id, name) VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Trujillo');

CREATE TABLE customers (
  dni VARCHAR(8) PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(9) NOT NULL,
  city_id UUID REFERENCES cities(id)
);


