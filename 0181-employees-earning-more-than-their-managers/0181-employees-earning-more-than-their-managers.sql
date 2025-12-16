-- Write your PostgreSQL query statement below
SELECT e.name as Employee
FROM Employee AS e
INNER JOIN Employee AS m ON e.managerId = m.id
WHERE e.salary > m.salary;