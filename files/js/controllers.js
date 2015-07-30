'use strict';

angular.module('bookingControllers', []);

angular.module('bookingControllers').controller('LoginCtrl', ['$scope', 'User', function ($scope, User) {
  $scope.users = User.query();
}]);