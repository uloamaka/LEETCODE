-- Write your PostgreSQL query statement below

SELECT T1.employee_id,T1.name, COUNT(T2.employee_id) AS reports_count , ROUND(AVG(T2.age)) AS average_age 
FROM Employees T1
INNER JOIN Employees T2 ON T1.employee_id  = T2.reports_to  
GROUP BY T1.employee_id, T1.name
ORDER BY employee_id;
