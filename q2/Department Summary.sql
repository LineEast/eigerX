create table department (
	ID serial,
	name text not null,
	location text not null,

	primary key (ID)
)

create table employee (
	ID serial,
	name text not null,
	salary integer,
	dept_id integer not null,

	foreign key (dept_id) references department (ID),

	primary key (ID)
)

-- собрать сводку по кол-ву сотров в каждом отделе
-- Перечислить каждый отдел, даже если никого нет
-- Отсортировать по кол-ву сотров и по алфовитному, если кол-во сотров совпадает

select department.name, count(employee.ID) as count
from department left join employee
on department.ID = employee.dept_id
group by department.ID, department.name
order by department.name, count;