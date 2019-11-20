-- Exported from QuickDBD: https://www.quickdatabasediagrams.com/
-- NOTE! If you have used non-SQL datatypes in your design, you will have to change these here.

CREATE TABLE housing (
    "housing_id" VARCHAR  NOT NULL PRIMARY KEY,
    "type" INTEGER,
    "surface_area" INTEGER,
    "rooms" INTEGER,
    "heating_system" VARCHAR,
    "year" INTEGER,
    "street_number" VARCHAR,
    "street" VARCHAR,
    "postcode" VARCHAR,
    "city" VARCHAR
);

CREATE TABLE consumption (
    "consumption_id" INTEGER,
    "housing_id" VARCHAR,
    "power_kw" VARCHAR,
    "date" date,
    CONSTRAINT fk_consumption_housing_id FOREIGN KEY (housing_id) REFERENCES housing(housing_id)
);

CREATE TABLE access (
    "access_id" INTEGER NOT NULL PRIMARY KEY,
    "housing_id" VARCHAR,
    "login" VARCHAR,
    "password" VARCHAR,
    "is_admin" BOOLEAN,
    CONSTRAINT fk_credentials_housing_id FOREIGN KEY (housing_id) REFERENCES housing(housing_id)
);



CREATE TABLE landlord (
    "landlord_id" INTEGER  NOT NULL PRIMARY KEY,
    "housing_id" VARCHAR,
    "lastname" VARCHAR,
    "firstname" VARCHAR,
    "company" VARCHAR,
    "address" VARCHAR,
    CONSTRAINT fk_landlord_housing_id FOREIGN KEY (housing_id) REFERENCES housing(housing_id)
);

CREATE TABLE tenant (
    "tenant_id" INTEGER   NOT NULL PRIMARY KEY,
    "housing_id" VARCHAR,
    "firstname" VARCHAR,
    "lastname" VARCHAR,
    CONSTRAINT fk_tenant_housing_id FOREIGN KEY (housing_id) REFERENCES housing(housing_id)
);
