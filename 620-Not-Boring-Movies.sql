-- Write your PostgreSQL query statement below
SELECT * FROM Cinema
WHERE MOD(id, 2) <> 0 AND Cinema.description NOT IN ('boring')  
ORDER BY rating DESC;