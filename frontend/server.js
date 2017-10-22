var express = require('express');
var http = require('http');
var bodyParser = require('body-parser');
var app     = express();
var fs = require('fs');

//Note that in version 4 of express, express.bodyParser() was
//deprecated in favor of a separate 'body-parser' module.
app.use(bodyParser.urlencoded({ extended: true })); 

//app.use(express.bodyParser());

app.post('/myaction', function(req, res) {
    
        var data = JSON.stringify({
            first_name: req.body.first_name,
            last_name: req.body.last_name,
            email: req.body.email,
            password: req.body.password,
            address: {
            street_number: req.body.street_number,
            street_name: req.body.street_name,
            city: req.body.city,
            state: req.body.state,
            zip: req.body.zip
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
        
        var options;
        var a = req.body.h1;
        if(req.body.h1)
            options = optionsB;
        else
            options = optionsA;
        console.log("made it!");
        var req = http.request(options, function(res) {
            res.setEncoding('utf8');
            res.on(data, function (chunk) {
                console.log("body: " + chunk);
            });
        });
        req.write(data);
        req.end();

        if(a){
            response.writeHead(301, {
            'Location': 'http://45.55.182.73:8000/lenderSignedln.html'});
          response.end();
        }
        else
        { response.writeHead(301, {
            'Location': 'http://45.55.182.73:8000/borrowerSignedln.html'});
          response.end();
        }
});

app.listen(8080, function() {
  console.log('Server running at http://127.0.0.1:8080/');
});
