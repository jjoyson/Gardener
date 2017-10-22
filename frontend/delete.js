var http = require("http");

var d = '59ec1b72b390353c953a160f';


var optionsA = {
    host: '45.55.182.73',
    port: 8000,
    path: '/account/loaners/'+d,
  method: 'DELETE',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var optionsB = {
      host: '45.55.182.73',
      port: 8000,
      path: '/account/donors/'+d,
  method: 'DELETE',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var request = http.request(options2);

request.end();