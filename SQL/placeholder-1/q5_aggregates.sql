-- Aggregates
-- Details: Get the number of products, average unit price(rounded to 2 decimal places), minimum unit price,
-- maximum unit price, and total units on order for categories containing greater than 10 products.
-- Order by `Category Id`. Your output should look like `Beverages|12|37.98|4.5|263.5|60`

SELECT
    c.CategoryName,
    COUNT(*),
    ROUND(AVG(p.UnitPrice), 2),
    MIN(p.UnitPrice),
    MAX(p.UnitPrice),
    SUM(p.UnitsOnOrder)
FROM Product AS p, Category as c
WHERE c.ID = p.CategoryId AND
    (
        SELECT COUNT(*) > 10
        FROM Product AS p, Category AS c
        WHERE c.ID = p.CategoryId
        GROUP BY p.CategoryId
    )
GROUP BY p.CategoryId
ORDER BY p.CategoryId;
