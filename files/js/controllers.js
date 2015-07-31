'use strict';

angular.module('bookingControllers', []);

angular.module('bookingControllers').controller('LoginCtrl', ['$scope', '$log', '$location', 'AuthService', function ($scope, $log, $location, AuthService) {
  $scope.user = {};
  $scope.submit = function () {
    if (AuthService.verifyLogin($scope.user)) {
      $log.debug('login success => redirect to create booking');
      $scope.user = {};
      $location.path('/shooting/create');
    } else {
      $log.debug('login failed');
      //$scope.loginForm.$error = 'login failed';
    }
  };
}]);

angular.module('bookingControllers').controller('ShootingCreateCtrl', ['$scope', '$log', '$location', 'ShootingService', function ($scope, $log, $location, ShootingService) {
  $scope.shooting = {};
  $scope.submit = function () {
    if (ShootingService.createShooting($scope.shooting)) {
      $log.debug('create shooting success');
      $scope.shooting = {};
      $location.path('/shooting/show');
    } else {
      $log.debug('create shooting failed');
    }
  };
}]);

angular.module('bookingControllers').controller('ShootingListCtrl', ['$scope', 'ShootingService', function ($scope, ShootingService) {
}]);

angular.module('bookingControllers').controller('ShootingShowCtrl', ['$scope', 'ShootingService', function ($scope, ShootingService) {
  $scope.shooting = ShootingService.getShooting();
}]);
