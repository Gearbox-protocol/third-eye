alter table pools add _version integer, add name varchar(200);
alter table credit_managers add name varchar(200);
update pools set _version=1;