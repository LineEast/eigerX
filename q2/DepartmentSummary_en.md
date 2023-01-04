Given 2 tables in a database (MySQL / PostgreSQL), Employee and Department, generate a summary of the # of employees in each department. List each department, even if it has no employees. Sort the results from high to low by number of employees, and then alphabetically by department if departments have the same number of employees. The results should list 1: department name and 2: employee count in appropriate column names.

### Schema
##### EMPLOYEE:
- ID (Integer) - Employee ID number. This is a primary key.
- NAME (String) - Employee name
- SALARY (Integer) - Employee salary
- DEPT_ID (Integer) - ID of the employee's department, a foreign key to DEPARTMENT.ID.
#### DEPARTMENT:
- ID (Integer) - Department ID. This is a primary key.
- NAME (String) - Department name.
- LOCATION (String) - Department location.

### Sample Data Tables
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


#### Sample Output
|||
|---|---|
| Executive | 2 |
| Technical | 2 |
| Production | 1 |
| Management | 0 |
| Resources | 0 |

