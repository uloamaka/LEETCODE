-- Write your PostgreSQL query statement below
SELECT 
    project_id, 
    ROUND(SUM(E.experience_years) / COUNT(P.employee_id )::numeric, 2) AS average_years 
FROM Project P
LEFT JOIN Employee E ON P.employee_id = E.employee_id
GROUP BY P.project_id
ORDER BY P.project_id;
