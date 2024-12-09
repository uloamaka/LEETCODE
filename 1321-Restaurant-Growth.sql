-- Write your PostgreSQL query statement below
-- FIND THE MOVING AVG WITH A 7 DAYS INTERVAL

SELECT * 
FROM (SELECT
    visited_on, 
    SUM(amount) OVER (
        ORDER BY visited_on
        ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
    ) AS amount,
    ROUND(AVG(amount) OVER (
        ORDER BY visited_on
        ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
    ), 2) AS average_amount
FROM (SELECT visited_on, SUM(amount) AS amount FROM Customer GROUP BY visited_on)
)
WHERE visited_on > (SELECT MIN(visited_on) + INTERVAL '5 days' FROM Customer);