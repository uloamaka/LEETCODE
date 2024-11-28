-- Write your PostgreSQL query statement below
-- SELECT  
--     teacher_id,
--     COUNT(DISTINCT subject_id) AS cnt
-- FROM Teacher 
-- GROUP BY 1;
SELECT teacher_id, COUNT(subject_id) AS cnt
FROM (SELECT DISTINCT ON (subject_id, teacher_id) * FROM Teacher)
GROUP BY teacher_id