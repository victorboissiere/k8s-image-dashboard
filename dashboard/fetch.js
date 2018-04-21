const dashboard = require('./');
const moment = require('moment');

let timestamp = null;
let cachedData = null;

function getWarningMessageForImage(image) {
  if (image.indexOf(':') === -1 || image.indexOf(':latest') !== -1) {
    return 'This image does not use a specific version';
  }

  return null;
}

function getRepositoryURLForImage(image) {
  if (process.env.REPO_REGEX) {
    const repoRegex = process.env.REPO_REGEX.split('|');
    const regexp = new RegExp(repoRegex[0]);
    if (image.match(regexp)) {
      return image.replace(regexp, repoRegex[1]);
    }
  }
  return null;
}

function getImages(pod) {
  return pod.spec.containers.map(container => ({
    name: container.image,
    warning: getWarningMessageForImage(container.image),
    url: getRepositoryURLForImage(container.image),
  }));
}

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
          images: getImages(item),
          lastCreation: moment(item.metadata.creationTimestamp).fromNow(),
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
