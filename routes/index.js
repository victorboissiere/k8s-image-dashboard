const express = require('express');
const router = express.Router();
const fetchDashboard = require("../dashboard/fetch");

router.get('/', (req, res, next) => {
  return fetchDashboard().then((dashboardData) => {
    res.render('index.html', { data: dashboardData });
  }).catch((error) => {
    res.render('error.html', { error: JSON.stringify(error, null, 2) });
  })
});

router.get('/health', (req, res, next) => {
  return res.status(200).send('Healthy');
});

module.exports = router;
