-- Write your PostgreSQL query statement below


WITH CTE_DELIVERY AS 
    ( SELECT DISTINCT ON (customer_id) customer_id,
        CASE WHEN order_date = customer_pref_delivery_date THEN 1 ELSE NULL END AS "immediate",
        CASE WHEN order_date != customer_pref_delivery_date THEN 1 ELSE NULL END AS "scheduled"
FROM Delivery   
ORDER BY customer_id, order_date)
SELECT ROUND(COUNT(immediate)::NUMERIC / COUNT(*) * 100, 2) AS immediate_percentage FROM CTE_DELIVERY