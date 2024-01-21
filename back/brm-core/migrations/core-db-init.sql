CREATE TABLE companies IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'name' VARCHAR(100),
    'description' VARCHAR(1000),
    'industry' INTEGER,
    'owner_id' INTEGER,
    'rating' FLOAT,
    'creation_date' DATE,
    'is_deleted' BOOLEAN,
    'deletion_date' DATE,
);

CREATE TABLE industries IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'name' VARCHAR(100),
);

CREATE TABLE employees IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'company_id' INTEGER,
    'first_name' VARCHAR(100),
    'second_name' VARCHAR(100),
    'email' VARCHAR(100),
    'job_title' VARCHAR(100),
    'department' VARCHAR(100),
    'creation_date' DATE,
    'is_deleted' BOOLEAN,
    'deletion_date' DATE,
)

CREATE TABLE contact_shard01 IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'owner_id' INTEGER,
    'employee_id' INTEGER,
    'notes' VARCHAR(500),
    'creation_date' DATE,
    'is_deleted' BOOLEAN,
    'deletion_date' DATE,
)

CREATE TABLE contact_shard02 IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'owner_id' INTEGER,
    'employee_id' INTEGER,
    'notes' VARCHAR(500),
    'creation_date' DATE,
    'is_deleted' BOOLEAN,
    'deletion_date' DATE,
)

CREATE TABLE contact_shard03 IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'owner_id' INTEGER,
    'employee_id' INTEGER,
    'notes' VARCHAR(500),
    'creation_date' DATE,
    'is_deleted' BOOLEAN,
    'deletion_date' DATE,
)

CREATE TABLE contact_shard04 IF NOT EXIST (
    'id' SERIAL PRIMARY KEY,
    'owner_id' INTEGER,
    'employee_id' INTEGER,
    'notes' VARCHAR(500),
    'creation_date' DATE,
    'is_deleted' BOOLEAN,
    'deletion_date' DATE,
)
