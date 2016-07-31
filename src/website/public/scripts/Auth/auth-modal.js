'use strict';

(function () {

    angular.module("app.modal", ['app.rest', 'app.auth']).controller('AuthModalCtrl', AuthModalCtrl);

    AuthModalCtrl.$inject = ['$rootScope', '$auth', '$api', '$uibModalInstance', '$cookies', '$timeout', 'forRegister'];

    function AuthModalCtrl($rootScope, $auth, $api, $uibModalInstance, $cookies, $timeout, forRegister) {

        var vm = this;
        var ApiUser = $api.user();

        vm.forRegister = forRegister;
        vm.firstName = "";
        vm.lastName = "";
        vm.email = "";
        vm.password = "";
        vm.error = "";
        vm.facebookId = "";

        vm.dismiss = function () {
            $uibModalInstance.dismiss('cancel');
        };

        vm.toggleRegister = function (toggle) {
            vm.forRegister = toggle;
        };

        vm.authenticate = function () {

            var defer = $q.defer();

            //$timeout(function () {
            //
            //    var authData = {
            //        email: vm.email,
            //        password: vm.password
            //    };
            //
            //    auth.$signin(function (res) {
            //        $rootScope.currentUser = res.data;
            //        $cookies.putObject('user', res.data);
            //        defer.resolve();
            //        $timeout(function () {
            //            $uibModalInstance.close();
            //        }, 300);
            //    }, function (res) {
            //        vm.error = res.data.message;
            //        defer.reject(res.data.message);
            //    });
            //
            //}, 500);

            return defer.promise;
        };

        vm.login = function () {
            vm.error = "";
            vm.loading = true;
            ApiUser.login({
                email: vm.email,
                password: vm.password
            }, function (res) {
                vm.loading = false;
                $auth.setUser(res);
                $uibModalInstance.close(res);
            }, function (res) {
                vm.error = "Login failed: " + res.data;
                vm.loading = false;
            });
        };

        vm.loginFB = function () {
            FB.login(function (response) {
                console.log(response);
                // The response object is returned with a status field that lets the
                // app know the current login status of the person.
                // Full docs on the response object can be found in the documentation
                // for FB.getLoginStatus().
                if (response.status === 'connected') {
                    vm.loginFBSuccess();
                } else if (response.status === 'not_authorized') {
                    // The person is logged into Facebook, but not your app.
                } else {
                    // The person is not logged into Facebook, so we're not sure if
                    // they are logged into this app or not.
                }
            }, {scope: 'email'});
        };

        vm.loginFBSuccess = function () {
            FB.api('/me', {fields: 'last_name,first_name,email'}, function (me) {
                vm.email = me.email;
                vm.firstName = me.first_name;
                vm.lastName = me.last_name;
                vm.facebookId = me.id;
                ApiUser.loginFacebook({
                    email: vm.email,
                    firstName: vm.firstName,
                    lastName: vm.lastName,
                    facebookId: vm.facebookId
                }, function (res) {
                    vm.loading = false;
                    $auth.setUser(res);
                    $uibModalInstance.close(res);
                }, function (res) {
                    vm.error = "Login failed: " + res.data;
                    vm.loading = false;
                });
            });
        };

        vm.register = function () {
            vm.error = "";
            vm.loading = true;
            ApiUser.register({
                email: vm.email,
                password: vm.password,
                firstName: vm.firstName,
                lastName: vm.lastName
            }, function (res) {
                vm.login();
            }, function (res) {
                vm.error = "Register failed: " + res.data;
                vm.loading = false;
            })
        };
    }

})();
