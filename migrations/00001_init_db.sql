-- +goose Up
CREATE TABLE IF NOT EXISTS messages (
                              id BIGSERIAL PRIMARY KEY,
                              from varchar NOT NULL,
                              to varchar NOT NULL,
                              text varchar,
                              datetime TIMESTAMP NOT NULL DEFAULT NOW(),
                              removed BOOLEAN DEFAULT FALSE,
                              created TIMESTAMP NOT NULL DEFAULT NOW(),
                              updated TIMESTAMP
);

CREATE TYPE event_type AS ENUM ('Created', 'Updated', 'Removed');

CREATE TYPE event_status AS ENUM ('lock', 'unlock');

CREATE TABLE messages_events (
                                    id BIGSERIAL PRIMARY KEY,
                                    message_id BIGINT NOT NULL REFERENCES messages(id),
                                    type event_type,
                                    status event_status,
                                    payload JSONB,
                                    updated TIMESTAMP
);

CREATE INDEX messages_index ON messages(id);
CREATE INDEX messages_events_index ON messages_events(id);

-- +goose Down
DROP TABLE messages_events;
DROP TYPE event_status;
DROP TYPE event_type;
DROP TABLE messages;