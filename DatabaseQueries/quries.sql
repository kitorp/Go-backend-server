create table user_information(
    userid int not null primary key auto_increment,
    email varchar(255) not null,
    password varchar(65) not null,
    usertype int not null default 0,
    deleted int not null default 0,
    token varchar(255),
    resource blob,
    resource_count int default 0,
    qouta int default -1
)