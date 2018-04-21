const dashboard = require('./');
const moment = require('moment');

function fetch() {
  return dashboard.getNamespaces().then((request) => {
    const namespaces = request.body.items.map(item => item.metadata.name);

    return Promise.all(
      namespaces.map(namespace => dashboard.getPodsFromNamespace((namespace)))
    ).then((podsRequest) => namespaces.map((namespace, index) => ({
        namespace,
        pods: podsRequest[index].body.items.map(item => ({
          name: item.metadata.name,
          lastCreation: moment(item.metadata.creationTimestamp).fromNow(),
          images: item.spec.containers.map(container => container.image),
        })),
      }))
    );
  });
}

module.exports = fetch;
