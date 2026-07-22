INSERT INTO customer_profiles
(
    user_id,
    first_name,
    last_name,
    country
)
SELECT
    id,
    first_name,
    last_name,
    'USA'
FROM users
WHERE NOT EXISTS
(
    SELECT 1
    FROM customer_profiles
    WHERE customer_profiles.user_id = users.id
);