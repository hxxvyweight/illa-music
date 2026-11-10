-- 1. Insert Artists
INSERT INTO artists (id, name) VALUES 
(1, 'hxxvyweight'),
(2, 'Bicep'),
(3, 'Charlotte de Witte');

-- 2. Insert Genres
INSERT INTO genres (id, name) VALUES 
(1, 'Techno'),
(2, 'Hardgroove'),
(3, 'UKG'),
(4, 'House');

-- 3. Insert Venues
INSERT INTO venues (id, name, location) VALUES 
(1, 'Printworks', 'London, UK'),
(2, 'Berghain', 'Berlin, Germany'),
(3, 'Fabric', 'London, UK');

-- 4. Map Artists to Genres (Many-to-Many)
INSERT INTO artists_genres (artist_id, genre_id) VALUES 
(1, 1), -- hxxvyweight plays Techno
(1, 2), -- hxxvyweight plays Hardgroove
(2, 3), -- Bicep plays UKG
(2, 4), -- Bicep plays House
(3, 1); -- Charlotte de Witte plays Techno

-- 5. Map Artists to Venues (Many-to-Many)
INSERT INTO artists_venues (artist_id, venue_id) VALUES 
(1, 3), -- hxxvyweight played at Fabric
(2, 1), -- Bicep played at Printworks
(2, 3), -- Bicep played at Fabric
(3, 2); -- Charlotte de Witte played at Berghain
