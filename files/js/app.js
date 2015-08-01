'use strict';

angular.module('bookingApp', [
  'ngRoute',
  'bookingControllers',
  'bookingDirectives',
  'bookingFilters',
  'bookingServices'
]);

angular.module('bookingApp').config(['$routeProvider', function ($routeProvider) {
  $routeProvider.
    when('/login', {
      templateUrl: 'partials/login.html',
      controller: 'LoginCtrl'
    }).
    when('/shooting/create', {
      templateUrl: 'partials/shooting/create.html',
      controller: 'ShootingCreateCtrl'
    }).
    when('/shooting/list', {
      templateUrl: 'partials/shooting/list.html',
      controller: 'ShootingListCtrl'
    }).
    when('/shooting/show/:shootingId', {
      templateUrl: 'partials/shooting/show.html',
      controller: 'ShootingShowCtrl'
    }).
    when('/model/create', {
      templateUrl: 'partials/model/create.html',
      controller: 'ModelCreateCtrl'
    }).
    when('/model/list', {
      templateUrl: 'partials/model/list.html',
      controller: 'ModelListCtrl'
    }).
    when('/model/complete/:token', {
      templateUrl: 'partials/model/complete.html',
      controller: 'ModelCompleteCtrl'
    }).
    otherwise({
      redirectTo: '/shooting/create'
    });
}]);
