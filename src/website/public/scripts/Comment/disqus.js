'use strict';

(function () {

    angular.module('HoneyStart').directive('disqus', Disqus);
    Disqus.$inject = ['$timeout', '$api', '$http'];

    function Disqus($timeout, $api, $http) {
        return {
            templateUrl: 'scripts/Comment/disqus.html?v=' + window.app_version,
            restrict: 'E',
            replace: true,
            link: function (scope, element, attrs) {
                $timeout(scope.loadScript(), 0);
            },
            controller: function ($scope, $rootScope) {

                $scope.rootScope = $rootScope;

                $scope.loadScript = function () {
                    $api.disqus().success(function (data) {
                        $('#disqus_wrapper').append(data);
                    });
                };


            }
        }
    }

})();
