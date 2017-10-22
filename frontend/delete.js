var http = require("http");

var q = {
    query: {email:"jithinjoyson1997@gmail.com"}
};

var payload = JSON.stringify(q);

var options = {
    host: '45.55.182.73',
    port: 8000,
    path: '/account/loaner',
  method: 'POST',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var options = {
    host: '45.55.182.73',
    port: 8000,
    path: '/account/donor',
  method: 'POST',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var request = http.request(options);

request.write(payload);
request.end();