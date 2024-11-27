-- Write your PostgreSQL query statement below


-- WITH CTE_TEST1 AS (SELECT 
--     MIN(event_date) AS  first_login,
--     MIN(event_date) + 1  AS  next_day_after_first_login_date,
--     player_id
-- FROM Activity
-- GROUP BY player_id) 
-- SELECT 
--     ROUND(1.0 * COUNT(A.player_id) / COUNT(CTE_TEST1.player_id), 2) AS fraction
-- FROM Activity A 
-- RIGHT JOIN CTE_TEST1 
--     ON A.player_id = CTE_TEST1.player_id AND A.event_date = CTE_TEST1.next_day_after_first_login_date;

SELECT ROUND(
    COUNT(player_id)::NUMERIC / 
    (SELECT COUNT(DISTINCT player_id)
    FROM Activity), 2) AS fraction
FROM Activity
WHERE (player_id, event_date) IN (
    SELECT player_id, MIN(event_date) + 1
    FROM Activity
    GROUP BY player_id
)