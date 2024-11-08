create table addresses
(
    id         bigint unsigned auto_increment
        primary key,
    name       varchar(191)    not null,
    donor_id   bigint unsigned not null,
    city_id    bigint unsigned not null,
    town_id    bigint unsigned not null,
    latitude   varchar(191)    not null,
    longitude  varchar(191)    not null,
    created_at timestamp       null,
    updated_at timestamp       null
)
    collate = utf8mb4_unicode_ci;

create table beneficiaries_requests
(
    id             bigint unsigned auto_increment
        primary key,
    car_available  tinyint(1)                                                                                                                                      not null,
    is_scheduled   tinyint(1) default 0                                                                                                                            not null,
    note           text                                                                                                                                            null,
    count_unit     int                                                                                                                                             not null,
    relationship   enum ('father', 'mother', 'brother', 'sister', 'husband', 'wife', 'son', 'daughter', 'grandfather', 'grandmother', 'grandson', 'granddaughter') not null,
    request_type   enum ('emergency', 'normal')                                                                                                                    not null,
    blood_type     enum ('plasma', 'platelets')                                                                                                                    not null,
    status         enum ('pending', 'in_progress', 'accepted', 'rejected', 'cancelled', 'done')                                                                    not null,
    blood_group_id bigint unsigned                                                                                                                                 not null,
    disease_id     bigint unsigned                                                                                                                                 not null,
    hospital_id    bigint unsigned                                                                                                                                 not null,
    city_id        bigint unsigned                                                                                                                                 not null,
    beneficiary_id bigint unsigned                                                                                                                                 not null,
    created_at     timestamp                                                                                                                                       null,
    updated_at     timestamp                                                                                                                                       null
)
    collate = utf8mb4_unicode_ci;

create table blood_groups
(
    id         bigint unsigned auto_increment
        primary key,
    code       varchar(191) not null,
    created_at timestamp    null,
    updated_at timestamp    null,
    constraint blood_groups_code_unique
        unique (code)
)
    collate = utf8mb4_unicode_ci;

create table blood_types
(
    id         bigint unsigned auto_increment
        primary key,
    name       varchar(191) not null,
    created_at timestamp    null,
    updated_at timestamp    null
)
    collate = utf8mb4_unicode_ci;

create table cache
(
    `key`      varchar(191) not null
        primary key,
    value      mediumtext   not null,
    expiration int          not null
)
    collate = utf8mb4_unicode_ci;

create table cache_locks
(
    `key`      varchar(191) not null
        primary key,
    owner      varchar(191) not null,
    expiration int          not null
)
    collate = utf8mb4_unicode_ci;

create table cities
(
    id         bigint unsigned auto_increment
        primary key,
    name       varchar(191)    not null,
    country_id bigint unsigned not null,
    created_at timestamp       null,
    updated_at timestamp       null
)
    collate = utf8mb4_unicode_ci;

create table countries
(
    id         bigint unsigned auto_increment
        primary key,
    name       varchar(191) not null,
    created_at timestamp    null,
    updated_at timestamp    null
)
    collate = utf8mb4_unicode_ci;

create table diseases
(
    id               bigint unsigned auto_increment
        primary key,
    name             varchar(191)         not null,
    request_repeated tinyint(1) default 0 not null,
    period varchar (191) default '0' not null,
    created_at       timestamp            null,
    updated_at       timestamp            null
)
    collate = utf8mb4_unicode_ci;

create table donor_archives
(
    id                       bigint unsigned auto_increment
        primary key,
    donor_id                 bigint unsigned not null,
    beneficiaries_request_id bigint unsigned not null,
    created_at               timestamp       null,
    updated_at               timestamp       null
)
    collate = utf8mb4_unicode_ci;

