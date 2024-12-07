CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE pet_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE pets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type_id INT NOT NULL,
    age INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (type_id) REFERENCES pet_types(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (name, email) VALUES
    ('田中太郎', 'tanaka@example.com'),
    ('山田花子', 'yamada@example.com'),
    ('佐藤次郎', 'sato@example.com');

INSERT INTO pet_types (name) VALUES
    ('犬'),
    ('猫'),
    ('鳥'),
    ('ハムスター');

INSERT INTO pets (name, type_id, age, user_id) VALUES
    ('ポチ', 1, 3, 1),   -- 田中さんの犬
    ('タマ', 2, 5, 1),   -- 田中さんの猫
    ('ピー子', 3, 2, 2); -- 山田さんの鳥