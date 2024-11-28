-- Write your PostgreSQL query statement below
SELECT  product_id, year AS first_year, quantity,price FROM Sales
WHERE (product_id,year) IN (SELECT DISTINCT product_id, MIN(year)FROM Sales GROUP BY product_id);