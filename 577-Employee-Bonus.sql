-- Write your PostgreSQL query statement below
SELECT 
    e.name,
    b.bonus
FROM Employee AS e
FULL OUTER JOIN Bonus AS b ON e.empId = b.empId
WHERE b.bonus IS NULL OR b.bonus < 1000;
