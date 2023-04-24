CREATE TABLE IF NOT EXISTS UGSN
(
    ID       UUID PRIMARY KEY,
    TITLE    VARCHAR(255)                                 NOT NULL,
    CODE     VARCHAR(8)                                   NOT NULL,
    LEVEL_ID UUID REFERENCES LEVEL (ID) ON DELETE CASCADE NOT NULL
);