'use strict';

(function () {

    angular.module('app.payment').controller('PaymentSuccessCtrl', PaymentSuccessCtrl);

    PaymentSuccessCtrl.$inject = ['$timeout', '$api', '$injectedModel'];

    function PaymentSuccessCtrl($timeout, $api, $injectedModel) {

        var duration = 30000;
        var vm = this;
        var ApiTransaction = $api.transaction();

        vm.LOADING = 'LOADING';
        vm.SUCCESS = 'SUCCESS';
        vm.ERROR = 'ERROR';
        vm.TIMEOUT = 'TIMEOUT';

        vm.backerId = $injectedModel.order_code;
        vm.error = '';
        vm.state = vm.LOADING;

        vm.timeout = function () {
            if (vm.state == vm.LOADING) {
                vm.state = vm.TIMEOUT;
            }
        };

        vm.verifyTransaction = function () {
            vm.state = vm.LOADING;
            vm.expired = new Date().getTime() + duration;
            $timeout(function () {
                vm.timeout();
            }, duration);
            var postData = {
                transactionInfo: $injectedModel.transaction_info,
                price: $injectedModel.price,
                paymentId: $injectedModel.payment_id,
                paymentType: $injectedModel.payment_type,
                errorText: $injectedModel.error_text || "",
                secureCode: $injectedModel.secure_code,
                orderCode: $injectedModel.order_code,
            };
            ApiTransaction.verify({}, postData,
                function (res) {
                    vm.state = vm.SUCCESS;
                },
                function (res) {
                    //TODO show error
                    vm.state = vm.ERROR;
                    vm.error = res.data;
                }
            );
        };

        vm.verifyTransaction();
    }

})();
