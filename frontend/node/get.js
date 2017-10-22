var http = require("http");

var d = '59ec330eb390353c953a1615';

var optionsB = {
  timeout: 10000,
  host: '162.243.160.168',
  port: 8000,
  path: '/account/donors/'+d,
  method: 'GET',
  headers: {
    'Content-Type': 'application/json'
  }
};

var optionsA = {
  host: '162.243.160.168',
  port: 8000,
  path: '/account/loaners/'+d,
  method: 'GET',
  headers: {
    "Content-Type": "application/json"
  }
};

var req = http.request(options, function(res) {
  //res.setEncoding("UTF-8");
  var response = "";
  res.on("data", function(chunk){
    response += chunk;
  });

  res.on("end", function(){
    console.log(response);
  });
  var data = JSON.stringify(response);
});
req.end();