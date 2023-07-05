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
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT user_info_pkey PRIMARY KEY (uid)
);

-- vote table
CREATE TABLE IF NOT EXISTS vote
(
    vid uuid DEFAULT uuid_generate_v4(),
    uid uuid NOT NULL,
    item_name TEXT NOT NULL UNIQUE,
    item_description TEXT,
    vote_count INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT vote_pkey PRIMARY KEY (vid)
);