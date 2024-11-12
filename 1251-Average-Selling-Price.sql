-- Write your PostgreSQL query statement below
SELECT 
    p.product_id,
    ROUND(COALESCE(SUM(p.price * COALESCE(s.units, 0)) / NULLIF(SUM(s.units)::numeric, 0), 0), 2) AS average_price
FROM Prices AS P
LEFT JOIN UnitsSold AS s ON p.product_id = s.product_id
AND s.purchase_date BETWEEN p.start_date AND p.end_date
GROUP BY p.product_id;