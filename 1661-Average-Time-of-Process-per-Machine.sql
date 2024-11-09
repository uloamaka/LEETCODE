-- Write your PostgreSQL query statement below
-- Select the machine_id with the average time, to get the complete time subract the start from end timestamp, then get the avg()
-- A1 - END 
-- A0 - START
SELECT 
    machine_id,
    round(avg(subQ.processing_time)::numeric, 3) AS processing_time
FROM (    
SELECT 
    A1.machine_id, 
    A1.process_id,
    A1.timestamp AS A1,
    A0.timestamp AS A0,
    (A1.timestamp - A0.timestamp) AS processing_time
FROM Activity A1
INNER JOIN Activity A0 ON A1.machine_id = A0.machine_id 
AND A1.process_id = A0.process_id
AND A1.activity_type = 'end'
AND A0.activity_type = 'start'
) subQ
GROUP BY subQ.machine_id;
