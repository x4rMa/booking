'use strict';

angular.module('bookingServices', ['ngResource']);

angular.module('bookingServices').factory('AuthService', ['$log', function ($log) {
  var service = {};

  service.verifyLogin = function (user) {
    var result = user.login == 'bborbe' && user.password == 'mazdazx';
    $log.debug('verifyLogin login: ' + user.login + ' password: ' + user.password + ' => ' + result);
    return result;
  };

  return service;
}]);

angular.module('bookingServices').factory('ShootingService', ['$log', 'Shooting', function ($log, Shooting) {
  var service = {};

  service.create = function (data) {
    $log.debug('create shooting');
    return Shooting.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update shooting with id: ' + data.id);
    return Shooting.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete shooting with id: ' + id);
    return Shooting.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get shooting with id: ' + id);
    return Shooting.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list shootings');
    return Shooting.query().$promise;
  };

  return service;
}]);


angular.module('bookingServices').factory('ModelService', ['$log', 'Model', function ($log, Model) {
  var service = {};

  service.create = function (data) {
    $log.debug('create model');
    return Model.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update model with id: ' + data.id);
    return Model.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete model with id: ' + id);
    return Model.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get model with id: ' + id);
    return Model.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list models');
    return Model.query().$promise;
  };

  service.findByToken = function (token) {
    $log.debug('find model by token: ' + token);
    return Model.query({token: token}).$promise;
  };

  return service;
}]);

angular.module('bookingServices').factory('DateService', ['$log', 'Date', function ($log, Date) {
  var service = {};

  service.create = function (data) {
    $log.debug('create date');
    return Date.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update date with id: ' + data.id);
    return Date.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete date with id: ' + id);
    return Date.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get date with id: ' + id);
    return Date.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list dates');
    return Date.query().$promise;
  };

  return service;
}]);

angular.module('bookingServices').factory('Model', ['$resource', function ($resource) {
  return $resource('/model/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);

angular.module('bookingServices').factory('Shooting', ['$resource', function ($resource) {
  return $resource('/shooting/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);


angular.module('bookingServices').factory('Date', ['$resource', function ($resource) {
  return $resource('/date/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);
