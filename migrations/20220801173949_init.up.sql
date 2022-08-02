-- countries --
CREATE TABLE countries
(
    id   SERIAL PRIMARY KEY,
    name TEXT UNIQIE NOT NULL
);

-- seed some base countries --
INSERT INTO countries (name)
VALUES ('Greece'),
       ('Nepal'),
       ('Syria'),
       ('France'),
       ('Japan'),
       ('Cyprus'),
       ('Germany');

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