CREATE TABLE erc20(
    id integer NOT NULL,
    name text NOT NULL,
    symbol text NOT NULL,
    decimals integer NOT NULL,
    total_supply_low integer NOT NULL,
    total_supply_high integer NOT NULL,
    address integer NOT NULL
);

CREATE TABLE owners(
    id integer NOT NULL,
    account integer NOT NULL,
    balance integer NOT NULL,
    is_proxy integer NOT NULL,
    version integer NOT NULL
);

CREATE TABLE allowances(
    id integer NOT NULL,
    owner integer NOT NULL,
    spender integer NOT NULL,
    amount integer NOT NULL
);