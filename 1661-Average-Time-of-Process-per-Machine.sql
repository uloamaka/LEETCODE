-- Write your PostgreSQL query statement below
-- Select the machine_id with the average time, to get the complete time subract the start from end timestamp, then get the avg()
-- A1 - END 
-- A0 - START
SELECT 
    A1.machine_id,
    round(avg(A1.timestamp - A0.timestamp)::numeric, 3) AS processing_time
FROM Activity A1
INNER JOIN Activity A0 ON A1.machine_id = A0.machine_id 
WHERE A0.activity_type = 'start' AND A1.activity_type = 'end'
GROUP BY A1.machine_id;
