CREATE TABLE public.carts (
    id text NOT NULL,
    user_id text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);