create table users (
  id SERIAL PRIMARY KEY,
  first_name TEXT NOT NULL,
  last_name TEXT,
  city TEXT,
  created_at timestamp without time zone
);
