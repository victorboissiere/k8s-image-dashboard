const Client = require('kubernetes-client').Client;
const config = require('kubernetes-client').config;
const client = new Client({ config: config.fromKubeconfig(), version: '1.9' });

function getNamespaces() {
  return client.api.v1.namespaces.get().then((results) => {
    return results;
  });
}

function getPodsFromNamespace(namespace) {
  return client.api.v1.namespaces(namespace).pods.get().then((results) => {
    return results;
  });
}

module.exports = {
  getNamespaces,
  getPodsFromNamespace,
};