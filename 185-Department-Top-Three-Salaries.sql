-- Write your PostgreSQL query statement below
WITH RankedSalaries AS (
    SELECT
        departmentId,
        name AS Employee,
        salary AS Salary,
        DENSE_RANK() OVER (
            PARTITION BY departmentId 
            ORDER BY salary DESC
    ) AS rank_in_department
FROM
    Employee 
) 
  
SELECT 
    Dp.name as Department,
    Employee,
    Salary
FROM 
    RankedSalaries CTE
JOIN 
    Department Dp
ON 
    CTE.departmentId = Dp.id
WHERE 
    rank_in_department < 4;