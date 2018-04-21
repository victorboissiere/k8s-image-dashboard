const express = require('express');
const path = require('path');
const cookieParser = require('cookie-parser');
const logger = require('morgan');
const mustacheExpress = require('mustache-express');

const indexRouter = require('./routes/index');

const app = express();

app.use(logger('dev', { skip: (req, res) =>  req.url === '/health' || req.url.startsWith('/images') }));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.set('views', __dirname + '/views');
app.engine('html', mustacheExpress());
app.set('view engine', 'mustache');

app.use('/', indexRouter);

module.exports = app;
