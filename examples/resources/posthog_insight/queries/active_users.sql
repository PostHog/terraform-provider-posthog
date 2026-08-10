SELECT
    toDate(timestamp) AS day,
    count(DISTINCT person_id) AS active_users
FROM events
WHERE timestamp >= now() - INTERVAL 7 DAY
GROUP BY day
ORDER BY day
