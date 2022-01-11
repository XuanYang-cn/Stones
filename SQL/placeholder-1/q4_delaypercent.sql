-- DelayPercent
-- For each `Shipper`, find the percentage of orders which are late.
-- Details: An order is considered late if `ShippedData > RequiredDate`.
-- Print the following format, order by descending percentage, rounded to the nearest hundredths, like `United Package|23.44`
SELECT
    s.CompanyName,
    COUNT(*),
    ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM 'Order' AS o, Shipper as s WHERE o.ShipVia = s.Id GROUP BY s.ID), 2) as percentage
FROM 'Order' AS o, Shipper as s
WHERE o.ShippedDate > o.RequiredDate AND o.ShipVia = s.Id
GROUP BY s.Id
ORDER BY percentage DESC;
