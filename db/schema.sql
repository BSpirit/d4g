-- Exported from QuickDBD: https://www.quickdatabasediagrams.com/
-- NOTE! If you have used non-SQL datatypes in your design, you will have to change these here.

CREATE TABLE "consumption" (
    "consumption_id" INTEGER   ,
    "housing_id" VARCHAR   ,
    "power_kw" VARCHAR   ,
    "date" date   ,
);

CREATE TABLE "credentials" (
    "credentials_id" INTEGER NOT NULL PRIMARY KEY,
    "housing_id" VARCHAR   ,
    "login" VARCHAR   ,
    "password" VARCHAR   ,
    "is_admin" BOOLEAN   ,
);

CREATE TABLE "housing" (
    "housing_id" VARCHAR  NOT NULL PRIMARY KEY  ,
    "type" INTEGER   ,
    "surface_area" INTEGER   ,
    "rooms" INTEGER   ,
    "heating_system" VARCHAR   ,
    "year" INTEGER   ,
    "street_number" VARCHAR   ,
    "street" VARCHAR   ,
    "postcode" VARCHAR   ,
    "city" VARCHAR   ,
);

CREATE TABLE "landlord" (
    "landlord_id" INTEGER  NOT NULL PRIMARY KEY  ,
    "housing_id" VARCHAR   ,
    "lastname" VARCHAR   ,
    "firstname" VARCHAR   ,
    "company" VARCHAR   ,
    "address" VARCHAR   ,
);

CREATE TABLE "tenant" (
    "tenant_id" INTEGER   NOT NULL PRIMARY KEY ,
    "housing_id" VARCHAR   ,
    "firstname" VARCHAR   ,
    "lastname" VARCHAR   ,
);

ALTER TABLE "consumption" ADD CONSTRAINT "fk_consumption_housing_id" FOREIGN KEY("housing_id")
REFERENCES "housing" ("housing_id");

ALTER TABLE "credentials" ADD CONSTRAINT "fk_credentials_housing_id" FOREIGN KEY("housing_id")
REFERENCES "housing" ("housing_id");

ALTER TABLE "landlord" ADD CONSTRAINT "fk_landlord_housing_id" FOREIGN KEY("housing_id")
REFERENCES "housing" ("housing_id");

ALTER TABLE "tenant" ADD CONSTRAINT "fk_tenant_housing_id" FOREIGN KEY("housing_id")
REFERENCES "housing" ("housing_id");

