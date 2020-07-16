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
        unique (id)
);

alter table user
    add primary key (id);

INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (1, 'xuchu', 'lalala', null, '2020-07-15 15:24:34', '2020-07-15 15:24:34', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (2, 'xuchu', 'lalala', null, '2020-07-15 15:25:42', '2020-07-15 15:25:42', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (3, 'xuchu', 'lalala', null, '2020-07-15 15:27:49', '2020-07-15 15:27:49', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (4, 'xuchu', 'lalala', null, '2020-07-15 15:29:16', '2020-07-15 15:29:16', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (5, 'xuchu', 'lalala', 'dev', '2020-07-15 15:29:46', '2020-07-15 15:29:46', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (6, 'xuchu', 'lalala', null, '2020-07-15 15:30:42', '2020-07-15 15:30:42', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (7, 'xuchu', 'lalala', null, '2020-07-15 15:33:48', '2020-07-15 15:33:48', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (8, 'xuchu', 'lalala', null, '2020-07-15 15:35:30', '2020-07-15 15:35:30', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (9, 'xuchu', 'lalala', null, '2020-07-15 15:37:32', '2020-07-15 15:37:32', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (10, 'xuchu', 'lalala', null, '2020-07-15 15:42:33', '2020-07-15 15:42:33', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (11, 'xuchu', 'lalala', 'dev', '2020-07-15 15:47:47', '2020-07-15 15:47:47', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (12, '', '', 'dev', '2020-07-15 15:52:54', '2020-07-15 15:52:54', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (13, '', '', 'dev', '2020-07-15 15:54:10', '2020-07-15 15:54:10', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (14, 'u01', 'p01', 'dev', '2020-07-15 15:58:20', '2020-07-15 15:58:20', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (15, 'u01', '', 'dev', '2020-07-15 16:01:00', '2020-07-15 16:01:00', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (16, 'u01', '', 'dev', '2020-07-15 16:02:45', '2020-07-15 16:02:45', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (17, 'u01', '', 'dev', '2020-07-15 16:03:21', '2020-07-15 16:03:21', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (18, 'u01', '', 'dev', '2020-07-15 16:03:25', '2020-07-15 16:03:25', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (19, 'u01', '', 'dev', '2020-07-15 16:03:44', '2020-07-15 16:03:44', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (20, 'u01', '', 'dev', '2020-07-15 16:06:11', '2020-07-15 16:06:11', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (21, 'u01', '', 'dev', '2020-07-15 16:07:11', '2020-07-15 16:07:11', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (22, 'u01', 'p01', 'dev', '2020-07-15 16:09:58', '2020-07-15 16:09:58', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (23, 'u01', 'p03', 'dev', '2020-07-15 17:14:51', '2020-07-15 17:14:51', null);
INSERT INTO loki.user (id, username, password, user_type, created_at, updated_at, deleted_at) VALUES (24, 'u01', 'p03', 'dev', '2020-07-15 20:45:22', '2020-07-15 20:45:22', null);