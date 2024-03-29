-- q7 order lags
-- For the first 10 orders by `CustomerId BLONP`: get the Order's `Id`, `OrderDate`, previous
-- `OrderDate`, and difference between the previous and current. Return results ordered by
-- `OrderDate` (ascending)
-- Details: The "previous" `OrderDate` for the first order should default to itself(lag time = 0)
-- Use the `julianday()` function for date arithmetic.
-- Use lag(expr, offset, default) for grabbing previous dates.
-- Please round the lage time to the nearest hundredth, formatted like:
-- 17361|2012-09-19 12:13:21|2021-09-18 22:37:15|0.57
