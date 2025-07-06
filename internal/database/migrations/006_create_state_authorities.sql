CREATE TABLE state_authorities (
    id SERIAL PRIMARY KEY,
    state VARCHAR(50) NOT NULL,
    authority_name VARCHAR(100) NOT NULL,
    twitter_handle VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE
); 