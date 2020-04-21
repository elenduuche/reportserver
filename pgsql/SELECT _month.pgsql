SELECT *
FROM payments
WHERE (extract (month FROM createdon) = 4)
AND (extract (year from createdon) = 2021)