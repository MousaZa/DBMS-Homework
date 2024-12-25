INSERT INTO categories (name)
VALUES ('Fiction'), ('Science'), ('Biography'), ('History'), ('Fantasy');

INSERT INTO books (title, author, categoryId, language, summary, available)
VALUES 
  ('The Great Gatsby', 'F. Scott Fitzgerald', 1, 'English',  'A classic novel set in the 1920s.', TRUE),
  ('Brief History of Time', 'Stephen Hawking', 2, 'English',  'A book about cosmology and black holes.', TRUE),
  ('Steve Jobs', 'Walter Isaacson', 3, 'English',  'Biography of Apples co-founder.', TRUE),
  ('World War II', 'Max Hastings', 4, 'English',  'A detailed account of WWII.', TRUE),
  ('Harry Potter', 'J.K. Rowling', 5, 'English',  'A wizards adventure.', TRUE);

INSERT INTO costumers (username, email, password, canborrow)
VALUES 
  ('Mousa', 'm@gmail.com', '1234', TRUE);

INSERT INTO admins (username, email, password, libraryName)
VALUES ('Admin', 'a@gmail.com', '1234', 'Main Library');