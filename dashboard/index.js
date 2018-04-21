const Client = require('kubernetes-client').Client;
const config = require('kubernetes-client').config;
const client = new Client({ config: config.fromKubeconfig() });

async function getNamespaces() {
  await client.loadSpec();
  return await client.api.v1.namespaces.get();
}

async function getPodsFromNamespace(namespace) {
  await client.loadSpec();
  return client.api.v1.namespaces(namespace).pods.get();
}

module.exports = {
  getNamespaces,
  getPodsFromNamespace,
};