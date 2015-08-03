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
    when('/shooting/update/:Id', {
      templateUrl: 'partials/shooting/form.html',
      controller: 'ShootingUpdateCtrl'
    }).
    when('/shooting/list', {
      templateUrl: 'partials/shooting/list.html',
      controller: 'ShootingListCtrl'
    }).
    when('/shooting/show/:Id', {
      templateUrl: 'partials/shooting/show.html',
      controller: 'ShootingShowCtrl'
    }).
    when('/shooting/delete/:Id', {
      template: '',
      controller: 'ShootingDeleteCtrl'
    }).
    // model
    when('/model/create', {
      templateUrl: 'partials/model/form.html',
      controller: 'ModelCreateCtrl'
    }).
    when('/model/update/:Id', {
      templateUrl: 'partials/model/form.html',
      controller: 'ModelUpdateCtrl'
    }).
    when('/model/list', {
      templateUrl: 'partials/model/list.html',
      controller: 'ModelListCtrl'
    }).
    when('/model/complete/:Token', {
      templateUrl: 'partials/model/complete.html',
      controller: 'ModelCompleteCtrl'
    }).
    when('/model/show/:Id', {
      templateUrl: 'partials/model/show.html',
      controller: 'ModelShowCtrl'
    }).
    when('/model/delete/:Id', {
      template: '',
      controller: 'ModelDeleteCtrl'
    }).
    // date
    when('/date/create', {
      templateUrl: 'partials/date/form.html',
      controller: 'DateCreateCtrl'
    }).
    when('/date/update/:Id', {
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
    when('/date/delete/:Id', {
      template: '',
      controller: 'DateDeleteCtrl'
    }).
    otherwise({
      redirectTo: '/login'
    });
}]);
