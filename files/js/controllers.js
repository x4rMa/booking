'use strict';

var bookingControllers = angular.module('bookingControllers', []);

bookingControllers.controller('UserListCtrl', ['$scope', 'User', function($scope, User) {
	$scope.users = User.query();
}]);