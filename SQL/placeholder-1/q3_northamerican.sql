-- North American
--  Indicate if an order's `ShipCountry` is in North America. For our purposes, this is 'USA', 'Mexico', 'Canada'.
SELECT
    ID,
    ShipCountry,
    (
        CASE
            WHEN ShipCountry IN ('USA', 'Mexico', 'Canada')
                THEN (SELECT 'NorthAmerica')
            ELSE
                (SELECT 'OtherPlace')
        END
    )
FROM 'Order' WHERE ID >= 15445 LIMIT 20;
