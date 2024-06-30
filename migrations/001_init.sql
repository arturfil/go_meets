-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "is_admin" boolean NOT NULL DEFAULT FALSE,
  "is_teacher" boolean DEFAULT FALSE,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE roles (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "description" varchar
);

CREATE TABLE role_relations (
    "user_id" uuid REFERENCES users(id),
    "role_id" uuid REFERENCES roles(id),
    PRIMARY KEY ("user_id", "role_id")

);

CREATE TABLE requests (
    "id" uuid PRIMARY KEY REFERENCES users(id),
    "status" varchar
);

CREATE TABLE meetings (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "subject_id" uuid NOT NULL,
  "student_id" uuid NOT NULL,
  "teacher_id" uuid NOT NULL,
  "student_attended" bool,
  "meeting_time" timestamptz NOT NULL,
  "meeting_day" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE subjects (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE teachings (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "teacher_id" uuid NOT NULL,
  "subject_id" uuid NOT NULL,
  "opening_time" timestamptz NOT NULL,
  "closing_time" timestamptz NOT NULL,
  "hourly_rate" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE bills (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "meeting_id" uuid NOT NULL,
  "balance_outstanding" float NOT NULL,
  "is_paid" boolean DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "meetings" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");
ALTER TABLE "meetings" ADD FOREIGN KEY ("student_id") REFERENCES "users" ("id");
ALTER TABLE "meetings" ADD FOREIGN KEY ("teacher_id") REFERENCES "users" ("id");

ALTER TABLE "teachings" ADD FOREIGN KEY ("teacher_id") REFERENCES "users" ("id");
ALTER TABLE "teachings" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "bills" ADD FOREIGN KEY ("meeting_id") REFERENCES "meetings" ("id");

CREATE INDEX ON users ("email");
CREATE INDEX ON meetings ("student_id");
CREATE INDEX ON bills ("meeting_id");

-- +goose Down
DROP TABLE IF EXISTS users             CASCADE;
DROP TABLE IF EXISTS meetings          CASCADE;
DROP TABLE IF EXISTS subjects          CASCADE;
DROP TABLE IF EXISTS teachings         CASCADE;
DROP TABLE IF EXISTS bills             CASCADE;
DROP TABLE IF EXISTS requests          CASCADE;
DROP TABLE IF EXISTS roles             CASCADE;
DROP TABLE IF EXISTS role_relations    CASCADE;
