var http = require("http");
var https = require("https");

var options = {
    host: '45.55.182.73',
    port: 8000,
    path: '/account/donor',
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
};

var options = {
    host: '45.55.182.73',
    port: 8000,
    path: '/account/loaner',
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
};

var req = http.get(options, function(res) {
    console.log('STATUS: ' + res.statusCode);
    console.log('HEADERS: ' + JSON.stringify(res.headers));