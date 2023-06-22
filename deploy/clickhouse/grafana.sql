
SELECT toStartOfMinute(time_iso8601) as time, concat(request_method, ' ', uri) as metric, COUNT(1) as Req
FROM test_log
WHERE $__timeFilter(time) AND http_host IN ($domain) AND metric IN (
  SELECT concat(request_method, ' ', uri) as metric FROM test_log 
  WHERE $__timeFilter(time_iso8601) AND http_host IN ($domain)
  GROUP BY metric ORDER BY COUNT(1) DESC limit 10
)
GROUP BY metric, time
ORDER BY time, Req DESC
LIMIT 1000;

