EFK Stack (Elasticsearch-Flentd-Kibana) docker-compose
================================================
An example docker-compose.yaml file and configurations file for testing collecting(fluent-bit), storing(elasticsearch) and visualizing(Kinbana) of Docker container logs with automated Fluent-bit Logging-driver.  

### Usage

```bash
docker-compose up -d
```

### Dashboard  

    - Elasticsearch: 
```
    http://localhost:9200/  
    http://localhost:9200/_cat/indices
```

    - Kibana: 
```
    http://localhost:5601/  
```

    - Sample Web Server with Httpd for log access_log from stdout: 
```
    http://localhost:8082/  
```


