'use strict';

(function () {

    angular.module("app.auth", ['app.modal']).service('$auth', AuthService);

    AuthService.$inject = ['$rootScope', '$cookies', '$uibModal'];

    function AuthService($rootScope, $cookies, $uibModal) {

        var vm = this;

        vm.setUser = function (user) {
            $cookies.putObject('user', user);
            $rootScope.user = user;
        };

        vm.getUser = function () {
            if (!$rootScope.user) {
                $rootScope.user = $cookies.getObject('user');
            }
            return $rootScope.user;œ
        };

        vm.getToken = function () {
            var user = vm.getUser();
            if (user) {
                return user.token;
            }
        };

        vm.showModal = function(forRegister, onSuccess) {
            var modal = $uibModal.open({
                animation: true,
                size: 'sm',
                templateUrl: '/scripts/Auth/auth-modal.html?v=' + window.app_version,
                controller: 'AuthModalCtrl',
                controllerAs: 'vm',
                resolve: {
                    forRegister: function () {
                        return forRegister;
                    }
                }
            });
            modal.result.then(onSuccess);
        };
    }

})();