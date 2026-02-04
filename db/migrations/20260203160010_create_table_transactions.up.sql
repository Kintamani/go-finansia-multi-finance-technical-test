create table transactions
(
    id                 varchar(100) not null,
    contact_id         varchar(100) not null,
    contract_number    varchar(100) not null,
    tenor              int          not null,
    channel            varchar(50)  not null,
    otr                bigint       not null,
    admin_fee          bigint       not null,
    installment_amount bigint       not null,
    interest_amount    bigint       not null,
    asset_name         varchar(150) not null,
    created_at         timestamp    not null default current_timestamp,
    updated_at         timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key uk_transactions_contract_number (contract_number),
    foreign key fk_transactions_contact_id (contact_id) references contacts (id)
) engine = innodb;
