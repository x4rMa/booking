'use strict';

angular.module('bookingServices', []);

angular.module('bookingServices').factory('AuthService', ['$log',
  function ($log) {
    var service = {};
    service.verifyLogin = function (user) {
      var result = user.login == 'bborbe' && user.password == 'mazdazx';
      $log.debug('verifyLogin login: ' + user.login + ' password: ' + user.password + ' => ' + result);
      return result;
    };
    return service;
  }]);

angular.module('bookingServices').factory('ShootingService', ['$log',
  function ($log) {
    var service = {};
    service.createShooting = function (shooting) {
      $log.debug('create shooting with name: ' + shooting.name);
      return true;
    };
    service.listShootings = function () {
      return [];
    };
    service.getShooting = function () {
      return {
        'link': 'http://example.com'
      };
    };
    return service;
  }]);


