'use strict';

angular.module('bookingDirectives', []);

angular.module('bookingDirectives').directive('datetime', ['$log', function ($log) {
  return {
    require: 'ngModel',
    link: function (scope, elm, attrs, ctrl) {
      ctrl.$parsers.unshift(function (viewValue) {
        $log.debug('datetime');
        var INTEGER_REGEXP = /^\d\d\d\d-\d\d-\d\d[\sT]\d\d:\d\d(:\d\d)$/;
        if (INTEGER_REGEXP.test(viewValue)) {
          // it is valid
          ctrl.$setValidity('integer', true);
          return viewValue;
        } else {
          // it is invalid, return undefined (no model update)
          ctrl.$setValidity('integer', false);
          return undefined;
        }
      });
    }
  }
}]);
