create table user_information(
    userid int not null primary key auto_increment,
    email varchar(255) not null unique,
    password varchar(65) not null,
    usertype int not null default 0,
    deleted int not null default 0,
    token varchar(255),
    resource blob,
    resource_count int default 0,
    qouta int default -1
);

Alter table user_information auto_increment = 1000000;

insert into user_information(email, password, usertype)  values('protik2095@gmail.com', '$2a$04$AiKxHSQj9kXFIVdE6Vj9h.LIYMoDgzHWIpr7MYreGQiJ.a4YAWQzW', 1)