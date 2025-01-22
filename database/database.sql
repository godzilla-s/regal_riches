CREATE DATABASE regal_riches;

CREATE TABLE userinfo (
    username  varchar(100) not null default '',
);

CREATE TABLE rr_account (
    userid  int not null default '',
    balance int not null default '',
)

CREATE TABLE rr_txn_detail (
    userid int not null default '',
    txn_id varchar(66) not null default '',
    txn_amoutn int not null default 0
)

CREATE TABLE rr_source_type (
    id  int not null default '',
    name  varchar(33) not null,
);