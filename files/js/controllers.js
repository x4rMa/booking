'use strict';

angular.module('bookingControllers', []);

angular.module('bookingControllers').controller('LoginCtrl', ['$scope', '$log', '$location', 'AuthenticationService', function ($scope, $log, $location, AuthenticationService) {
  $scope.reset = function () {
    $log.debug('reset login form');
    $scope.user = {};
    $scope.messages = [];
  };
  $scope.submit = function () {
    AuthenticationService.login($scope.user).then(function (result) {
      $log.debug('login = ' + result);
      if (result) {
        $log.debug('login success => redirect');
        $scope.reset();
        $location.path('/shooting/list');
      } else {
        $log.debug('login failed');
        $scope.messages.push('login failed');
      }
    });
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('LogoutCtrl', ['$scope', '$log', '$location', 'AuthenticationService', function ($scope, $log, $location, AuthenticationService) {
  $log.debug('logout');
  AuthenticationService.logout();
  $location.path('/');
}]);

angular.module('bookingControllers').controller('NaviCtrl', ['$scope', '$log', 'AuthenticationService', 'AuthorizationService', function ($scope, $log, AuthenticationService, AuthorizationService) {
  $log.debug('navi');
  $scope.isAdministrator = AuthorizationService.hasRole('administrator');
  $scope.isOrganizer = AuthorizationService.hasRole('organizer');
  $scope.isParticipant = AuthorizationService.hasRole('participant');
  $scope.isLoggedIn = AuthenticationService.isLoggedIn();
}]);

angular.module('bookingControllers').controller('ShootingCreateCtrl', ['$scope', '$log', '$location', 'ShootingService', 'ModelService', 'DateService', function ($scope, $log, $location, ShootingService, ModelService, DateService) {
  ModelService.list().then(function (result) {
    $scope.models = result;
  });
  DateService.list().then(function (result) {
    $scope.dates = result;
  });
  $scope.reset = function () {
    $log.debug('reset create shooting form');
    $scope.shooting = {};
    $scope.messages = [];
  }
  $scope.submit = function () {
    $scope.shooting.model_id = parseInt($scope.shooting.model_id);
    ShootingService.create($scope.shooting).then(function (result) {
      $log.debug('create shooting success with id: ' + shooting.id);
      $scope.reset();
      $location.path('/shooting/show/' + shooting.id);
    }, function (error) {
      $log.debug('create shooting failed: ' + error);
      $scope.messages.push(error);
    });
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('ShootingShowCtrl', ['$scope', '$routeParams', '$log', '$location', 'ShootingService', function ($scope, $routeParams, $log, $location, ShootingService) {
  $log.debug('show shooting with id: ' + $routeParams.Id);
  ShootingService.get($routeParams.Id).then(function (result) {
    $log.debug('shooting found');
    $scope.shooting = result;
  }, function (error) {
    $log.debug('shooting not found: ' + error);
    $location.path('/shooting/list');
  });
}]);

angular.module('bookingControllers').controller('ShootingUpdateCtrl', ['$scope', '$routeParams', '$log', '$location', 'ShootingService', 'ModelService', 'DateService', function ($scope, $routeParams, $log, $location, ShootingService, ModelService, DateService) {
  $log.debug('show shooting with id: ' + $routeParams.Id);
  ModelService.list().then(function (result) {
    $scope.models = result;
  });
  DateService.list().then(function (result) {
    $scope.dates = result;
  });
  ShootingService.get($routeParams.Id).then(function (result) {
    $scope.shooting = result;
  }, function (error) {
    $log.debug('shooting not found');
    $location.path('/shooting/list');
  });
  $scope.submit = function () {
    ShootingService.update($scope.shooting).then(function (result) {
      $log.debug('update shooting success');
      $location.path('/shooting/show/' + shooting.id);
    }, function (error) {
      $log.debug('update shooting failed: ' + error);
    });
  };
}]);

angular.module('bookingControllers').controller('ShootingListCtrl', ['$scope', '$log', 'ShootingService', function ($scope, $log, ShootingService) {
  ShootingService.list().then(function (result) {
    $log.debug('list ' + result.length + ' shootings');
    $scope.shootingsWithoutDate = [];
    $scope.shootingsWithDate = [];
    angular.forEach(result, function (shooting) {
      if (shooting.date_id && shooting.date_id > 0) {
        $log.debug('push to shootingsWithDate');
        $scope.shootingsWithDate.push(shooting);
      } else {
        $log.debug('push to shootingsWithoutDate');
        $scope.shootingsWithoutDate.push(shooting);
      }
    });
  }, function (error) {
    $log.debug('list shootings failed: ' + error);
  });
}]);

angular.module('bookingControllers').controller('ShootingDeleteCtrl', ['$scope', '$routeParams', '$log', '$location', 'ShootingService', function ($scope, $routeParams, $log, $location, ShootingService) {
  $log.debug('delete shooting with id: ' + $routeParams.Id);
  ShootingService.delete($routeParams.Id).then(function (result) {
    $log.debug('shooting deleted');
    $location.path('/shooting/list');
  }, function (error) {
    $log.debug('delete shooting failed: ' + error);
    $location.path('/shooting/list');
  });
}]);

angular.module('bookingControllers').controller('ShootingSelectCtrl', ['$scope', '$log', '$location', 'ShootingService', function ($scope, $log, $location, ShootingService) {
  ShootingService.list().then(function (result) {
    $log.debug('list ' + result.lenght + ' shootings for select success');
    $scope.shootingsWithoutDate = [];
    $scope.shootingsWithDate = [];
    angular.forEach(result, function (shooting) {
      if (shooting.date_id && shooting.date_id > 0) {
        $log.debug('push to shootingsWithDate');
        $scope.shootingsWithDate.push(shooting);
      } else {
        $log.debug('push to shootingsWithoutDate');
        $scope.shootingsWithoutDate.push(shooting);
      }
    });
  }, function (error) {
    $log.debug('list shootings for select failed: ' + error);
  });
}]);

angular.module('bookingControllers').controller('ModelListCtrl', ['$scope', '$log', 'ModelService', function ($scope, $log, ModelService) {
  ModelService.list().then(function (result) {
    $log.debug('list models success');
    $scope.models = result;
  }, function (error) {
    $log.debug('list models failed: ' + error);
  });
}]);

angular.module('bookingControllers').controller('ModelCreateCtrl', ['$scope', '$log', '$location', 'ModelService', function ($scope, $log, $location, ModelService) {
  $scope.reset = function () {
    $log.debug('reset create model form');
    $scope.model = {};
    $scope.messages = [];
  }
  $scope.submit = function () {
    ModelService.create($scope.model).then(function (result) {
      $log.debug('create model success with id: ' + result.id);
      $scope.reset();
      $location.path('/model/show/' + result.id);
    }, function (error) {
      $log.debug('create model failed: ' + error);
      $scope.messages.push(error);
    });
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('ModelShowCtrl', ['$scope', '$routeParams', '$log', '$location', 'ModelService', function ($scope, $routeParams, $log, $location, ModelService) {
  $log.debug('show model with id: ' + $routeParams.Id);
  ModelService.get($routeParams.Id).then(function (result) {
    $scope.link = $location.protocol() + '://' + $location.host() + ':' + $location.port() + '/#/model/auth/' + result.token;
    $scope.model = result;
  }, function (error) {
    $log.debug('model not found: ' + error);
    $location.path('/model/list');
  });
}]);

angular.module('bookingControllers').controller('ModelUpdateCtrl', ['$scope', '$routeParams', '$log', '$location', 'ModelService', function ($scope, $routeParams, $log, $location, ModelService) {
  $log.debug('update model with id: ' + $routeParams.Id);
  ModelService.get($routeParams.Id).then(function (result) {
    $scope.model = result;
  }, function (error) {
    $log.debug('model not found: ' + error);
    $location.path('/model/list');
  });
  $scope.submit = function () {
    ModelService.update($scope.model).then(function (data) {
      $log.debug('update model success');
      $location.path('/model/show/' + data.id);
    }, function (error) {
      $log.debug('update model failed: ' + error);
    });
  };
}]);

angular.module('bookingControllers').controller('ModelDeleteCtrl', ['$scope', '$routeParams', '$log', '$location', 'ModelService', function ($scope, $routeParams, $log, $location, ModelService) {
  $log.debug('delete model with id: ' + $routeParams.Id);
  ModelService.delete($routeParams.Id).then(function (result) {
    $log.debug('delete model success');
    $location.path('/model/list');
  }, function (error) {
    $log.debug('delete model failed: ' + error);
    $location.path('/model/list');
  });
}]);

angular.module('bookingControllers').controller('ModelCompleteCtrl', ['$scope', '$routeParams', '$log', '$location', 'ModelService', function ($scope, $routeParams, $log, $location, ModelService) {
  $log.debug('complete model with token: ' + $routeParams.Token);
  ModelService.findByToken($routeParams.Token).then(function (result) {
    $scope.model = result[0];
  }, function (error) {
    $log.debug('find model by token failed: ' + error);
    $location.path('/');
  });
  $scope.submit = function () {
    ModelService.update($scope.model).then(function (result) {
      $log.debug('update model success');
      $location.path('/shooting/select');
    }, function (error) {
      $log.debug('update model failed: ' + error);
    });
  };
}]);

angular.module('bookingControllers').controller('ModelAuthCtrl', ['$scope', '$routeParams', '$log', '$location', 'AuthenticationService', function ($scope, $routeParams, $log, $location, AuthenticationService) {
  $log.debug('auth model with token: ' + $routeParams.Token);
  AuthenticationService.login({'token': $routeParams.Token}).then(function (result) {
    if (result) {
      $log.debug('auth model success');
      $location.path('/model/complete');
    } else {
      $log.debug('auth model failed');
      $location.path('/');
    }
  }, function (error) {
    $log.debug('auth model failed: ' + error);
    $location.path('/');
  });
}]);

angular.module('bookingControllers').controller('DateCreateCtrl', ['$scope', '$log', '$location', 'DateService', function ($scope, $log, $location, DateService) {
  $scope.reset = function () {
    $log.debug('reset create date form');
    $scope.date = {};
    $scope.messages = [];
  }
  $scope.submit = function () {
    DateService.create($scope.date).then(function (result) {
      $log.debug('create date success with id: ' + result.id);
      $scope.reset();
      $location.path('/date/list');
    }, function (error) {
      $log.debug('create date failed: ' + error);
      $scope.messages.push(error);
    });
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('DateUpdateCtrl', ['$scope', '$routeParams', '$log', '$location', 'DateService', function ($scope, $routeParams, $log, $location, DateService) {
  $log.debug('update date with id: ' + $routeParams.Id);
  DateService.get($routeParams.Id).then(function (result) {
    $scope.date = result;
  }, function (error) {
    $log.debug('date not found');
    $location.path('/date/list');
  });
  $scope.submit = function () {
    DateService.update($scope.date).then(function (result) {
      $log.debug('update date success');
      $location.path('/date/list');
    }, function (error) {
      $log.debug('update date failed: ' + error);
    });
  };
}]);

angular.module('bookingControllers').controller('DateListCtrl', ['$scope', '$log', 'DateService', function ($scope, $log, DateService) {
  DateService.list().then(function (result) {
    $log.debug('list dates success');
    $scope.dates = result;
  }, function (error) {
    $log.debug('list dates failed: ' + error);
  });
}]);

angular.module('bookingControllers').controller('DateSelectCtrl', ['$scope', '$routeParams', '$log', 'ShootingService', 'DateService', function ($scope, $routeParams, $log, ShootingService, DateService) {
  $log.debug('select date for shooting with id: ' + $routeParams.ShootingId);
  ShootingService.get($routeParams.ShootingId).then(function (result) {
    $log.debug('get shooting success');
    $scope.shooting = result;
  }, function (error) {
    $log.debug('get shooting failed: ' + error);
  });
  DateService.listFree().then(function (result) {
    $log.debug('list free dates success');
    $scope.dates = result;
  }, function (error) {
    $log.debug('list free dates failed: ' + error);
  });
  $scope.book = function (date_id) {
    $log.debug('book ' + date_id + ' for shooting ' + $routeParams.ShootingId);
    ShootingService.book(date_id, $routeParams.ShootingId).then(function (result) {
      $log.debug('book date for shooting success');
    }, function (error) {
      $log.debug('book date for shooting failed: ' + error);
    });
  };
}]);

angular.module('bookingControllers').controller('DateDeleteCtrl', ['$scope', '$routeParams', '$log', '$location', 'DateService', function ($scope, $routeParams, $log, $location, DateService) {
  $log.debug('delete date with id: ' + $routeParams.Id);
  DateService.delete($routeParams.Id).then(function (result) {
    $log.debug('delete date sucess');
    $location.path('/date/list');
  }, function (error) {
    $log.debug('delete date failed: ' + error);
    $location.path('/date/list');
  });
}]);

angular.module('bookingControllers').controller('UserListCtrl', ['$scope', '$log', 'UserService', function ($scope, $log, UserService) {
  UserService.list().then(function (result) {
    $log.debug('list users success');
    $scope.users = result;
  }, function (error) {
    $log.debug('list users failed: ' + error);
  });
}]);

angular.module('bookingControllers').controller('UserCreateCtrl', ['$scope', '$log', '$location', 'UserService', function ($scope, $log, $location, UserService) {
  $scope.reset = function () {
    $log.debug('reset create user form');
    $scope.user = {};
    $scope.messages = [];
  }
  $scope.submit = function () {
    UserService.create($scope.user).then(function (result) {
      $log.debug('create user success with id: ' + result.id);
      $scope.reset();
      $location.path('/user/show/' + result.id);
    }, function (error) {
      $log.debug('create user failed: ' + error);
      $scope.messages.push(error);
    });
  };
  $scope.reset();
}]);

angular.module('bookingControllers').controller('UserShowCtrl', ['$scope', '$routeParams', '$log', '$location', 'UserService', function ($scope, $routeParams, $log, $location, UserService) {
  $log.debug('show user with id: ' + $routeParams.Id);
  UserService.get($routeParams.Id).then(function (result) {
    $scope.user = result;
  }, function (error) {
    $log.debug('user not found: ' + error);
    $location.path('/user/list');
  });
}]);

angular.module('bookingControllers').controller('UserUpdateCtrl', ['$scope', '$routeParams', '$log', '$location', 'UserService', function ($scope, $routeParams, $log, $location, UserService) {
  $log.debug('update user with id: ' + $routeParams.Id);
  UserService.get($routeParams.Id).then(function (result) {
    $scope.user = result;
  }, function (error) {
    $log.debug('user not found: ' + error);
    $location.path('/user/list');
  });
  $scope.submit = function () {
    UserService.update($scope.user).then(function (data) {
      $log.debug('update user success');
      $location.path('/user/show/' + data.id);
    }, function (error) {
      $log.debug('update user failed: ' + error);
    });
  };
}]);

angular.module('bookingControllers').controller('UserDeleteCtrl', ['$scope', '$routeParams', '$log', '$location', 'UserService', function ($scope, $routeParams, $log, $location, UserService) {
  $log.debug('delete user with id: ' + $routeParams.Id);
  UserService.delete($routeParams.Id).then(function (result) {
    $log.debug('delete user success');
    $location.path('/user/list');
  }, function (error) {
    $log.debug('delete user failed: ' + error);
    $location.path('/user/list');
  });
}]);