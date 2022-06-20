-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
	id UUID NOT NULL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	start_at TIMESTAMP NOT NULL DEFAULT NOW(),
	finish_at TIMESTAMP NOT NULL DEFAULT NOW(),
	user_id UUID NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE calendar;
-- +goose StatementEnd
