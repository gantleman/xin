var express = require('express');
var app = express();
var bodyParser = require('body-parser');
app.use(bodyParser.json()); // for parsing application/json
var Web3 = require('web3');

var sqlite3 = require('sqlite3').verbose();
var db = new sqlite3.Database('mydb');
web3 = new Web3();

var xintokenContract;
var xintoken;
var ethserv = 'http://192.168.56.104:8547';
web3.setProvider(new web3.providers.HttpProvider(ethserv));

var coinbase = web3.eth.coinbase;
console.log(coinbase);

var balance = web3.eth.getBalance(coinbase);
console.log(balance.toString(10));

xintokenContract = web3.eth.contract([{"constant":true,"inputs":[{"name":"car","type":"address"},{"name":"fhash","type":"address"}],"name":"geturl","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"car","type":"address"},{"name":"fhash","type":"address"},{"name":"url","type":"string"},{"name":"price","type":"uint256"},{"name":"stamp","type":"uint256"}],"name":"add_repair","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"seller","type":"address"},{"name":"value","type":"uint256"}],"name":"selfcollect","outputs":[{"name":"","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"index","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"buyer","type":"address"},{"name":"value","type":"uint256"}],"name":"addcoin","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"info","outputs":[{"name":"url","type":"string"},{"name":"price","type":"uint256"},{"name":"stamp","type":"uint256"},{"name":"balance","type":"uint256"},{"name":"seller","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"x","type":"address"}],"name":"address2str","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"car","type":"address"},{"name":"fhash","type":"address"},{"name":"buyer","type":"address"}],"name":"isbuy","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"coin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"uint256"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"car","type":"address"},{"name":"fhash","type":"address"}],"name":"buy","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[],"name":"getselfcoin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"car","type":"address"},{"name":"fhash","type":"address"},{"name":"irate","type":"uint256"},{"name":"comment","type":"string"}],"name":"add_rate","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[],"name":"error","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"seller","type":"address"}],"name":"getcoin","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"collect","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"car","type":"address"}],"name":"getaddress","outputs":[{"name":"rets","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"i","type":"uint256"}],"name":"uint2str","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[],"name":"admin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"padmin","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]);
xintoken = xintokenContract.at("0x0a114949c6a5dcdd9243bf264acf74ed180336c4");

console.log('web3 connecting')

function isConnected(){
	if(!web3.isConnected()){
		web3.setProvider(new web3.providers.HttpProvider(ethserv));
		return web3.isConnected();
	}else
		return 1;
}

app.get('/eth', function (req, res) {
	res.send('Hello World');

	var coinbase = web3.eth.coinbase;
	console.log(coinbase);

	var balance = web3.eth.getBalance(coinbase);
	console.log(balance.toString(10));
})

app.get('/', function (req, res) {
   res.status(200).send('ok');
})


app.post('/eth/gasCharge',function (req, res) {
	console.log('/eth/gasCharge');

	if (!isConnected()){
		console.log('eth fail');
		res.status(201).send('eth fail');
		return;
	}

	addr = req.body.addr;

	console.log("body:"+req.body.addr);

	if (req.body.addr == 'undefined'){
		res.status(201).send('addr fail');
		return;
	}

	var balance = web3.fromWei(web3.eth.getBalance(addr), "ether")
	if (balance < 1)
	{
		web3.eth.sendTransaction({from: web3.eth.coinbase, to: addr, value: web3.toWei(10,"ether")}, 
		function(err, hash){
			 if (err){
					console.log('FAIL on sendTransaction ' + err);
					res.status(201).send('add fail');
				}else{
					res.status(200).send('add ok');
				}
			});
	}else
		res.status(200).send('no add ok');
});

app.post('/eth/coinInit', function (req, res, next) {
	console.log('/eth/coinInit');

	if (!isConnected()){
		console.log('eth fail');
		res.status(201).send('eth fail');
		return;
	}

	console.log("body:"+req.body.addr);
	console.log("coinbase:"+web3.eth.coinbase);

	if (req.body.addr == 'undefined'){
		res.status(201).send('addr fail');
		return;
	}

	db.serialize(function() {
		db.run("CREATE TABLE IF NOT EXISTS t(a TEXT PRIMARY KEY)",
			function(err){
            if (err){
                console.log('FAIL on creating table ' + err);
				res.status(201).send('fail');
            }else{
				db.run("INSERT INTO t VALUES (?)", req.body.addr, function(err){
					if (err){
						 console.log('FAIL on INSERT INTO ' + err);
						 res.status(201).send('fail');
					}else{
						xintoken.addcoin.sendTransaction(req.body.addr, 1000, {from:web3.eth.coinbase});
						console.log('ok');
						res.status(200).send('ok');
						return;
					}
				});
			}
        });
	});
});


var server = app.listen(8080, function () {

  var host = server.address().address;
  var port = server.address().port;

  console.log("应用实例，访问地址为 http://%s:%s", host, port);
})

