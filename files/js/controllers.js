'use strict';

angular.module('bookingControllers', []);

angular.module('bookingControllers').controller('LoginCtrl', ['$scope', '$log', '$location', 'AuthService', function ($scope, $log, $location, AuthService) {
  $scope.reset = function () {
    $log.debug('reset login form');
    $scope.user = {};
  };
  $scope.submit = function () {
    if (AuthService.verifyLogin($scope.user)) {
      $log.debug('login success => redirect to create booking');
      $scope.reset();
      $location.path('/shooting/create');
    } else {
      $log.debug('login failed');
      //$scope.loginForm.$error = 'login failed';
    }
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('ShootingCreateCtrl', ['$scope', '$log', '$location', 'ShootingService', function ($scope, $log, $location, ShootingService) {
  $scope.reset = function () {
    $log.debug('reset create shooting form');
    $scope.shooting = {};
  }
  $scope.submit = function () {
    var shooting = ShootingService.createShooting($scope.shooting);
    if (shooting) {
      $log.debug('create shooting success');
      $scope.reset();
      $location.path('/shooting/show/' + shooting.id);
    } else {
      $log.debug('create shooting failed');
    }
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('ShootingListCtrl', ['$scope', 'ShootingService', function ($scope, ShootingService) {
  $scope.shootings = ShootingService.listShootings();
}]);

angular.module('bookingControllers').controller('ShootingShowCtrl', ['$scope', '$routeParams', '$log', 'ShootingService', function ($scope, $routeParams, $log, ShootingService) {
  $log.debug('show shooting with id: ' + $routeParams.shootingId);
  $scope.shooting = ShootingService.getShooting($routeParams.shootingId);
}]);

angular.module('bookingControllers').controller('ModelListCtrl', ['$scope', 'ModelService', function ($scope, ModelService) {
  $scope.models = ModelService.listModels();
}]);

angular.module('bookingControllers').controller('ModelCreateCtrl', ['$scope', '$log', 'ModelService', function ($scope, $log, ModelService) {
  $scope.reset = function () {
    $log.debug('reset create model form');
    $scope.model = {};
  }
  $scope.submit = function () {
    var model = ModelService.createModel($scope.model);
    if (model) {
      $log.debug('create model success');
      $scope.reset();
      $location.path('/model/show/' + model.id);
    } else {
      $log.debug('create model failed');
    }
  };
  $scope.reset();
}]);
