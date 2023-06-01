# Test prometheus alerts

Notifications sent by AlertManager can be tested by creating a test alert:

```bash
curl -H 'Content-Type: application/json' -d '[{"labels":{"alertname":"testalert"}}]' https://alertmanager/api/v1/alerts
```

## Related

- https://fabianlee.org/2022/07/03/prometheus-sending-a-test-alert-through-alertmanager/

    #prometheus #alerts #monitoring
