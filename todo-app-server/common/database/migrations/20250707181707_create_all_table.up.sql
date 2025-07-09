CREATE TYPE tag_enum as ENUM ('inactive', 'active');

CREATE TABLE users
(
    id         INTEGER GENERATED ALWAYS AS IDENTITY,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(100) NOT NULL,
    password   TEXT         NOT NULL,
    created_at TIMESTAMP with time zone DEFAULT now(),
    updated_at TIMESTAMP with time zone DEFAULT now(),
    PRIMARY KEY (id),
    CONSTRAINT email_unique UNIQUE (email)
);

CREATE TABLE todos
(
    id          INTEGER GENERATED ALWAYS AS IDENTITY,
    user_id     INTEGER      NOT NULL,
    title       VARCHAR(100) NOT NULL,
    description TEXT         NOT NULL,
    tag         tag_enum     NOT NULL,
    due_date    TIMESTAMPTZ  NOT NULL,
    created_at  TIMESTAMP with time zone DEFAULT now(),
    updated_at  TIMESTAMP with time zone DEFAULT now(),
    PRIMARY KEY (id),
    CONSTRAINT fk_user_todos FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_user_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER trigger_update_todo_updated_at
    BEFORE UPDATE ON todos
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE INDEX idx_user_todos_id ON todos(user_id);
CREATE INDEX idx_todos_tag_id ON todos(tag);
CREATE INDEX idx_todos_title_id ON todos(title);