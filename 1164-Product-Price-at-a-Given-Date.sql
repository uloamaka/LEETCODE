-- Write your PostgreSQL query statement below
-- #We want to filter out only the date before '2019-08-16' inclusively, and rank the price based on latest date using window function rank() over (). We could use first_value() over () as well. We rank the price grouped by product_id, so we use 'partition by product_id', and we want to keep the latest record only, so we use order by change_date desc.
-- We can further filter out the latest price by saying that we want only the price who is ranked #1.

-- Then we need to consider whether there's product changed price after '2019-08-16'. We can make sure of that by using union function. We want to find out which product changed price at the first time after '2019-08-16', and set the original price at $10.

WITH CTE_RANK AS (
 SELECT 
    product_id, 
    new_price AS price,
    RANK() OVER (PARTITION BY product_id ORDER BY change_date DESC) AS LATEST_DATE
 FROM Products
 WHERE change_date <= '2019-08-16'::DATE
)
SELECT
    product_id,
    price
FROM CTE_RANK
WHERE LATEST_DATE = 1
UNION 
SELECT DISTINCT product_id, 10 AS price
FROM Products
WHERE product_id NOT IN (SELECT product_id FROM Products WHERE change_date <= '2019-08-16');
