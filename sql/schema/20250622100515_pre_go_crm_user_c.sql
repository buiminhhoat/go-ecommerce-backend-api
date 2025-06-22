-- +goose Up
-- +goose StatementBegin
create table pre_go_crm_user_c
(
    usr_id               int unsigned auto_increment comment 'Account ID'
        primary key,
    usr_email            varchar(30) default '' not null comment 'Email',
    usr_phone            varchar(15) default '' not null comment 'Phone Number',
    usr_username         varchar(30) default '' not null comment 'Username',
    usr_password         varchar(32) default '' not null comment 'Password',
    usr_created_at       int         default 0  not null comment 'Creation Time',
    usr_updated_at       int         default 0  not null comment 'Update Time',
    usr_create_ip_at     varchar(12) default '' not null comment 'Creation IP',
    usr_last_login_at    int         default 0  not null comment 'Last Login Time',
    usr_last_login_ip_at varchar(12) default '' not null comment 'Last Login IP',
    usr_login_times      int         default 0  not null comment 'Login Times',
    usr_status           tinyint(1)  default 0  not null comment 'Status 1:enable, 0:disable, -1:deleted'
)
    comment 'Account';

create index idx_email
    on pre_go_crm_user_c (usr_email);

create index idx_phone
    on pre_go_crm_user_c (usr_phone);

create index idx_username
    on pre_go_crm_user_c (usr_username);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pre_go_crm_user_c;
-- +goose StatementEnd
