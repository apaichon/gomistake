# gomistake
POC the mistake way in golang

## TEST STEPS
1. Start mq-server.
```sh
go run main.go 
```
The http server run on port 8888.
2. Go to load-test project and install node libraries.
```sh
cd load-test
npm install
```
3. Set Load test capacity on file test-360-deposit.js
```javascript
 // Set a title for the load test
    title: 'Deposit 360 Creation Load Test',
    connections:10, // total connection simulate 
    duration: 300, // duration to test in seconds
    amount:1_000 // amount of records to load test (1000/10) = 100 records per connection in 5 mins
  
  };
```
4. Run load test.
```sh
node test-360-deposit.js
```
5. See load test performance at folder /reports
6. Monitor memory usage in console or Activity Monitoring.
