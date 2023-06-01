# Test prometheus alerts

Notifications sent by AlertManager can be tested by creating a test alert:

```bash
curl -H 'Content-Type: application/json' -d '[{"labels":{"alertname":"testalert"}}]' https://alertmanager/api/v1/alerts
```

    #prometheus #alerts #monitoring
