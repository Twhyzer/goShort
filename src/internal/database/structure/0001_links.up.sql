CREATE TABLE links (
   id serial PRIMARY KEY,
   inputurl TEXT NOT NULL,
   shorturl TEXT NOT NULL,
   redirects INTEGER DEFAULT '0',
   creation TIMESTAMP DEFAULT current_timestamp
);