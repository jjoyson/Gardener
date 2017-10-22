var http = require('http');

    var data = JSON.stringify({
        amount: 10000.50,
        credit_card_number: "4485218429391984",
        });

    var optionsA = {
        host: '45.55.182.73',
        port: 8000,
        path: '/payments/loaners/59ec5cb4b390353c953a1625',
      method: 'POST',
     headers:{
        'Content-Type': 'application/x-www-form-urlencoded',
        }
    };

    var optionsB = {
        host: '45.55.182.73',
        port: 8000,
        path: '/payments/donors',
      method: 'POST',
     headers:{
        'Content-Type': 'application/x-www-form-urlencoded',
        }
        
    };

    var req = http.request(optionsA, function(res) {
        res.setEncoding('utf8');
        res.on(data, function (chunk) {
            console.log("body: " + chunk);
        });
    });
    console.log(data);
    req.write(data);
    req.end();