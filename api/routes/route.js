'use strict'

module.exports = function(app) {

	const controller = require('../controllers/controller');

	const bodyParser = require('body-parser');
	app.use(bodyParser.urlencoded({extended: true}));
	app.use(bodyParser.json({limit: '50mb'}));
	//app.use(express.json({limit: '50mb'}));

	app.route('/asset')
		//.get(controller.getAll);
	
	app.route('/asset/:peerId/:key')
		.get(controller.getItem)
		.put(controller.putItem);

        app.route('/asset/:peerId/search/')
                .post(controller.getItemWithParam);

};
