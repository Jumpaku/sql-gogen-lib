CREATE TABLE A (
    PK INT64 NOT NULL,
    Col_01 BOOL,
    Col_02 BOOL NOT NULL,
    Col_03 BYTES(50),
    Col_04 BYTES(50) NOT NULL,
    Col_05 DATE,
    Col_06 DATE NOT NULL,
    Col_07 FLOAT64,
    Col_08 FLOAT64 NOT NULL,
    Col_09 INT64,
    Col_10 INT64 NOT NULL,
    Col_11 JSON,
    Col_12 JSON NOT NULL,
    Col_13 NUMERIC,
    Col_14 NUMERIC NOT NULL,
    Col_15 STRING(50),
    Col_16 STRING(50) NOT NULL,
    Col_17 TIMESTAMP,
    Col_18 TIMESTAMP NOT NULL,
) PRIMARY KEY (PK);