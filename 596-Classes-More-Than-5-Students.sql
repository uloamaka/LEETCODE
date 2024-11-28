-- Write your PostgreSQL query statement below
SELECT DISTINCT class    
FROM Courses 
GROUP BY class
HAVING COUNT(student) >= 5;