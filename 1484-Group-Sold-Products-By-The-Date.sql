-- Write your PostgreSQL query statement below
SELECT 
    sell_date,
    COUNT(DISTINCT product) AS num_sold,
    array_to_string(array_agg(DISTINCT product), ',') AS products
FROM 
    Activities
GROUP BY 
    sell_date
ORDER BY 
    sell_date;

