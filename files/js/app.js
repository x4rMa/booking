'use strict';

angular.module('bookingApp', [
  'ngRoute',
  'bookingControllers',
  'bookingFilters',
  'bookingServices'
]);

angular.module('bookingApp').config(['$routeProvider',
  function ($routeProvider) {
    $routeProvider.
      when('/login', {
        templateUrl: 'partials/login.html',
        controller: 'LoginCtrl'
      }).
      otherwise({
        redirectTo: '/login'
      });
  }]);
