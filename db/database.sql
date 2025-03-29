CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "Users"(
    "id" VARCHAR PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "users_id_email_name_index" ON "Users"("id", "email", "name");

CREATE TABLE "Products"(
    "id" VARCHAR PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "code" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "products_id_name_index" ON "Products"("id", "name");

CREATE TABLE "Orders"(
     "id" VARCHAR PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
     "code" VARCHAR(255) NOT NULL,
     "user_id" UUID NOT NULL,
     "total_price" INTEGER NOT NULL,
     "status" VARCHAR(255) CHECK
         ("status" IN('new', 'progress', 'done', 'cancelled')) NOT NULL DEFAULT 'new',
     "created_at" DATE NOT NULL,
     "updated_at" DATE NOT NULL,
     "deleted_at" DATE NULL
);

CREATE TABLE "OrderLines"(
     "id" VARCHAR PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
     "product_id" UUID NOT NULL,
     "order_id" UUID NOT NULL,
     "quantity" SMALLINT NOT NULL,
     "price" INTEGER NOT NULL,
     "created_at" DATE NOT NULL,
     "updated_at" DATE NOT NULL,
     "deleted_at" DATE NULL
);

CREATE TABLE "Carts"(
    "id" VARCHAR PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "user_id" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE NULL
);

CREATE TABLE "CartLines"(
    "id" VARCHAR PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "cart_id" UUID NOT NULL,
    "product_id" UUID NOT NULL,
    "quantity" SMALLINT NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE NULL
);

ALTER TABLE "OrderLines" ADD CONSTRAINT "orderlines_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "Orders"("id");
ALTER TABLE "OrderLines" ADD CONSTRAINT "orderlines_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "Products"("id");
ALTER TABLE "CartLines" ADD CONSTRAINT "cartlines_cart_id_foreign" FOREIGN KEY("cart_id") REFERENCES "Carts"("id");
ALTER TABLE "CartLines" ADD CONSTRAINT "cartlines_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "Products"("id");
ALTER TABLE "Orders" ADD CONSTRAINT "orders_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "Users"("id");
ALTER TABLE "Carts" ADD CONSTRAINT "carts_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "Users"("id");