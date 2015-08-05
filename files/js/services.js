'use strict';

angular.module('bookingServices', ['ngResource', 'ngCookies']);

angular.module('bookingServices').factory('AuthenticationService', ['$log', '$cookieStore', '$http', '$q', function ($log, $cookieStore, $http, $q) {
  var service = {};
  var userkey = 'user';

  service.isLoggedIn = function () {
    var result = !!service.getCurrent();
    $log.debug('isLoggedIn: ' + result);
    return result;
  };

  service.getCurrent = function () {
    $log.debug('getCurrent');
    return $cookieStore.get(userkey);
  };

  service.verifyLogin = function (user) {
    $log.debug('verifyLogin for user: ' + user.login);
    return $http.post('/user/verifyLogin', user);
  };

  service.login = function (user) {
    var deferred = $q.defer();
    service.verifyLogin(user).then(function (result) {
      if (result) {
        $cookieStore.put(userkey, user);
      }
      deferred.resolve(result);
    }, function (error) {
      deferred.reject(error);
    });
    return deferred.promise;
  };

  service.logout = function () {
    $cookieStore.remove(userkey);
  };

  return service;
}]);

angular.module('bookingServices').factory('AuthorizationService', ['$log', '$q', '$location', '$rootScope', 'AuthenticationService', function ($log, $q, $location, $rootScope, AuthenticationService) {
  var service = {};

  service.hasRole = function (role) {
    var result = false;
    if (role == 'administrator' && AuthenticationService.isLoggedIn() && AuthenticationService.getCurrent().login == 'admin') {
      result = true;
    }
    if (role == 'organizer' && AuthenticationService.isLoggedIn()) {
      result = true;
    }
    if (role == 'participant' && AuthenticationService.isLoggedIn()) {
      result = true;
    }
    $log.debug('has role: ' + role + ' => ' + result);
    return result;
  };

  service.checkPermission = function (permission_list) {
    var deferred = $q.defer();
    $log.debug("hasPermission");
    if (AuthenticationService.isLoggedIn()) {
      deferred.resolve();
    } else {
      //If user does not have required access,
      //we will route the user to unauthorized access page
      $location.path('/login');
      //As there could be some delay when location change event happens,
      //we will keep a watch on $locationChangeSuccess event
      // and would resolve promise when this event occurs.
      $rootScope.$on('$locationChangeSuccess', function (next, current) {
        deferred.resolve();
      });
    }
    return deferred.promise;
  };

  return service;
}]);

angular.module('bookingServices').factory('Shooting', ['$resource', function ($resource) {
  return $resource('/shooting/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);

angular.module('bookingServices').factory('ShootingService', ['$log', 'Shooting', function ($log, Shooting) {
  var service = {};

  service.create = function (data) {
    $log.debug('create shooting');
    return Shooting.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update shooting with id: ' + data.id);
    return Shooting.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete shooting with id: ' + id);
    return Shooting.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get shooting with id: ' + id);
    return Shooting.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list shootings');
    return Shooting.query().$promise;
  };

  service.book = function (date_id, shooting_id) {
    $log.debug('book shooting with id: ' + shooting_id);
    return $http.post('/shooting/book', {'id': shooting_id, 'date_id': date_id});
  };

  return service;
}]);

angular.module('bookingServices').factory('Model', ['$resource', function ($resource) {
  return $resource('/model/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);

angular.module('bookingServices').factory('ModelService', ['$log', 'Model', function ($log, Model) {
  var service = {};

  service.create = function (data) {
    $log.debug('create model');
    return Model.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update model with id: ' + data.id);
    return Model.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete model with id: ' + id);
    return Model.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get model with id: ' + id);
    return Model.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list models');
    return Model.query().$promise;
  };

  service.findByToken = function (token) {
    $log.debug('find model by token: ' + token);
    return Model.query({token: token}).$promise;
  };

  return service;
}]);

angular.module('bookingServices').factory('Date', ['$resource', function ($resource) {
  return $resource('/date/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);

angular.module('bookingServices').factory('DateService', ['$log', 'Date', function ($log, Date) {
  var service = {};

  service.create = function (data) {
    $log.debug('create date');
    return Date.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update date with id: ' + data.id);
    return Date.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete date with id: ' + id);
    return Date.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get date with id: ' + id);
    return Date.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list dates');
    return Date.query().$promise;
  };

  return service;
}]);

angular.module('bookingServices').factory('User', ['$resource', function ($resource) {
  return $resource('/user/:Id', {}, {
    query: {
      method: 'GET', params: {}, isArray: true
    },
    create: {
      method: 'POST', params: {}
    },
    update: {
      method: 'PUT', params: {}
    },
    delete: {
      method: 'DELETE', params: {}
    }
  });
}]);

angular.module('bookingServices').factory('UserService', ['$log', 'User', function ($log, User) {
  var service = {};

  service.create = function (data) {
    $log.debug('create user');
    return User.create(data).$promise;
  };

  service.update = function (data) {
    $log.debug('update user with id: ' + data.id);
    return User.update(data).$promise;
  };

  service.delete = function (id) {
    $log.debug('delete user with id: ' + id);
    return User.delete({Id: id}).$promise;
  };

  service.get = function (id) {
    $log.debug('get user with id: ' + id);
    return User.get({Id: id}).$promise;
  };

  service.list = function () {
    $log.debug('list users');
    return User.query().$promise;
  };

  return service;
}]);