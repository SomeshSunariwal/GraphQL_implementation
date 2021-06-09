CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    book_name VARCHAR NOT NULL UNIQUE,
    author VARCHAR NOT NULL,
    seller VARCHAR NOT NULL,
    available BOOLEAN NOT NULL
);

CREATE TABLE availability_zone (
    id SERIAL PRIMARY KEY,
    book_name VARCHAR NOT NULL,
    address VARCHAR NOT NULL
);
ALTER TABLE availability_zone ADD CONSTRAINT book_fk FOREIGN KEY (book_name) REFERENCES books(book_name);

-- Temp Data to start with
INSERT INTO  books (id, book_name, author, seller, available) 
    VALUES (1, 'book 1', 'author 1', 'seller 1', 'true'),
     (2, 'book 2', 'author 2', 'seller 2', 'true'),
     (3, 'book 3', 'author 3', 'seller 3', 'true'),
     (4, 'book 4', 'author 4', 'seller 4', 'true'),
     (5, 'book 5', 'author 5', 'seller 5', 'true');

INSERT INTO availability_zone (id, book_name, address) 
    VALUES (1, 'book 1', 'RAJ'), 
     (2, 'book 1', 'IXJ'), 
     (3, 'book 1', 'SXR'), 
     (4, 'book 2', 'KER'), 
     (5, 'book 2', 'PNB');