
CREATE TABLE "urls"
(
    id   serial      not null unique,
    url  text        not null unique,
    link varchar(10) not null unique
);

CREATE UNIQUE INDEX link ON urls (link) INCLUDE (url);