const Client = require('kubernetes-client').Client;
const config = require('kubernetes-client').config;
const client = new Client({ config: process.env.LOCAL ? config.fromKubeconfig() : config.getInCluster() });

function getNamespaces() {
  return client.loadSpec().then(() => client.api.v1.namespaces.get().then(results => results));
}

function getPodsFromNamespace(namespace) {
  return client.loadSpec().then(() => client.api.v1.namespaces(namespace).pods.get().then(results => results));
}

module.exports = {
  getNamespaces,
  getPodsFromNamespace,
};