create table if not exists tb_user (
    -- 主键
    id int primary key auto_increment,
    -- 用户名
    user_name varchar(16) not null unique,
    check (length(user_name) >= 6 and length(user_name) <= 16),
    -- 密码的哈希
    password_hash text not null,
    -- 头像url
    avatar_url text,
    -- 创建时间
    create_time timestamp not null default current_timestamp,
    -- 是否是管理员
    is_admin tinyint not null default false,
    -- 状态 0:正常 -1:删除, 可能会有其他状态
    status tinyint not null default 0,
    -- 备注
    remark text,
    -- 预留字段0
    reserve0 text,
    -- 预留字段1
    reserve1 text,
    -- 预留字段2
    reserve2 text,
    -- 预留字段3
    reserve3 text,
    -- 预留字段4
    reserve4 text
);
