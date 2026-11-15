-- +goose Up
CREATE TABLE hooks (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE requests (
    id SERIAL PRIMARY KEY,
    hook_id UUID NOT NULL REFERENCES hooks(id),
    method TEXT NOT NULL,
    path TEXT NOT NULL,
    headers JSONB NOT NULL,
    body TEXT NOT NULL,
    ip TEXT NOT NULL,
    received_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_requests_hook_id ON requests(hook_id);

-- +goose Down
DROP TABLE requests;
DROP TABLE hooks;