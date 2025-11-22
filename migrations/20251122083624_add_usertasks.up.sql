ALTER TABLE tasks
    ADD COLUMN user_id INTEGER REFERENCES users(id) ON DELETE CASCADE;

-- Только индекс отдельно
CREATE INDEX idx_tasks_user_id ON tasks(user_id);