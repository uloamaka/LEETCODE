-- Write your PostgreSQL query statement below
-- Rank with the turns colummn 
-- Sum up the weight column, to get up to 1000
-- select the name of the passenger that made the total weight to 1000

WITH CTE_ORD AS (
    SELECT 
        person_name,
        SUM(weight) OVER (ORDER BY turn) AS total_weight
    FROM 
    Queue)
    
    SELECT 
       LAST_VALUE(person_name) OVER (ORDER BY total_weight DESC) AS person_name
    FROM  CTE_ORD
    WHERE total_weight::NUMERIC <= 1000
    LIMIT 1
;