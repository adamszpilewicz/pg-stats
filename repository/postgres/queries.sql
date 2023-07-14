SELECT datname       AS database_name,
       numbackends   AS number_of_connections,
       xact_commit   AS total_transactions_committed,
       xact_rollback AS total_transactions_rolledback,
       blks_read     AS blocks_read,
       blks_hit      AS blocks_hit,
       tup_returned  AS tuples_returned,
       tup_fetched   AS tuples_fetched,
       tup_inserted  AS tuples_inserted,
       tup_updated   AS tuples_updated,
       tup_deleted   AS tuples_deleted
FROM pg_stat_database;

select * from pg_stat_database;
select * from pg_stat_user_tables;

create table test (id int, name varchar(20));

CREATE TABLE people (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(120) DEFAULT 'first name',
                         age INT DEFAULT 35,
                         location VARCHAR(120) DEFAULT 'Poland'
);

insert into people (name, age, location) values ('John', 25, 'USA');
select * from people;


CREATE TABLE jobs (
                        id SERIAL PRIMARY KEY,
                        employer VARCHAR(120) DEFAULT 'employer',
                        salary float4 DEFAULT 38000.00,
                        job_title VARCHAR(120) DEFAULT 'software engineer'
);

insert into jobs (employer, salary, job_title) values ('Google', 100000.00, 'software engineer');
select * from jobs;

select * from pg_stat_user_tables;

select * from pg_stat_;


