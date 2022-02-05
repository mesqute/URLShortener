
FROM postgres:latest

# копирование файла для использования встроенного механизма инициализации postgre контейнера
COPY postgresql.sql /docker-entrypoint-initdb.d/postgresql.sql
