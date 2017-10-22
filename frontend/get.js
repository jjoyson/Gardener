var http = require('http');

var options = {
    host: '45.55.182.73',
    port: 8000,
    path: '/account/loaner',
  method: 'GET',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    }
};

var q = {email:"jithinjoyson1997@gmail.com"};

var options = {
    uri:'http://45.55.182.73:8000/account/leanders/2',
    host: '45.55.182.73',
    port: 8000,
    path: '/account/donor',
  method: 'GET',
 headers:{
    'Content-Type': 'application/x-www-form-urlencoded',
    },
    qs:q
};

var req = http.get(options, function(res) {
  console.log('STATUS: ' + res.statusCode);
  console.log('HEADERS: ' + JSON.stringify(res.headers));

  // Buffer the body entirely for processing as a whole.
  var bodyChunks = [];
  res.on('data', function(chunk) {
    // You can process streamed parts here...
    bodyChunks.push(chunk);
  }).on('end', function() {
    var body = Buffer.concat(bodyChunks);
    console.log('BODY: ' + body);
    // ...and/or process the entire body here.
  })
});

req.on('error', function(e) {
  console.log('ERROR: ' + e.message);
});