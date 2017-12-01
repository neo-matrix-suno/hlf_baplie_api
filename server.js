'use strict'

const
express = require('express'),
app = express(),
port = 3000,
cors = require('cors')(),
route = require('./api/routes/route');

app.use(cors);
app.listen(port);
route(app);


console.log("Server started on " + port);

