CREATE TABLE ads (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER,
    "title" VARCHAR(200),
    "text" VARCHAR(1000),
    "industry" INTEGER,
    "price" INTEGER,
    "creation_date" DATE,
    "created_by" INTEGER,
    "responsible" INTEGER,
    "is_deleted" BOOLEAN
);

CREATE TABLE responses_shard01 (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER,
    "employee_id" INTEGER,
    "ad_id" INTEGER,
    "creation_date" DATE
);

CREATE TABLE responses_shard02 (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER,
    "employee_id" INTEGER,
    "ad_id" INTEGER,
    "creation_date" DATE
);

CREATE TABLE responses_shard03 (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER,
    "employee_id" INTEGER,
    "ad_id" INTEGER,
    "creation_date" DATE
);

CREATE TABLE responses_shard04 (
    "id" SERIAL PRIMARY KEY,
    "company_id" INTEGER,
    "employee_id" INTEGER,
    "ad_id" INTEGER,
    "creation_date" DATE
);
