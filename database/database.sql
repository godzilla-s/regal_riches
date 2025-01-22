CREATE DATABASE regal_riches;

CREATE TABLE userinfo (
    id serial primary key,
    name  varchar(100) not null default '',
    active boolean not null default true,
    state varchar(33) not null default '',
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz not null default now()
);

CREATE TABLE rr_account (
    userid  int not null default '',
    balance int not null default '',
)

CREATE TABLE rr_txn_detail (
    id  serial primary key,
    userid int not null default '',
    txn_amount int not null default 0,
    source_id int not null default 0,
    type varchar(33) not null default '',
    created_at timestamptz not null default now()
);

CREATE TABLE rr_source_type (
    id  int not null default '',
    name  varchar(33) not null,
);

CREATE TABLE ton_txn_detail (
    txn_id  serial primary key,
    user_id int not null,
    type int not null,
    amount int not null default 0,
    created_at timestamptz not null default now()
);
