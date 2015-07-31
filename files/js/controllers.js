'use strict';

angular.module('bookingControllers', []);

angular.module('bookingControllers').controller('LoginCtrl', ['$scope', '$log', 'AuthService', '$location', function ($scope, $log, AuthService, $location) {
  $scope.user = {};
  $scope.submit = function () {
    if (AuthService.verifyLogin($scope.user)) {
      $log.debug('login success => redirect to create booking');
      $scope.user = {};
      $location.path('/booking');
    } else {
      $log.debug('login failed');
      //$scope.loginForm.$error = 'login failed';
    }
  };
}]);

angular.module('bookingControllers').controller('BookingCreateCtrl', ['$scope', function ($scope) {
}]);