'use strict'

const
express = require('express'),
app = express(),
port = 3000,
cors = require('cors')(),
route = require('./api/routes/route');

app.use(cors);
//app.use(express.json({limit: '50mb'}));
//app.use(express.urlencoded({limit: '50mb'}));
app.listen(port);
route(app);


console.log("Server started on " + port);

