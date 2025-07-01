CREATE TABLE public.products (
    id text NOT NULL,
    code text,
    name text,
    image_url text,
    description text,
    price numeric,
    active boolean DEFAULT true,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
