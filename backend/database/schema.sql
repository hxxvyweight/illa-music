CREATE TABLE IF NOT EXISTS artists (
	id INTEGER PRIMARY KEY, 
	name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS genres (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS venues (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	location TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS artists_genres (
	artist_id INTEGER, 
	genre_id INTEGER, 
	FOREIGN KEY (artist_id) REFERENCES artists(id), 
	FOREIGN KEY (genre_id) REFERENCES genres(id)
);
CREATE TABLE IF NOT EXISTS artists_venues (
	artist_id INTEGER,
	venue_id INTEGER,
	FOREIGN KEY (artist_id) REFERENCES artists(id),
	FOREIGN KEY (venue_id) REFERENCES venues(id)
);
