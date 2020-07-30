create table user
(
    id         int auto_increment,
    username   varchar(200)              not null,
    password   varchar(200)              not null,
    user_type  varchar(20) default 'dev' null,
    created_at datetime                  null,
    updated_at datetime                  null,
    deleted_at datetime                  null,
    constraint user_id_uindex
        unique (id),
    constraint user_username_uindex
        unique (username)
);

alter table user
    add primary key (id);

INSERT INTO odin-inner.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (1, 'admin@cmcm.com', '$2a$10$CiPHfAbvfC8vlxd8Eu3Oa.nsePiyjGIc4AqUUFDdVZEUbnXKCOONu', 'admin', null, null, null);