CREATE TABLE peoples
(
    uuid                UUID NOT NULL PRIMARY KEY DEFAULT GEN_RANDOM_UUID() ,
    passport_serie      INT NOT NULL,
    passport_number     INT NOT NULL,
    surname             TEXT NOT NULL,
    name                TEXT NOT NULL,
    patronymic          TEXT NULL,
    address             TEXT NOT NULL,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP(0)
);
