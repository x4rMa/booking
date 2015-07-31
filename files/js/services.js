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
  service.list = [];

  service.createShooting = function (shooting) {
    $log.debug('create shooting with name: ' + shooting.name);

    if (!shooting.name) {
      $log.debug('create shooting failed without name');
      return;
    }
    service.counter = service.counter + 1;
    $log.debug('increase counter to ' + service.counter);
    var data = {
      'id': service.counter,
      'name': shooting.name,
      'token': 'abc',
    };
    service.list.push(data);
    $log.debug('create shooting successful with id: ' + data.id);
    return data;
  };

  service.listShootings = function () {
    return service.list;
  };

  service.getShooting = function (id) {
    var result;
    angular.forEach(service.list, function (shooting) {
      if (id == shooting.id) {
        result = shooting;
      }
    });
    return result;
  };

  return service;
}]);


