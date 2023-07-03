-- create uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- encrypt password
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- user info table
CREATE TABLE IF NOT EXISTS user_info
(
    uid uuid DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT user_info_pkey PRIMARY KEY (uid)
);