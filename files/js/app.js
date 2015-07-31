'use strict';

angular.module('bookingApp', [
  'ngRoute',
  'bookingControllers',
  'bookingDirectives',
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
      when('/booking', {
        templateUrl: 'partials/booking_create.html',
        controller: 'BookingCreateCtrl'
      }).
      otherwise({
        redirectTo: '/login'
      });
  }]);
