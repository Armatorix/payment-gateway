CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE merchants
(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  name VARCHAR NOT NULL,
  api_secret VARCHAR NOT NULL
);

CREATE TYPE CARD_STATE AS ENUM('active', 'voided', 'refunded');

CREATE TABLE credit_cards
(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  holder VARCHAR NOT NULL,
  card_number VARCHAR NOT NULL,
  cvv integer NOT NULL,
  expiry_month integer NOT NULL,
  expiry_year integer NOT NULL,
  amount integer NOT NULL,
  currency VARCHAR,
  card_state CARD_STATE,

  merchant_id UUID,
  CONSTRAINT fk_merchant
    FOREIGN KEY(merchant_id)
      REFERENCES merchants(id)
);

