const dashboard = require('./');
const moment = require('moment');

let timestamp = null;
let cachedData = null;

function fetchDashboardData() {
  return dashboard.getNamespaces().then((request) => {
    const excludeNamespaces = process.env.EXCLUDE_NAMESPACES ? process.env.EXCLUDE_NAMESPACES.split(',') : [];
    const namespaces = request.body.items
      .map(item => item.metadata.name)
      .filter(namespace => excludeNamespaces.indexOf(namespace) === -1);

    return Promise.all(
      namespaces.map(namespace => dashboard.getPodsFromNamespace((namespace)))
    ).then((podsRequest) => namespaces.map((namespace, index) => ({
        namespace,
        pods: podsRequest[index].body.items.map(item => ({
          name: item.metadata.name,
          lastCreation: moment(item.metadata.creationTimestamp).fromNow(),
          images: item.spec.containers.map(container => container.image),
        })),
      })).filter(namespace => namespace.pods.length > 0)
    );
  });
}

function fetch() {
  if (cachedData == null || moment().unix() - timestamp > 10) {
    return fetchDashboardData().then((data) => {
      timestamp = moment().unix();
      cachedData = data;
      return data;
    });
  }

  return Promise.resolve(cachedData);
}

module.exports = fetch;
