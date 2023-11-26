CREATE TABLE "users" (
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
);