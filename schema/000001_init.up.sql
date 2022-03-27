BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "languages" (
    "id" uuid DEFAULT uuid_generate_v4(),
    "title" varchar(255) NOT NULL,
    "rating" int NOT NULL,
    "developer" varchar(255) NOT NULL,
    "date_of_creation" int NOT NULL,
    PRIMARY KEY ("id")
);

COMMIT;

INSERT INTO
    languages(title, rating, developer, date_of_creation)
VALUES
    ('JavaScript', 1, 'Brendan Eich', 1995),
    (
        'Python',
        2,
        'Python Software Foundation, Guido van Rossum',
        1991
    );