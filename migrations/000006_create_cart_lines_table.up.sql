CREATE TABLE public.cart_lines (
    id text NOT NULL,
    cart_id text,
    product_id text,
    quantity bigint,
    price numeric,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);