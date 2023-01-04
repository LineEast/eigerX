Даны 2 таблицы в базе данных (MySQL / PostgreSQL):
- Eployee
- Department

Cгенерируйте сводку о количестве сотрудников в каждом отделе. 
Перечислите каждый отдел, даже если в нем нет сотрудников. 
Отсортируйте результаты от высокого к низкому по количеству сотрудников, а затем в алфавитном порядке по отделам, если в отделах одинаковое количество сотрудников.
В результатах должны быть указаны 
1) название отдела 
2) количество сотрудников в соответствующих названиях столбцов.

### Схема таблиц
##### EMPLOYEE:
- ID (Integer) - Employee ID number. This is a primary key.
- NAME (String) - Employee name
- SALARY (Integer) - Employee salary
- DEPT_ID (Integer) - ID of the employee's department, a foreign key to DEPARTMENT.ID.
#### DEPARTMENT:
- ID (Integer) - Department ID. This is a primary key.
- NAME (String) - Department name.
- LOCATION (String) - Department location.

### Пример таблиц
#### EMPLOYEE:

| ID | NAME | SALARY | DEPT_ID |
|----|------|--------|---------|
| 1 | Candice | 4685 | 1 |
| 2 | Julia | 2559 | 2 |
| 3 | Bob | 4405 | 4 |
| 4 | Scarlet | 2350 | 1 |
| 5 | Ileana | 1151 | 4 |

#### DEPARTMENT:

| ID | NAME | LOCATION | 
|----|------|----------|
| 1 | Executive | Sydney |
| 2 | Production | Sydney |
| 3 | Resources | Cape Town |
| 4 | Technical | Texas |
| 5 | Management | Paris |


#### Пример результата

|||
|---|---|
| Executive | 2 |
| Technical | 2 |
| Production | 1 |
| Management | 0 |
| Resources | 0 |