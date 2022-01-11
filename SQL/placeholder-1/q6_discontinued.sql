-- q6
-- Discontinued
-- For each of the 8 discontinued products in the database, which custormer make the first ever order for the
-- product? Output the customer's `CompanyName` and `ContactName`
-- Details: Print the following format, order by `ProductName` alphabetically:
-- `Alic Mutton|Consolidated Holdings|Elizabeth Brown`
SELECT Productname, CompanyName, Contactname FROM (
    SELECT
        p.ProductName,
        c.CompanyName,
        c.ContactName,
        RANK() OVER ( PARTITION BY p.ProductName
                            ORDER BY o.OrderDate) AS r
    FROM Product AS p, 'Order' AS o, 'Customer' AS c, 'OrderDetail' AS od
    WHERE
        p.Discontinued = 1 AND
        od.ProductId = p.Id AND
        od.OrderId = o.Id AND
        o.CustomerId = c.Id

    ) AS ranking
WHERE ranking.r = 1
ORDER BY ranking.ProductName
;
