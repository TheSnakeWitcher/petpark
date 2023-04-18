-- migrate:up
CREATE TABLE pets (

    -- auto
    id       UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    -- required
    picked   DATE NOT NULL,
    address  TEXT NOT NULL,
    contact  VARCHAR(20) NOT NULL,
    details  TEXT NOT NULL

) ;

-- migrate:down
DROP TABLE IF EXISTS pets ;