create table donors
(
    id             bigint unsigned auto_increment
        primary key,
    first_name     varchar(191)            not null,
    last_name      varchar(191)            not null,
    phone_number   varchar(191)            not null,
    blood_group_id bigint unsigned         not null,
    birth_date     timestamp               not null,
    gender         enum ('male', 'female') not null,
    car_available  tinyint(1)              not null,
    paid           tinyint(1) default 0    not null,
    created_at     timestamp               null,
    updated_at     timestamp               null,
    constraint donors_phone_number_unique
        unique (phone_number)
)
    collate = utf8mb4_unicode_ci;

create table failed_jobs
(
    id         bigint unsigned auto_increment
        primary key,
    uuid       varchar(191)                        not null,
    connection text                                not null,
    queue      text                                not null,
    payload    longtext                            not null,
    exception  longtext                            not null,
    failed_at  timestamp default CURRENT_TIMESTAMP not null,
    constraint failed_jobs_uuid_unique
        unique (uuid)
)
    collate = utf8mb4_unicode_ci;

create table hospitals
(
    id         bigint unsigned auto_increment
        primary key,
    name       varchar(191) not null,
    created_at timestamp    null,
    updated_at timestamp    null
)
    collate = utf8mb4_unicode_ci;

create table job_batches
(
    id             varchar(191) not null
        primary key,
    name           varchar(191) not null,
    total_jobs     int          not null,
    pending_jobs   int          not null,
    failed_jobs    int          not null,
    failed_job_ids longtext     not null,
    options        mediumtext   null,
    cancelled_at   int          null,
    created_at     int          not null,
    finished_at    int          null
)
    collate = utf8mb4_unicode_ci;

create table jobs
(
    id           bigint unsigned auto_increment
        primary key,
    queue        varchar(191)     not null,
    payload      longtext         not null,
    attempts     tinyint unsigned not null,
    reserved_at  int unsigned     null,
    available_at int unsigned     not null,
    created_at   int unsigned     not null
)
    collate = utf8mb4_unicode_ci;

create index jobs_queue_index
    on jobs (queue);

create table migrations
(
    id        int unsigned auto_increment
        primary key,
    migration varchar(191) not null,
    batch     int          not null
)
    collate = utf8mb4_unicode_ci;

create table password_reset_tokens
(
    email      varchar(191) not null
        primary key,
    token      varchar(191) not null,
    created_at timestamp    null
)
    collate = utf8mb4_unicode_ci;

create table request_feedback
(
    id                 bigint unsigned auto_increment
        primary key,
    request_id         bigint unsigned not null,
    beneficiaries_note text            not null,
    donor_note         text            not null,
    donor_rate         int             not null,
    beneficiaries_rate int             not null,
    created_at         timestamp       null,
    updated_at         timestamp       null
)
    collate = utf8mb4_unicode_ci;

create table scheduled_requests
(
    id             bigint unsigned auto_increment
        primary key,
    blood_type     enum ('plasma', 'platelets') not null,
    count_unit     int                          not null,
    period int not null,
    beneficiary_id bigint unsigned              not null,
    hospital_id    bigint unsigned              not null,
    disease_id     bigint unsigned              not null,
    city_id        bigint unsigned              not null,
    blood_group_id bigint unsigned              not null,
    created_at     timestamp                    null,
    updated_at     timestamp                    null
)
    collate = utf8mb4_unicode_ci;

create table sessions
(
    id            varchar(191)    not null
        primary key,
    user_id       bigint unsigned null,
    ip_address    varchar(45)     null,
    user_agent    text            null,
    payload       longtext        not null,
    last_activity int             not null
)
    collate = utf8mb4_unicode_ci;

create index sessions_last_activity_index
    on sessions (last_activity);

create index sessions_user_id_index
    on sessions (user_id);

create table users
(
    id                bigint unsigned auto_increment
        primary key,
    name              varchar(191) not null,
    email             varchar(191) not null,
    email_verified_at timestamp    null,
    password          varchar(191) not null,
    remember_token    varchar(100) null,
    created_at        timestamp    null,
    updated_at        timestamp    null,
    constraint users_email_unique
        unique (email)
)
    collate = utf8mb4_unicode_ci;

