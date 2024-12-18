-- classDiagram
--     C_1 <|-- C_2
--     C_2 <|-- C_3
--     C_2 <|-- C_4
--     C_3 <|-- C_5
--     C_4 <|-- C_5
CREATE TABLE C_1 (
    PK_11 INT64 NOT NULL,
    PK_12 INT64 NOT NULL,
) PRIMARY KEY (PK_11, PK_12);

CREATE TABLE C_2 (
    PK_21 INT64 NOT NULL,
    PK_22 INT64 NOT NULL,
    CONSTRAINT FK_C_2_1 FOREIGN KEY (PK_21, PK_22) REFERENCES C_1 (PK_11, PK_12),
) PRIMARY KEY (PK_21, PK_22);

CREATE TABLE C_3 (
    PK_31 INT64 NOT NULL,
    PK_32 INT64 NOT NULL,
    CONSTRAINT FK_C_3_2 FOREIGN KEY (PK_31, PK_32) REFERENCES C_2 (PK_21, PK_22),
) PRIMARY KEY (PK_31, PK_32);

CREATE TABLE C_4 (
    PK_41 INT64 NOT NULL,
    PK_42 INT64 NOT NULL,
    CONSTRAINT FK_C_4_2 FOREIGN KEY (PK_41, PK_42) REFERENCES C_2 (PK_21, PK_22),
) PRIMARY KEY (PK_41, PK_42);

CREATE TABLE C_5 (
    PK_51 INT64 NOT NULL,
    PK_52 INT64 NOT NULL,
    CONSTRAINT FK_C_5_3 FOREIGN KEY (PK_51, PK_52) REFERENCES C_3 (PK_31, PK_32),
    CONSTRAINT FK_C_5_4 FOREIGN KEY (PK_51, PK_52) REFERENCES C_4 (PK_41, PK_42),
) PRIMARY KEY (PK_51, PK_52);