'use strict'

const request = require('request');
const options = {
	url: 'http://localhost:3000/asset/TestAdmin/CAR20',
	//body: {'args': ["1981", 'SIN', 'DiamondBlue', 'Bright']},
	//body: {AAA:["1981", 'SIN', 'DiamondBlue', 'Bright']},
	body: {'args': [colour:"1981", make:'SIN', model:'DiamondBlue', owner:'Bright']},
	json: true
};

request.put(options, function(err, res, body) {
	if (err) {
		console.log('code: ' + err);
	}

	if (res) {
		console.log(res.body);
	}
});

