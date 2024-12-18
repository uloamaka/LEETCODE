# Write your MySQL query statement below
-- GROUP BY MONTH AND COUNTRY, 
SELECT 
    DATE_FORMAT( trans_date, '%Y-%m' ) month,
    country,
    COUNT(state) trans_count,
    SUM(state = 'approved') approved_count,
    SUM(amount) trans_total_amount,
    SUM((state = 'approved') * amount)  approved_total_amount
FROM Transactions 
GROUP BY month, country