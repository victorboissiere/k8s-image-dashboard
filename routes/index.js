const express = require('express');
const router = express.Router();
const fetchDashboard = require("../dashboard/fetch");
const dashboard = require("../dashboard");

router.get('/', (req, res, next) => {
  return fetchDashboard().then((dashboardData) => {
    res.render('index.html', { data: dashboardData });
  }).catch((error) => {
    res.render('error.html', { error: JSON.stringify(error, null, 2) });
  })
});

router.get('/nodes', (req, res, next) => {
  return dashboard.getNodes().then((data) => {
    console.log(JSON.stringify(data, null, 2));
    res.render('nodes.html', {
      nodes: data.body.items.map(item => ({
        name: item.metadata.name,
        labels: Object.keys(item.metadata.labels).map(label => `${label}: ${item.metadata.labels[label]}`),
        annotations: Object.keys(item.metadata.annotations).map(annotation => `${annotation}: ${item.metadata.annotations[annotation]}`),
        ip: item.status.addresses.find(addr => addr.type == 'InternalIP').address,
        kubeVersion: item.status.nodeInfo.kubeletVersion,
      })),
    });
  }).catch((error) => {
    res.render('error.html', { error: JSON.stringify(error, null, 2) });
  });
});

router.get('/health', (req, res, next) => {
  return res.status(200).send('Healthy');
});

module.exports = router;
