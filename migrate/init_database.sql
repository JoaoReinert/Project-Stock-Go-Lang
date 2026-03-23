create table if not exists category
(
    id   int auto_increment primary key,
    name varchar(150) null
);

create table if not exists sub_category
(
    id          int auto_increment primary key,
    id_category int          null,
    name        varchar(150) null,
    foreign key (id_category) references category (id)
);

create table if not exists unit_stock
(
    id   int auto_increment primary key,
    name varchar(150) null
);

create table if not exists equipment
(
    id              int auto_increment primary key,
    name            varchar(150) null,
    id_sub_category int          null,
    id_unit_stock   int          null,
    foreign key (id_sub_category) references sub_category (id),
    foreign key (id_unit_stock) references unit_stock (id)
);

create table if not exists stock_item
(
    id            int auto_increment primary key,
    id_equipment  int null,
    id_unit_stock int null,
    status_code   int null,
    code          int null,
    foreign key (id_equipment) references equipment (id),
    foreign key (id_unit_stock) references unit_stock (id)
);

create table if not exists user
(
    id       int auto_increment primary key,
    name     varchar(100) not null,
    email    varchar(255) null,
    password varchar(255) not null,
    salt     varchar(255) not null
);

create table if not exists movement
(
    id      int auto_increment primary key,
    date    date not null,
    id_user int  not null,
    foreign key (id_user) references user (id)
);

create table if not exists historic_movement
(
    id            int auto_increment primary key,
    id_movement   int null,
    id_stock_item int null,
    foreign key (id_movement) references movement (id),
    foreign key (id_stock_item) references stock_item (id)
);

