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
    // shooting
    when('/shooting/create', {
      templateUrl: 'partials/shooting/form.html',
      controller: 'ShootingCreateCtrl'
    }).
    when('/shooting/update/:shootingId', {
      templateUrl: 'partials/shooting/form.html',
      controller: 'ShootingUpdateCtrl'
    }).
    when('/shooting/list', {
      templateUrl: 'partials/shooting/list.html',
      controller: 'ShootingListCtrl'
    }).
    when('/shooting/show/:shootingId', {
      templateUrl: 'partials/shooting/show.html',
      controller: 'ShootingShowCtrl'
    }).
    // model
    when('/model/create', {
      templateUrl: 'partials/model/form.html',
      controller: 'ModelCreateCtrl'
    }).
    when('/model/update/:modelId', {
      templateUrl: 'partials/model/form.html',
      controller: 'ModelUpdateCtrl'
    }).
    when('/model/list', {
      templateUrl: 'partials/model/list.html',
      controller: 'ModelListCtrl'
    }).
    when('/model/complete/:modelId', {
      templateUrl: 'partials/model/complete.html',
      controller: 'ModelCompleteCtrl'
    }).
    when('/model/show/:modelId', {
      templateUrl: 'partials/model/show.html',
      controller: 'ModelShowCtrl'
    }).
    // date
    when('/date/create', {
      templateUrl: 'partials/date/form.html',
      controller: 'DateCreateCtrl'
    }).
    when('/date/update/:dateId', {
      templateUrl: 'partials/date/form.html',
      controller: 'DateUpdateCtrl'
    }).
    when('/date/list', {
      templateUrl: 'partials/date/list.html',
      controller: 'DateListCtrl'
    }).
    when('/date/select', {
      templateUrl: 'partials/date/select.html',
      controller: 'DateSelectCtrl'
    }).
    otherwise({
      redirectTo: '/login'
    });
}]);
