//express_demo.js 文件
var express = require('express');
var app = express();
var request = require('superagent')

var Fabric_Client = require('fabric-client');
var Fabric_CA_Client = require('fabric-ca-client');
var fabric_client = new Fabric_Client();
var fabric_ca_client = null;
var admin_user = null;
var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
var channel;
var peer;

app.get('/', function (req, res) {
  console.log("hello")
  res.status(200).send("working");
});

app.post('/api/uppic', function (req, res) {
  const sreq = request.post(HOST + req.originalUrl);
  req.body.app = 'blockchain'
  req.body.key = 'f6070v96i4mI4P5K'
  sreq.set('Content-Type', 'application/json')
      .send(req.body)

  sreq.pipe(res);
  sreq.on('end', function (error, res) {
    if (error) {
      console.log('post end ' + error);
      return;
    }
  });
});

app.post('/api/upfile', function (req, res) {
  const sreq = request.post(HOST + req.originalUrl);
  req.body.app = 'blockchain'
  req.body.key = 'f6070v96i4mI4P5K'
  sreq.set('Content-Type', 'application/json')
      .send(req.body)

  sreq.pipe(res);
  sreq.on('end', function (error, res) {
    if (error) {
      console.log('post end ' + error);
      return;
    }
  });
});

app.get('/api/getpic', function (req, res) {

  const request = {
		chaincodeId: 'xincc',
		fcn: 'isbuy',
		args: ['query,'+ req.user+','+req.pwd+','+req.car+','+req.fhash]
	};

	// send the query proposal to the peer
 channel.queryByChaincode(request).then((query_responses) => {
   if(query_responses[0].toString()=="1")
   {
    var sreq = request.get("http://upload.xin.com"+req.query.picurl)
    sreq.pipe(res);
    sreq.on('end', function (error, res) {
      if (error) {
        console.log('get end ' + error);
        return;
      }
    });
   }else{
     res.status(404).send();
   }
 });
})

var server = app.listen(8081, function () {

  Fabric_Client.newDefaultKeyValueStore({ path: store_path
  }).then((state_store) => {
        // assign the store to the fabric client
        fabric_client.setStateStore(state_store);
        var crypto_suite = Fabric_Client.newCryptoSuite();
        // use the same location for the state store (where the users' certificate are kept)
        // and the crypto store (where the users' keys are kept)
        var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
        crypto_suite.setCryptoKeyStore(crypto_store);
        fabric_client.setCryptoSuite(crypto_suite);
        var	tlsOptions = {
          trustedRoots: [],
          verify: false
        };
        // be sure to change the http to https when the CA is running TLS enabled
        fabric_ca_client = new Fabric_CA_Client('http://172.16.40.41:8054', tlsOptions , 'rca-org1', crypto_suite);
    
        // first check to see if the admin is already enrolled
        return fabric_client.getUserContext('user-admin-org1', true);
  }).then((user_from_store) => {
    if (user_from_store && user_from_store.isEnrolled()) {
        console.log('Successfully loaded admin from persistence');
        admin_user = user_from_store;
        return null;
    } else {
        // need to enroll it with CA server
        return fabric_ca_client.enroll({
          enrollmentID: 'user-admin-org1',
          enrollmentSecret: 'user-admin-org1pw'
        }).then((enrollment) => {
          console.log('Successfully enrolled admin user "user-admin"');
        }).catch((err) => {
          console.error('Failed to enroll and persist user-admin. Error: ' + err.stack ? err.stack : err);
          throw new Error('Failed to enroll user-admin');
        });
    }
}).then(() => {
  console.log('Assigned the admin user to the fabric client ::' + admin_user.toString());
}).catch((err) => {
  console.error('Failed to enroll user-admin: ' + err);
});

// setup the fabric network
channel = fabric_client.newChannel('mychannel');
peer = fabric_client.newPeer('grpc://172.16.40.41:7051');
channel.addPeer(peer);

var host = server.address().address
var port = server.address().port

console.log("http://%s:%s", host, port)
})
