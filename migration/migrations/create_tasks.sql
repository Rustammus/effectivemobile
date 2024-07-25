CREATE TABLE tasks
(
    uuid   uuid DEFAULT GEN_RANDOM_UUID() NOT NULL PRIMARY KEY ,
    people_uuid uuid NOT NULL ,
    name TEXT NULL ,
    start_time timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    end_time timestamptz NULL ,
    FOREIGN KEY (people_uuid) REFERENCES peoples(uuid) ON DELETE CASCADE
);

DROP TABLE tasks