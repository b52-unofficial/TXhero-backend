-- DROP TABLE public.transaction_info;
-- DROP TABLE public.bid;
-- DROP TABLE public.builder;
-- DROP TABLE public.reward;

CREATE TABLE public.transaction_info (
     id SERIAL PRIMARY KEY,
     tx_hash varchar NOT NULL,
     gas_fee float8 NOT NULL,
     from_address varchar NOT NULL,
     "timestamp" timestamp NOT NULL,
     status int4 NOT NULL DEFAULT 0,
     reward float8 NOT NULL DEFAULT 0,
     create_dt timestamp NOT NULL DEFAULT now(),
     update_dt timestamp NOT NULL DEFAULT now()
);

comment on table transaction_info is 'tx hero transaction data';
comment on column transaction_info.tx_hash is 'transaction hash';
comment on column transaction_info.gas_fee is 'transaction gas fee';
comment on column transaction_info.from_address is 'transaction from address';
comment on column transaction_info.timestamp is 'transaction timestamp';
comment on column transaction_info.status is 'transaction status';

alter table transaction_info
    owner to postgres;
GRANT ALL ON TABLE public.transaction_info TO postgres;


CREATE TABLE public.builder (
        id SERIAL PRIMARY KEY,
        builder_name varchar NULL,
        address varchar NOT NULL,
        endpoint varchar NOT NULL,
        description varchar NULL,
        create_dt timestamp NOT NULL DEFAULT now(),
        update_dt timestamp NOT NULL DEFAULT now()
);

comment on table builder is 'builder whitelist';
alter table builder
    owner to postgres;
GRANT ALL ON TABLE public.builder TO postgres;

CREATE TABLE public.bid (
    round SERIAL PRIMARY KEY NOT NULL ,
    top_bid float8 NOT NULL,
    total_gas_fee float8 NULL,
    builder_id int4 NOT NULL,
    start_timestamp timestamp NOT NULL DEFAULT now(),
    end_timestamp timestamp NOT NULL DEFAULT now(),
    create_dt timestamp NOT NULL DEFAULT now(),
    update_dt timestamp NOT NULL DEFAULT now(),
    CONSTRAINT builder_id FOREIGN KEY (builder_id) REFERENCES public.builder(id)
);

comment on table bid is 'bid info';
ALTER TABLE public.bid OWNER TO postgres;
GRANT ALL ON TABLE public.bid TO postgres;

CREATE TABLE public.reward (
    address varchar NOT NULL PRIMARY KEY,
    reward float8 NOT NULL
);

ALTER TABLE public.reward OWNER TO postgres;
GRANT ALL ON TABLE public.reward TO postgres;


