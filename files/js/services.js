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
