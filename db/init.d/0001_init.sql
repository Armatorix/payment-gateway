CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE merchants
(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  name VARCHAR NOT NULL,
  api_secret VARCHAR NOT NULL
);
