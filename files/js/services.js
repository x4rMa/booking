'use strict';

angular.module('bookingServices', ['ngResource']);

angular.module('bookingServices').factory('AuthService', ['$log', function ($log) {
  var service = {};

  service.verifyLogin = function (user) {
    var result = user.login == 'bborbe' && user.password == 'mazdazx';
    $log.debug('verifyLogin login: ' + user.login + ' password: ' + user.password + ' => ' + result);
    return result;
  };

  return service;
}]);

angular.module('bookingServices').factory('ShootingService', ['$log', function ($log) {
  var service = {};

  service.counter = 0;
  service.shootings = {};

  service.createShooting = function (data) {
    $log.debug('create shooting with name: ' + data.name);

    if (!data.name) {
      $log.debug('create shooting failed without name');
      return;
    }
    var shooting = {};
    shooting['id'] = ++service.counter;
    shooting['name'] = data.name;
    shooting['modelId'] = data.modelId;
    service.shootings[shooting.id] = shooting;
    $log.debug('create shooting successful with id: ' + shooting.id);
    return shooting;
  };

  service.listShootings = function () {
    $log.debug('listShootings');
    var list = [];
    for (var id in service.shootings) {
      $log.debug('add shooting with id: ' + id + ' to result');
      list.push(service.shootings[id]);
    }
    return list;
  };

  service.deleteShooting = function (id) {
    $log.debug('deleteShooting');
    delete service.shootings[id];
    return;
  };

  service.getShooting = function (id) {
    $log.debug('getShooting with id: ' + id);
    return angular.copy(service.shootings[id]);
  };


  service.updateShooting = function (data) {
    $log.debug('updateShooting with id: ' + data.id);
    var shooting = service.shootings[data.id];
    if (shooting) {
      for (var key in data) {
        shooting[key] = data[key];
      }
      return service.getShooting(shooting.id);
    } else {
      return;
    }
  };

  return service;
}]);


angular.module('bookingServices').factory('ModelService', ['$log', function ($log) {
  var service = {};

  service.counter = 0;
  service.models = {};

  service.createModel = function (data) {
    $log.debug('createModel');
    var model = {};
    model['id'] = ++service.counter;
    model['firstname'] = data.firstname;
    model['lastname'] = data.lastname;
    model['email'] = data.email;
    model['phone'] = data.phone;
    service.models[model.id] = model;
    return service.getModel(model.id);
  };

  service.updateModel = function (data) {
    $log.debug('updateModel with id: ' + data.id);
    var model = service.models[data.id];
    if (model) {
      for (var key in data) {
        model[key] = data[key];
      }
      return service.getModel(model.id);
    } else {
      return;
    }
  };

  service.deleteModel = function (id) {
    $log.debug('deleteModel');
    delete service.models[id];
    return;
  };

  service.getModel = function (id) {
    $log.debug('getModel with id: ' + id);
    return angular.copy(service.models[id]);
  };

  service.listModels = function () {
    $log.debug('listModels');
    var list = [];
    for (var id in service.models) {
      $log.debug('add model with id: ' + id + ' to result');
      list.push(service.models[id]);
    }
    return list;
  };

  return service;
}]);

angular.module('bookingServices').factory('DateService', ['$log', 'Date', function ($log, Date) {
  var service = {};

  service.createDate = function (data) {
    $log.debug('create date');
    return Date.create(data);
  };

  service.updateDate = function (data) {
    $log.debug('update date with id: ' + data.id);
    return Date.update(data);
  };

  service.deleteDate = function (id) {
    $log.debug('delete date with id: ' + id);
    Date.delete({Id: id});
    return;
  };

  service.getDate = function (id) {
    $log.debug('get date with id: ' + id);
    return Date.get({Id: id});
  };

  service.listDates = function () {
    $log.debug('list dates');
    return Date.query();
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
