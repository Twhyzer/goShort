CREATE TABLE links (
   id serial PRIMARY KEY,
   targetUrl TEXT NOT NULL,
   requestKey TEXT NOT NULL,
   redirects INTEGER DEFAULT '0',
   creation TIMESTAMP DEFAULT current_timestamp
);