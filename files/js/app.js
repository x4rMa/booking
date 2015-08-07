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
    when('/logout', {
      template: '',
      controller: 'LogoutCtrl'
    }).
    // user
    when('/user/create', {
      templateUrl: 'partials/user/form.html',
      controller: 'UserCreateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('administrator');
        },
      }
    }).
    when('/user/update/:Id', {
      templateUrl: 'partials/user/form.html',
      controller: 'UserUpdateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('administrator');
        },
      }
    }).
    when('/user/list', {
      templateUrl: 'partials/user/list.html',
      controller: 'UserListCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('administrator');
        },
      }
    }).
    when('/user/show/:Id', {
      templateUrl: 'partials/user/show.html',
      controller: 'UserShowCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('administrator');
        },
      }
    }).
    when('/user/delete/:Id', {
      template: '',
      controller: 'UserDeleteCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('administrator');
        },
      }
    }).
    // shooting
    when('/shooting/create', {
      templateUrl: 'partials/shooting/form.html',
      controller: 'ShootingCreateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/shooting/update/:Id', {
      templateUrl: 'partials/shooting/form.html',
      controller: 'ShootingUpdateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/shooting/list', {
      templateUrl: 'partials/shooting/list.html',
      controller: 'ShootingListCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/shooting/show/:Id', {
      templateUrl: 'partials/shooting/show.html',
      controller: 'ShootingShowCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/shooting/delete/:Id', {
      template: '',
      controller: 'ShootingDeleteCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/shooting/select', {
      templateUrl: 'partials/shooting/select.html',
      controller: 'ShootingSelectCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('participant');
        },
      }
    }).
    // model
    when('/model/create', {
      templateUrl: 'partials/model/form.html',
      controller: 'ModelCreateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/model/update/:Id', {
      templateUrl: 'partials/model/form.html',
      controller: 'ModelUpdateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/model/list', {
      templateUrl: 'partials/model/list.html',
      controller: 'ModelListCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/model/complete', {
      templateUrl: 'partials/model/complete.html',
      controller: 'ModelCompleteCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('participant');
        },
      }
    }).
    when('/model/show/:Id', {
      templateUrl: 'partials/model/show.html',
      controller: 'ModelShowCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/model/delete/:Id', {
      template: '',
      controller: 'ModelDeleteCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/model/auth/:Token', {
      template: '',
      controller: 'ModelAuthCtrl'
    }).
    // date
    when('/date/create', {
      templateUrl: 'partials/date/form.html',
      controller: 'DateCreateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/date/update/:Id', {
      templateUrl: 'partials/date/form.html',
      controller: 'DateUpdateCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/date/list', {
      templateUrl: 'partials/date/list.html',
      controller: 'DateListCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/date/select', {
      templateUrl: 'partials/date/select.html',
      controller: 'DateSelectCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('participant');
        },
      }
    }).
    when('/date/delete/:Id', {
      template: '',
      controller: 'DateDeleteCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('organizer');
        },
      }
    }).
    when('/date/select/:ShootingId', {
      templateUrl: 'partials/date/select.html',
      controller: 'DateSelectCtrl',
      resolve: {
        permission: function (AuthorizationService, $route) {
          return AuthorizationService.checkPermission('participant');
        },
      }
    }).
    otherwise({
      redirectTo: '/shooting/list'
    });
}]);
