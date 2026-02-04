create table contact_limits
(
    id           varchar(100) not null,
    contact_id   varchar(100) not null,
    tenor        int          not null,
    limit_amount bigint       not null,
    created_at   timestamp    not null default current_timestamp,
    updated_at   timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key uk_contact_limits_contact_tenor (contact_id, tenor),
    foreign key fk_contact_limits_contact_id (contact_id) references contacts (id)
) engine = innodb;
