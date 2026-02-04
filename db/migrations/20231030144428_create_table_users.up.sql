create table users
(
    id                     bigint       not null auto_increment,
    username               varchar(100) not null,
    name                   varchar(100) not null,
    password               varchar(100) not null,
    token                  varchar(100) null,
    failed_login_attempts  int          not null default 0,
    last_failed_login_at   bigint       not null default 0,
    locked_until           bigint       not null default 0,
    last_login_at          bigint       not null default 0,
    created_at             timestamp    not null default current_timestamp,
    updated_at             timestamp    not null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key uk_users_username (username)
) engine = InnoDB;
