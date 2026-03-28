CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    status TEXT,
    created_at DATE,
    updated_at DATE
);

CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);