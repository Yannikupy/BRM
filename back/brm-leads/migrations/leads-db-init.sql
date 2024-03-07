CREATE TABLE leads_shard01 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "description" VARCHAR(1000),
    "price" INTEGER NOT NULL,
    "status" INTEGER NOT NULL,
    "responsible" INTEGER NOT NULL,
    "company_id" INTEGER NOT NULL,
    "client_company" INTEGER NOT NULL,
    "client_employee" INTEGER NOT NULL,
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE leads_shard02 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "description" VARCHAR(1000),
    "price" INTEGER NOT NULL,
    "status" INTEGER NOT NULL,
    "responsible" INTEGER NOT NULL,
    "company_id" INTEGER NOT NULL,
    "client_company" INTEGER NOT NULL,
    "client_employee" INTEGER NOT NULL,
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE leads_shard03 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "description" VARCHAR(1000),
    "price" INTEGER NOT NULL,
    "status" INTEGER NOT NULL,
    "responsible" INTEGER NOT NULL,
    "company_id" INTEGER NOT NULL,
    "client_company" INTEGER NOT NULL,
    "client_employee" INTEGER NOT NULL,
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE leads_shard04 (
    "id" SERIAL PRIMARY KEY,
    "ad_id" INTEGER NOT NULL,
    "title" VARCHAR(200) NOT NULL,
    "description" VARCHAR(1000),
    "price" INTEGER NOT NULL,
    "status" INTEGER NOT NULL,
    "responsible" INTEGER NOT NULL,
    "company_id" INTEGER NOT NULL,
    "client_company" INTEGER NOT NULL,
    "client_employee" INTEGER NOT NULL,
    "creation_date" DATE NOT NULL,
    "is_deleted" BOOLEAN NOT NULL
);

CREATE TABLE statuses (
    "id" INTEGER PRIMARY KEY,
    "name" VARCHAR(100) PRIMARY KEY
);

INSERT INTO status (id, name)
VALUES
    (1, 'Новая сделка'),
    (3, 'Установка контакта'),
    (4, 'Обсуждение деталей'),
    (5, 'Заключительные детали'),
    (6, 'Завершено'),
    (7, 'Отклонено');
