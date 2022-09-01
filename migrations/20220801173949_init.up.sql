-- countries --
CREATE TABLE countries
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- companies --
CREATE TABLE companies
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT                          NOT NULL,
    country    INT REFERENCES countries (id) NOT NULL,
    website    TEXT UNIQUE                   NOT NULL,
    phone      TEXT UNIQUE                   NOT NULL,
    created_at TIMESTAMPTZ                   NOT NULL,
    updated_at TIMESTAMPTZ                   NOT NULL,
    deleted_at TIMESTAMPTZ
)