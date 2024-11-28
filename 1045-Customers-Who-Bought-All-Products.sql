-- Write your PostgreSQL query statement below
-- group the customer table by the customer_id
--  count the distinct product_key for each customer 
-- count the total numbers of products 
-- while checking the ttl_product_count and the customer_prd_count


SELECT customer_id
FROM Customer 
GROUP BY customer_id
HAVING COUNT(DISTINCT product_key) = (SELECT COUNT(product_key) AS ttl_product_count
FROM Product)
