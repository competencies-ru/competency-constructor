CREATE TABLE IF NOT EXISTS INDICATOR
(
    ID            uuid PRIMARY KEY,
    TITLE         VARCHAR(255)                                      NOT NULL,
    CODE          VARCHAR(255)                                      NOT NULL,
    SUBJECT_ID    uuid REFERENCES SUBJECT (ID)                      NOT NULL,
    COMPETENCY_ID uuid REFERENCES COMPETENCY (ID) ON DELETE cascade not null
);