CREATE TABLE public.orders (
    id text NOT NULL,
    code text,
    user_id text,
    total_price numeric,
    status text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
