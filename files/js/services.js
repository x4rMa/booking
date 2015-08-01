'use strict';

angular.module('bookingServices', []);

angular.module('bookingServices').factory('AuthService', ['$log', function ($log) {
  var service = {};

  service.verifyLogin = function (user) {
    var result = user.login == 'bborbe' && user.password == 'mazdazx';
    $log.debug('verifyLogin login: ' + user.login + ' password: ' + user.password + ' => ' + result);
    return result;
  };

  return service;
}]);

angular.module('bookingServices').factory('ShootingService', ['$log', function ($log) {
  var service = {};

  service.counter = 0;
  service.shootings = {};

  service.createShooting = function (data) {
    $log.debug('create shooting with name: ' + data.name);

    if (!data.name) {
      $log.debug('create shooting failed without name');
      return;
    }
    var shooting = {
      'id': ++service.counter,
      'name': data.name,
      'token': 'abc',
    };
    service.shootings[shooting.id] = shooting;
    $log.debug('create shooting successful with id: ' + shooting.id);
    return shooting;
  };

  service.listShootings = function () {
    $log.debug('listShootings');
    var list = [];
    for (var id in service.shootings) {
      $log.debug('add shooting with id: ' + id + ' to result');
      list.push(service.shootings[id]);
    }
    return list;
  };

  service.deleteShooting = function (id) {
    $log.debug('deleteShooting');
    delete service.shootings[id];
  };

  service.getShooting = function (id) {
    $log.debug('getShooting with id: '+id);
    return service.shootings[id];
  };

  return service;
}]);


angular.module('bookingServices').factory('ModelService', ['$log', function ($log) {
  var service = {};

  service.counter = 0;
  service.models = {};

  service.createModel = function (data) {
    $log.debug('createModel');
    var model = {
      'id': ++service.counter,
      'firstnme': data.firstname,
      'lastname': data.lastname,
      'email': data.email,
      'phone': data.phone
    };

    service.models[model.id] = model;
    return model;
  };

  service.deleteModel = function (id) {
    $log.debug('deleteModel');
    delete service.models[id];
  };

  service.getModel = function (id) {
    $log.debug('getModel with id: '+id);
    return service.models[id];
  };

  service.listModels = function () {
    $log.debug('listModels');
    var list = [];
    for (var id in service.models) {
      $log.debug('add model with id: ' + id + ' to result');
      list.push(service.models[id]);
    }
    return list;
  };

  return service;
}]);
