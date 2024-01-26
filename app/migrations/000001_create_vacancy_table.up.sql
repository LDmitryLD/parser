CREATE TABLE IF NOT EXISTS vacancy (
    id SERIAL PRIMARY KEY,
    context TEXT,
    type TEXT,
    date_posted TEXT,
    title TEXT,
    description TEXT,
    valid_through TEXT,
    job_location JSONB,
    job_location_type TEXT,
    employment_type TEXT
);

CREATE TABLE IF NOT EXISTS hiring_organization (
    id SERIAL PRIMARY KEY,
    vacancy_id INT REFERENCES vacancy(id) ON DELETE CASCADE,
    type TEXT,
    name TEXT,
    logo TEXT,
    same_as TEXT
);

CREATE TABLE IF NOT EXISTS identifier (
    id SERIAL PRIMARY KEY,
    vacancy_id INT  REFERENCES vacancy(id) ON DELETE CASCADE,
    type TEXT,
    name TEXT,
    value TEXT
);

-- CREATE TABLE IF NOT EXISTS job_location (
--     id SERIAL PRIMARY KEY,
--     vacancy_id INT REFERENCES vacancy(id) ON DELETE CASCADE,
--     type TEXT
-- );

-- CREATE TABLE IF NOT EXISTS address (
--     id SERIAL PRIMARY KEY,
--     job_location_id INT REFERENCES job_location(id) ON DELETE CASCADE,
--     type TEXT,
--     street_address TEXT,
--     address_locality TEXT
-- );

-- CREATE TABLE IF NOT EXISTS address_country (
--     id SERIAL PRIMARY KEY,
--     address_id INT REFERENCES address(id) ON DELETE CASCADE,
--     type TEXT,
--     name TEXT
-- );

--поменял  date на text