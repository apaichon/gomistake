const autocannon = require('autocannon');
const fs = require('fs');
const csvParser = require('csv-parser');

// CSV file path
const csvFilePath = './data/deposit.csv';

// API endpoint
const apiEndpoint = 'http://localhost:5055/api/deposits'; // Assuming the API endpoint for creating users is /api/users

// Read data from CSV and send POST requests

fs.createReadStream(csvFilePath)
  .pipe(csvParser())
  .on('data', (row) => {
    // Extract data from CSV row
    const { AccountID, Amount, DepositDate, CreatedBy} = row;

    // Define POST request payload
    const payload = JSON.stringify({ AccountID, Amount, DepositDate, CreatedBy });

    // Options for autocannon
    const opts = {
      url: apiEndpoint,
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: payload,
        // Set a title for the load test
        title: 'Deposit Load Test'
    };

    // Send POST request using autocannon
    autocannon(opts, (err, result) => {
      if (err) {
        console.error(err);
        return;
      }
      console.log(result);
    });
  })
  .on('end', () => {
    console.log('CSV file processed');
  });
