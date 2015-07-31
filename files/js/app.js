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
      when('/shooting/create', {
        templateUrl: 'partials/shooting_create.html',
        controller: 'ShootingCreateCtrl'
      }).
      when('/shooting/list', {
        templateUrl: 'partials/shooting_list.html',
        controller: 'ShootingListCtrl'
      }).
      when('/shooting/show/:shootingId', {
        templateUrl: 'partials/shooting_show.html',
        controller: 'ShootingShowCtrl'
      }).
      otherwise({
        redirectTo: '/login'
      });
  }]);
