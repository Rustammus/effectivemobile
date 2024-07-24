CREATE TABLE peoples
(
    uuid   uuid DEFAULT GEN_RANDOM_UUID() NOT NULL PRIMARY KEY ,
    passport_serie   INT NOT NULL,
    passport_number  INT NOT NULL,
    surname  VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100) NULL,
    address VARCHAR(100) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP(0)
);

DROP TABLE peoples