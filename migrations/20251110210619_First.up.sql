CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     completed BOOLEAN NOT NULL DEFAULT FALSE,
                                     created_at TIMESTAMP WITH TIME ZONE,
                                     updated_at TIMESTAMP WITH TIME ZONE
);
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     email TEXT NOT NULL,
                                     password TEXT NOT NULL,
                                     created_at TIMESTAMP WITH TIME ZONE,
                                     updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);