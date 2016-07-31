'use strict';

(function () {

    angular.module("app.home", ['app.rest']).controller('HomeCtrl', HomeCtrl);

    HomeCtrl.$inject = ['$rootScope', '$window', '$api'];

    function HomeCtrl($rootScope, $window, $api) {

        var vm = this;
        var ApiProject = $api.project();

        vm.features = [
            {
                icon: 'fa-rocket',
                title: 'FAST',
                description: 'High Life narwhal, banh mi PBR single-origin coffee Odd Future actually aliqua polaroid befor',
            },
            {
                icon: 'fa-credit-card',
                title: 'CONVENIENT',
                description: 'Neutra Thundercats craft beer, listicle meggings bicycle rights 90s XOXO beard cardiga',
            },
            {
                icon: 'fa-gift',
                title: 'MEANINGFUL GIFT',
                description: 'We design beautiful modern engaging websites that always latest responsive technologies.',
            },
            {
                icon: 'fa-lock',
                title: 'SECURE',
                description: 'Food truck master cleanse mixtape minim Portland, cardigan stumptown chambray',
            },
        ];

        vm.projects = [];

        vm.loadProjectsHome = function () {
            ApiProject.getHome(function (res) {
                vm.projects = res;
            });
        };

        vm.loadProjectsHome();
    }

})();
