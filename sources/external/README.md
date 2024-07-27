# Field Enablement 1
To quickly test how the app works with a vm run the following script to install the Open Telemetry Collector on your vm:

### Sample Command
```
curl ${OBSERVE_COLLECTION_ENDPOINT?}v1/http \
 -H "Authorization: Bearer ${OBSERVE_TOKEN?}" \
 -H "Content-type: application/json" \
 -d '{"ir": "app maker"}'
```

