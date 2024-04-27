CREATE TABLE IF NOT EXISTS "public"."transactions"(
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "note" TEXT,
    "amount" NUMERIC NOT NULL DEFAULT 0,
    "type" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP
);