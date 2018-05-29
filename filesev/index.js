//express_demo.js 文件
var express = require('express');
var app = express();
var request = require('superagent')

app.get('/', function (req, res) {
  console.log("hello")
});

app.post('/api/uppic', function (req, res) {
  const sreq = request.post(HOST + req.originalUrl);
  req.body.app = 'blockchain'
  req.body.key = 'f6070v96i4mI4P5K'
  sreq.set('Content-Type', 'application/json')
      .send(req.body)

  sreq.pipe(res);
  sreq.on('end', function (error, res) {
    console.log('post end ' + error);
  });
});

app.get('/api/getpic', function (req, res) {
  console.log("get+"+req.query.picurl);

  var sreq = request.get("http://upload.xin.com"+req.query.picurl)
  sreq.pipe(res);
  sreq.on('end', function (error, res) {
    console.log('get end ' + error);
  });
})

var server = app.listen(8081, function () {

  var host = server.address().address
  var port = server.address().port

  console.log("http://%s:%s", host, port)
})
