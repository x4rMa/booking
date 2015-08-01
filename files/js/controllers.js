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

angular.module('bookingControllers').controller('ShootingShowCtrl', ['$scope', '$routeParams', '$log', '$location', 'ShootingService', function ($scope, $routeParams, $log, $location, ShootingService) {
  $log.debug('show shooting with id: ' + $routeParams.shootingId);
  var shooting = ShootingService.getShooting($routeParams.shootingId);
  if (shooting) {
    $scope.shooting = shooting;
  } else {
    $log.debug('shooting not found');
    $location.path('/shooting/list');
  }
}]);

angular.module('bookingControllers').controller('ShootingUpdateCtrl', ['$scope', '$routeParams', '$log', '$location', 'ShootingService', function ($scope, $routeParams, $log, $location, ShootingService) {
  $log.debug('show shooting with id: ' + $routeParams.shootingId);
  var shooting = ShootingService.getShooting($routeParams.shootingId);
  if (shooting) {
    $scope.shooting = shooting;
  } else {
    $log.debug('shooting not found');
    $location.path('/shooting/list');
  }
  $scope.submit = function () {
    var shooting = ShootingService.updateShooting($scope.shooting);
    if (shooting) {
      $log.debug('update shooting success');
      $location.path('/shooting/show/' + shooting.id);
    } else {
      $log.debug('update shooting failed');
    }
  };
}]);

angular.module('bookingControllers').controller('ShootingListCtrl', ['$scope', 'ShootingService', function ($scope, ShootingService) {
  $scope.shootings = ShootingService.listShootings();
}]);

angular.module('bookingControllers').controller('ModelListCtrl', ['$scope', 'ModelService', function ($scope, ModelService) {
  $scope.models = ModelService.listModels();
}]);

angular.module('bookingControllers').controller('ModelCreateCtrl', ['$scope', '$log', '$location', 'ModelService', function ($scope, $log, $location, ModelService) {
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

angular.module('bookingControllers').controller('ModelShowCtrl', ['$scope', '$routeParams', '$log', '$location', 'ModelService', function ($scope, $routeParams, $log, $location, ModelService) {
  $log.debug('show model with id: ' + $routeParams.modelId);
  var model = ModelService.getModel($routeParams.modelId);
  if (model) {
    $scope.model = model;
  } else {
    $log.debug('model not found');
    $location.path('/model/list');
  }
}]);

angular.module('bookingControllers').controller('ModelUpdateCtrl', ['$scope', '$routeParams', '$log', '$location', 'ModelService', function ($scope, $routeParams, $log, $location, ModelService) {
  $log.debug('show model with id: ' + $routeParams.modelId);
  var model = ModelService.getModel($routeParams.modelId);
  if (model) {
    $scope.model = model;
  } else {
    $log.debug('model not found');
    $location.path('/model/list');
  }
  $scope.submit = function () {
    var model = ModelService.updateModel($scope.model);
    if (model) {
      $log.debug('update model success');
      $location.path('/model/show/' + model.id);
    } else {
      $log.debug('update model failed');
    }
  };
}]);
