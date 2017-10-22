var http = require("http");

var d = '59ec1b72b390353c953a160f';


var optionsA = {
    host: '162.243.160.168',
    port: 8000,
    path: '/account/loaners/'+d,
  method: 'DELETE',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var optionsB = {
      host: '162.243.160.168',
      port: 8000,
      path: '/account/donors/'+d,
  method: 'DELETE',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var request = http.request(optionsA);

request.end();