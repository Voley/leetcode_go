SELECT name
FROM Customer
WHERE NOT referee_id = 2 OR ISNULL(referee_id)
