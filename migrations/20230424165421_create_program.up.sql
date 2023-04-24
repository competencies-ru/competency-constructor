CREATE TABLE IF NOT EXISTS PROGRAM
(
    ID           UUID PRIMARY KEY,
    TITLE        VARCHAR(255)                                     NOT NULL,
    CODE         VARCHAR(11)                                      NOT NULL,
    SPECIALTY_ID UUID REFERENCES SPECIALTY (ID) ON DELETE CASCADE NOT NULL
);