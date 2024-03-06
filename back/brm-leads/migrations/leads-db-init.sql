CREATE TABLE leads_shard01 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER,
    "title" VARCHAR(200),
    "description" VARCHAR(1000),
    "price" INTEGER,
    "status" INTEGER,
    "responsible" INTEGER,
    "company_id" INTEGER,
    "client_company" INTEGER,
    "client_employee" INTEGER,
    "creation_date" DATE,
    "is_deleted" BOOLEAN
);

CREATE TABLE leads_shard02 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER,
    "title" VARCHAR(200),
    "description" VARCHAR(1000),
    "price" INTEGER,
    "status" INTEGER,
    "responsible" INTEGER,
    "company_id" INTEGER,
    "client_company" INTEGER,
    "client_employee" INTEGER,
    "creation_date" DATE,
    "is_deleted" BOOLEAN
);

CREATE TABLE leads_shard03 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER,
    "title" VARCHAR(200),
    "description" VARCHAR(1000),
    "price" INTEGER,
    "status" INTEGER,
    "responsible" INTEGER,
    "company_id" INTEGER,
    "client_company" INTEGER,
    "client_employee" INTEGER,
    "creation_date" DATE,
    "is_deleted" BOOLEAN
);

CREATE TABLE leads_shard04 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER,
    "title" VARCHAR(200),
    "description" VARCHAR(1000),
    "price" INTEGER,
    "status" INTEGER,
    "responsible" INTEGER,
    "company_id" INTEGER,
    "client_company" INTEGER,
    "client_employee" INTEGER,
    "creation_date" DATE,
    "is_deleted" BOOLEAN
);

CREATE TABLE status (
    "id" INTEGER PRIMARY KEY,
    "name" VARCHAR(100)
);

INSERT INTO status (id, name)
VALUES
    (1, 'Новая сделка'),
    (3, 'Установка контакта'),
    (4, 'Обсуждение деталей'),
    (5, 'Заключительные детали'),
    (6, 'Завершено'),
    (7, 'Отклонено');
