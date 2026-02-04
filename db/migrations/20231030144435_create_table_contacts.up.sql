create table contacts
(
    id           varchar(100) not null,
    nik          varchar(16)  not null,
    full_name    varchar(150) not null,
    legal_name   varchar(150) not null,
    birth_place  varchar(100) not null,
    birth_date   varchar(20)  not null,
    salary       bigint       not null,
    ktp_photo    varchar(500) not null,
    selfie_photo varchar(500) not null,
    user_id      bigint       not null,
    created_at   timestamp    not null default current_timestamp,
    updated_at   timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key fk_contacts_user_id (user_id) references users (id)
) engine = innodb;
