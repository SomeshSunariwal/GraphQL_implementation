CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    book_name VARCHAR NOT NULL,
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
INSERT INTO  books (id, book_name, author, seller) 
    VALUES (1, 'book 1', 'author 1', 'seller 1', true),
    VALUES (2, 'book 2', 'author 2', 'seller 2', true),
    VALUES (3, 'book 3', 'author 3', 'seller 3', true),
    VALUES (4, 'book 4', 'author 4', 'seller 4', true),
    VALUES (5, 'book 5', 'author 5', 'seller 5', true);

INSERT INTO availability_zone (id, book_name, address) 
    VALUES (1, "book1", "RAJ"), 
    VALUES (2, "book1", "IXJ"), 
    VALUES (3, "book1", "SXR"), 
    VALUES (4, "book2", "KER"), 
    VALUES (5, "book2", "PNB");