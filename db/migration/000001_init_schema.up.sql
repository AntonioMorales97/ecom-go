CREATE TABLE "product_type" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description_long" text,
  "description_short" varchar,
  "price" int NOT NULL,
  "product_type_id" bigint NOT NULL,
  "product_category_id" bigint,
  "product_inventory_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_inventory" (
  "id" bigserial PRIMARY KEY,
  "quantity" int NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_order" (
  "id" bigserial PRIMARY KEY,
  "quantity" int NOT NULL,
  "product_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "product" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_type" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("product_category_id") REFERENCES "product_category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("product_inventory_id") REFERENCES "product_inventory" ("id");

ALTER TABLE "product_order" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_inventory" ADD CONSTRAINT "CK_product_inventory_quantity" CHECK ("quantity" >= 0);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON product_type
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON product_category
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON product
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON product_inventory
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON product_order
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();