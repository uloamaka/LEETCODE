-- Write your PostgreSQL query statement below
-- Group the account into salary ranges 
-- count the numbers of each categories

WITH CTE_SUM AS (
    SELECT 
        SUM(CASE WHEN income < 20000 THEN 1 ELSE 0 END) AS "Low Salary",
        SUM(CASE WHEN income BETWEEN 20000 AND 50000 THEN 1 ELSE 0 END) AS "Average Salary",
        SUM(CASE WHEN income > 50000 THEN 1 ELSE 0 END) AS "High Salary"
    FROM Accounts
)

SELECT 'Low Salary' AS category, "Low Salary" AS accounts_count FROM CTE_SUM
UNION ALL
SELECT 'Average Salary' AS category, "Average Salary" AS accounts_count FROM CTE_SUM
UNION ALL
SELECT 'High Salary' AS category, "High Salary" AS accounts_count FROM CTE_SUM;

