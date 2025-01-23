# Write your MySQL query statement below
-- GROUP BY TEACHER_ID, SELECT THE DISTINCT SUBJECT
SELECT 
    teacher_id,
    COUNT(DISTINCT subject_id) AS cnt
FROM 
    Teacher
GROUP BY teacher_id;