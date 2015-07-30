'use strict';

angular.module('bookingServices', ['ngResource']);

angular.module('bookingServices').factory('User', ['$resource',
  function ($resource) {
    return $resource(':userId.json', {}, {
      query: {method: 'GET', params: {userId: 'users'}, isArray: true}
    });
  }]);