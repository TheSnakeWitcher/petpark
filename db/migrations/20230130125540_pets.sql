-- migrate:up
CREATE TABLE pets (
    id       INT PRIMARY KEY,
    name     VARCHAR(20),
    picked   DATE,
    location TEXT,
    data     JSON
) ;

-- migrate:down
DROP TABLE IF EXISTS pets ;
