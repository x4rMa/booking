'use strict';

angular.module('bookingDirectives', []);

angular.module('bookingDirectives').directive('datetime', ['$log', function ($log) {
  return {
    require: 'ngModel',
    link: function (scope, elm, attrs, ctrl) {
      ctrl.$parsers.unshift(function (viewValue) {
        $log.debug('datetime');
        // 2015-08-12T15:45:30+02:00
        var INTEGER_REGEXP = /^\d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d([+-]\d\d:\d\d)?$/;
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
