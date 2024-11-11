-- Write your PostgreSQL query statement below
SELECT T1.name
FROM Employee T1
INNER JOIN Employee T2 ON T1.id = T2.managerId 
GROUP BY T1.name, T2.managerId HAVING COUNT(T2.managerId) >= 5;