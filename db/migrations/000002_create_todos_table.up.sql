CREATE TABLE IF NOT EXISTS todos (
  id uuid PRIMARY KEY,
  created timestamp DEFAULT NOW(),
  description TEXT DEFAULT '' NOT NULL
);
