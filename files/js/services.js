'use strict';

var bookingServices = angular.module('bookingServices', ['ngResource']);

bookingServices.factory('User', ['$resource',
	function ($resource) {
		return $resource(':userId.json', {}, {
			query: {method: 'GET', params: {userId: 'users'}, isArray: true}
		});
	}]);