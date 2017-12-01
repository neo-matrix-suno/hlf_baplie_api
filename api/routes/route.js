'use strict'

module.exports = function(app) {

	const controller = require('../controllers/controller');

	const bodyParser = require('body-parser');
	app.use(bodyParser.urlencoded({extended: true}));
	app.use(bodyParser.json());

	app.route('/asset')
		//.get(controller.getAll);
	
	app.route('/asset/:peerId/:key')
		.get(controller.getItem)
		.put(controller.putItem);

	app.route('/asset/:peerId/carrier/:key')
		.get(controller.getItemByCarrier);

	app.route('/asset/:peerId/search/')
		.post(controller.getItemWithParam);

	app.route('/asset/:peerId/vessel/:key')
		.get(controller.getItemByVessel);

	app.route('/asset/:peerId/voyage/:key')
		.get(controller.getItemByVoyage);

	app.route('/asset/:peerId/vsldate/:key')
		.get(controller.getItemByVsldate);

	app.route('/asset/:peerId/snddate/:key')
		.get(controller.getItemBySnddate);

	app.route('/asset/:peerId/equiment/:key')
		.get(controller.getItemByEquiment);

};
