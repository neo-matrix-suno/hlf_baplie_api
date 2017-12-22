'use strict'

const path = require('path');

/*
exports.getAll = function(req, res) {
	const hfc_code = require('../../hfc_codes/get_all');
	let options = {
	    wallet_path: path.join(__dirname, '../../creds'),
	    user_id: 'MAERSK',
	    channel_id: 'mychannel',
	    chaincode_id: 'baplie',
	    network_url: 'grpc://localhost:7051',
	};
	return hfc_code.query(options, res);
}
*/
exports.getItem = function(req, res) {
	const hfc_code = require('../../hfc_codes/get_item');
	let options = {
	    wallet_path: path.join(__dirname, '../../creds'),
	    //user_id: req.params.peerId,
	    //user_id: 'MAERSK',
	    user_id: 'MAERSK',
	    channel_id: 'mychannel',
	    chaincode_id: 'baplie',
   	    command: 'queryBap',
   	    args: [req.params.key],
	    peer_url: 'grpc://localhost:7051',
		event_url: 'grpc://localhost:7053',
		orderer_url: 'grpc://localhost:7050',
	};
	return hfc_code.invoke(options, res);
}

exports.getItemWithParam = function(req, res) {
	const hfc_code = require('../../hfc_codes/get_item');
	//const condition = JSON.stringify(req.body.condition);
	const condition = req.body.condition;
	console.log("AAA: %s",condition);
	let options = {
	    wallet_path: path.join(__dirname, '../../creds'),
	    //user_id: req.params.peerId,
	    user_id: 'MAERSK',
	    channel_id: 'mychannel',
	    chaincode_id: 'baplie',
   	    command: 'queryWithParam',
   	    args: [condition],
   	    //args: [req.params.voyage],
   	    //args: [req.params.voyage],
	    peer_url: 'grpc://localhost:7051',
		event_url: 'grpc://localhost:7053',
		orderer_url: 'grpc://localhost:7050',
	};
	return hfc_code.invoke(options, res);
}


exports.putItem = function(req, res) {
	const hfc_code = require('../../hfc_codes/put_item');
	let options = {
	    wallet_path: path.join(__dirname, '../../creds'),
	    user_id: req.params.peerId,
	    //user_id: 'HJSTML',
	    channel_id: 'mychannel',
	    chaincode_id: 'baplie',
	    command: 'createBap',
   	    args: [req.params.key].concat(req.body.args),
		peer_url: 'grpc://localhost:7051',
		event_url: 'grpc://localhost:7053',
		orderer_url: 'grpc://localhost:7050',
	};
	return hfc_code.invoke(options, res);
}
