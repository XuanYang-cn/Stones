CREATE TABLE easy_drinks
(
  drink_name VARCHAR(30) NOT NULL,
  main VARCHAR(30) NOT NULL,
  amount1 DEC(3,1) NOT NULL,
  second VARCHAR(30) NOT NULL,
  amount2 DEC(2,2) NOT NULL,
  directions BLOB NOT NULL
);

INSERT INTO easy_drinks

VALUES

('Blackthorn', 'tonic water', 1.5, 'pineapple juice', 1, 'stir with ice, strain into cocktail glass with lemon twist'),
('Blue Moon', 'soda', 1.5, 'blueberry juice', 0.75, 'stir with ice, strain into cocktail glass with lemon twist'),
('Oh My Gosh', 'peach nectar', 1.0, 'pineapple juice', 1.00, 'stir with ice, strain into shot glass'),
('Lime Fizz', 'Sprite', 1.5, 'lime juice', 0.75, 'stir with ice, strain into cocktail glass'),
('Kiss on the Lips', 'cherry juice', 2.0, 'apricot nectar', 7.00, 'serve over ice with straw'),
('Hot Gold', 'peach nectar', 3.0, 'orange juice', 6.00, 'pour hot orange juice i mug an add peach nectar'),
('Lone Tree', 'soda', 1.5, 'cherry juice', 0.75, 'stir with ice, strain into cocktail glass'),
('Greyhound', 'soda', 1.5, 'grapefruit juice', 5.00, 'serve over ice with lime slice'),
('Indian Summer', 'apple juice', 2.0, 'hot tea', 6.00, 'add juice to mug and top off with hot tea'),
('Bull Frog', 'ice tea', 1.5, 'lemanade', 5.00, 'serve over ice with lime silce'),
('Soda and it', 'soda', 2.0, 'grape juice', 1.00, 'shake in cocktail glass, no ice');

CREATE TABLE drink_info
(
  drink_name VARCHAR(30) NOT NULL,
  cost DEC(3,1) NOT NULL,
  carbs DEC(4,1) NOT NULL,
  color VARCHAR(10) NOT NULL,
  ice VARCHAR(1) NOT NULL,
  caories INTEGER NOT NULL
);

INSERT INTO drink_info
VALUES
('Blackthorn', 3, 8.4, 'yellow', 'Y', 33),
('Blue Moon', 2.5, 3.2, 'blue', 'Y', 12),
('Oh My Gosh', 3.5, 8.6, 'orange', 'Y', 35),
('Lime Fizz', 2.5, 5.4, 'green', 'Y', 24),
('Kiss on the Lips', 5.5, 42.5, 'purple', 'Y', 171),
('Hot Gold', 3.2, 32.1, 'orange', 'N', 135),
('Lone Tree', 3.6, 4.2, 'red', 'Y', 17),
('Greyhound', 4, 14, 'yellow', 'Y', 50),
('Indian Summer', 2.8, 7.2, 'brown', 'N', 30),
('Bull Frog', 2.6, 21.5, 'tan', 'Y', 80),
('Soda ant iZt', 3.8, 4.7, 'red', 'N', 19);

CREATE TABLE clown_info
(
  name VARCHAR(20) NOT NULL,
  last_seen VARCHAR(128),
  apperance BLOB NOT NULL,
  activities VARCHAR(128)
);

INSERT INTO clown_info
VALUES
('Elsie', 'Cherry Hill Senior Center', 'F, red hair, green dress, huge feet', 'balloons, little car'),
('Pickles', 'Jack Green\'s party', 'M, orange hair, blue suit, huge feet', 'mime'),
('Snuggles', 'Ball-Mart', 'F, yellow shirt, baggy red pants', 'horn, unbrella'),
('Mr.Hobo', 'BG Circus', 'M, cigar, black hair, tiny hat', 'violin'),
('Clarabelle', 'Belmont Senior Center', 'F, pink hair, huge flower, blue dress', 'yelling, dancing'),
('Scooter', 'Oakland Hospital', 'M, blue hair, red suit, huge nose', 'balloons'),
('Zippo', 'Millstone Mail', 'F, orange suit, baggy pants', 'dancing'),
('Babe', 'Earl\'s Autos', 'F, all pink and sparkly', 'balancing, little car');
