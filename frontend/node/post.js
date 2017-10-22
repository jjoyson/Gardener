var http = require('http');

    var data = JSON.stringify({
        first_name: "Tom",
        last_name: "Cruz",
        email: "tommyboy@gmail.com",
        password: "password12",
        address: {
        street_number: "13651",
        street_name: "NW 131st PL",
        city: "New York City",
        state: "New York",
        zip: "75895"
        }
        });

    var optionsA = {
        host: '45.55.182.73',
        port: 8000,
        path: '/account/loaners',
      method: 'POST',
     headers:{
        'Content-Type': 'application/x-www-form-urlencoded',
        }
    };

    var optionsB = {
        host: '45.55.182.73',
        port: 8000,
        path: '/account/donors',
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
    req.write(data);
    req.end();